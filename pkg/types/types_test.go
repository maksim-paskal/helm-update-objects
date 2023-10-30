package types_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/maksim-paskal/helm-update-objects/pkg/types"
)

func TestRules(t *testing.T) {
	t.Parallel()

	kubernetesVersion, err := semver.NewVersion("1.0.0")
	if err != nil {
		t.Fatal(err)
	}

	rules := []types.Rule{
		{Regexp: `test1`, Replace: `test1/a`, AfterSemver: `1.0.0`},
		{Regexp: `test2`, Replace: `test2/b`},
		{Regexp: `test3`, Replace: `test3/c`},
		{Regexp: `test4`, Replace: `test4/d`, AfterSemver: `1.1.0`},
	}

	valid := make(map[string]string)

	valid["test1 test test1"] = `test1/a test test1/a`
	valid["test2 test test1"] = `test2/b test test1/a`
	valid["test1 test2 test3"] = `test1/a test2/b test3/c`
	valid["test2 test test1"] = `test2/b test test1/a`
	valid["test4"] = `test4`

	for k, v := range valid {
		input := k

		for _, rule := range rules {
			input = rule.ReplaceAllString(kubernetesVersion, input)
		}

		if input != v {
			t.Errorf("expected %s, got %s", v, input)
		}
	}
}
