package config

import "github.com/maksim-paskal/helm-update-objects/pkg/types"

// rules to replace.
// https://kubernetes.io/docs/reference/using-api/deprecation-guide/
func GetRules() []types.Rule { //nolint:funlen
	return []types.Rule{
		{
			Regexp:      `admissionregistration.k8s.io/v1beta1`,
			Replace:     `admissionregistration.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `apiextensions.k8s.io/v1beta1`,
			Replace:     `apiextensions.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `apiregistration.k8s.io/v1beta1`,
			Replace:     `apiregistration.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `authentication.k8s.io/v1beta1`,
			Replace:     `authentication.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `authorization.k8s.io/v1beta1`,
			Replace:     `authorization.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `certificates.k8s.io/v1beta1`,
			Replace:     `certificates.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `coordination.k8s.io/v1beta1`,
			Replace:     `coordination.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `extensions/v1beta1`,
			Replace:     `networking.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `networking.k8s.io/v1beta1`,
			Replace:     `networking.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `rbac.authorization.k8s.io/v1beta1`,
			Replace:     `rbac.authorization.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `scheduling.k8s.io/v1beta1`,
			Replace:     `scheduling.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `storage.k8s.io/v1beta1`,
			Replace:     `storage.k8s.io/v1`,
			AfterSemver: `1.22.0`,
		},
		{
			Regexp:      `policy/v1beta1`,
			Replace:     `policy/v1`,
			AfterSemver: `1.25.0`,
		},
		{
			Regexp:      `PodSecurityPolicy`,
			Replace:     `PodDisruptionBudget`, // replace for fake object
			AfterSemver: `1.25.0`,
		},
		{
			Regexp:      `autoscaling/v2beta1`,
			Replace:     `autoscaling/v2`,
			AfterSemver: `1.25.0`,
		},
		{
			Regexp:      `batch/v1beta1`,
			Replace:     `batch/v1`,
			AfterSemver: `1.25.0`,
		},
		{
			Regexp:      `discovery.k8s.io/v1beta1`,
			Replace:     `discovery.k8s.io/v1`,
			AfterSemver: `1.25.0`,
		},
		{
			Regexp:      `events.k8s.io/v1beta1`,
			Replace:     `events.k8s.io/v1`,
			AfterSemver: `1.25.0`,
		},
		{
			Regexp:      `node.k8s.io/v1beta1`,
			Replace:     `node.k8s.io/v1`,
			AfterSemver: `1.25.0`,
		},
		{
			Regexp:      `flowcontrol.apiserver.k8s.io/v1beta1`,
			Replace:     `flowcontrol.apiserver.k8s.io/v1beta3`,
			AfterSemver: `1.26.0`,
		},
		{
			Regexp:      `autoscaling/v2beta2`,
			Replace:     `autoscaling/v2`,
			AfterSemver: `1.26.0`,
		},
		{
			Regexp:      `flowcontrol.apiserver.k8s.io/v1beta2`,
			Replace:     `flowcontrol.apiserver.k8s.io/v1beta3`,
			AfterSemver: `1.29.0`,
		},
	}
}
