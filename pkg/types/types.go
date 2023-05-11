package types

import (
	"regexp"

	"github.com/Masterminds/semver/v3"
)

type Rule struct {
	Regexp      string
	Replace     string
	AfterSemver string
}

func (r *Rule) Ignore(kubernetesVersion *semver.Version) bool {
	if len(r.AfterSemver) == 0 {
		return false
	}

	afterSemverVersion, err := semver.NewVersion(r.AfterSemver)
	if err != nil {
		panic(err)
	}

	if kubernetesVersion.Equal(afterSemverVersion) {
		return false
	}

	return !kubernetesVersion.GreaterThan(afterSemverVersion)
}

func (r *Rule) ReplaceAllString(kubernetesVersion *semver.Version, s string) string {
	if r.Ignore(kubernetesVersion) {
		return s
	}

	ruleRegexp := regexp.MustCompile(r.Regexp)

	return ruleRegexp.ReplaceAllString(s, r.Replace)
}

func (r *Rule) Match(kubernetesVersion *semver.Version, s string) bool {
	if r.Ignore(kubernetesVersion) {
		return false
	}

	ruleRegexp := regexp.MustCompile(r.Regexp)

	return ruleRegexp.MatchString(s)
}
