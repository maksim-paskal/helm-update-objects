package internal

import (
	"context"
	"flag"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/maksim-paskal/helm-update-objects/pkg/client"
	"github.com/maksim-paskal/helm-update-objects/pkg/config"
	"github.com/maksim-paskal/helm-update-objects/pkg/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	namespace  = flag.String("namespace", "", "namespace to convert")
	name       = flag.String("release-name", "", "name of the release to convert")
	dryRun     = flag.Bool("dry-run", true, "dry run")
	k8sVersion = flag.String("kubernetes-version", "", "manual kubernetes version")
)

var rules = config.GetRules()

func Run(ctx context.Context) error { //nolint:cyclop,funlen
	kubernetesVersion, err := client.Clientset.ServerVersion()
	if err != nil {
		return errors.Wrap(err, "error getting kubernetes version")
	}

	kubernetesVersionString := kubernetesVersion.String()

	if len(*k8sVersion) > 0 {
		log.Warnf("Using manual kubernetes version %s, server version %s", *k8sVersion, kubernetesVersionString)
		kubernetesVersionString = *k8sVersion
	}

	semverKubernetes, err := semver.NewVersion(kubernetesVersionString)
	if err != nil {
		return errors.Wrap(err, "error cast to semver")
	}

	log.Infof("Kubernetes version %s", semverKubernetes.String())

	for _, rule := range rules {
		if !rule.Ignore(semverKubernetes) {
			log.Debugf("rule %s", rule.Regexp)
		}
	}

	releases, err := client.Clientset.CoreV1().Secrets(*namespace).List(ctx, metav1.ListOptions{
		LabelSelector: "owner=helm,status=deployed",
	})
	if err != nil {
		return errors.Wrap(err, "error selecting secrets")
	}

	found := false

	for _, secret := range releases.Items {
		releaseName, releaseNameOK := secret.Labels["name"]
		if !releaseNameOK {
			return errors.New("release name not found")
		}

		releaseNameFull := fmt.Sprintf("%s.%s", secret.Namespace, releaseName)

		// filter by name
		if len(*name) > 0 && releaseNameFull != *name {
			continue
		}

		log.Debugf("found %s", releaseNameFull)

		found = true

		if err := upgradeSecret(ctx, semverKubernetes, secret); err != nil {
			return errors.Wrapf(err, "error upgrading secret %s", releaseNameFull)
		}
	}

	if !found {
		return errors.New("release not found")
	}

	return nil
}

func upgradeSecret(ctx context.Context, kubernetesVersion *semver.Version, secret corev1.Secret) error { //nolint:cyclop
	releaseName, releaseNameOK := secret.Labels["name"]
	if !releaseNameOK {
		return errors.New("release name not found")
	}

	releaseRaw, releaseRawOK := secret.Data["release"]
	if !releaseRawOK {
		return errors.New("release field not found")
	}

	releaseNameFull := fmt.Sprintf("%s.%s", secret.Namespace, releaseName)

	release, err := utils.DecodeRelease(string(releaseRaw))
	if err != nil {
		return errors.Wrap(err, "error decoding release")
	}

	if *dryRun {
		for _, rule := range rules {
			if rule.Match(kubernetesVersion, release.Manifest) {
				log.Warningf("needs update %s (has %s it deprecated in %s)", releaseNameFull, rule.Regexp, rule.AfterSemver)
			}
		}

		return nil
	}

	replaced := release.Manifest

	for _, rule := range rules {
		replaced = rule.ReplaceAllString(kubernetesVersion, replaced)
	}

	if replaced == release.Manifest {
		log.Warningf("no changes %s", releaseNameFull)

		return nil
	}
	// return replaced object
	release.Manifest = replaced

	convertedReleaseRaw, err := utils.EncodeRelease(release)
	if err != nil {
		return errors.Wrap(err, "error encoding release")
	}

	convertedSecret := secret.DeepCopy()
	convertedSecret.Data["release"] = []byte(convertedReleaseRaw)

	_, err = client.Clientset.CoreV1().Secrets(secret.Namespace).Update(ctx, convertedSecret, metav1.UpdateOptions{})
	if err != nil {
		return errors.Wrap(err, "error saving secret")
	}

	log.Infof("updated %s", releaseNameFull)

	return nil
}
