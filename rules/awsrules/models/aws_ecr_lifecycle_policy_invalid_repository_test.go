// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsEcrLifecyclePolicyInvalidRepositoryRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ecr_lifecycle_policy" "foo" {
	repository = "example@com"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsEcrLifecyclePolicyInvalidRepositoryRule(),
					Message: `repository does not match valid pattern ^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ecr_lifecycle_policy" "foo" {
	repository = "example"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsEcrLifecyclePolicyInvalidRepositoryRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
