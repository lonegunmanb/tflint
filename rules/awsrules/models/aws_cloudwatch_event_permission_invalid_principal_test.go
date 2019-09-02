// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsCloudwatchEventPermissionInvalidPrincipalRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudwatch_event_permission" "foo" {
	principal = "-"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudwatchEventPermissionInvalidPrincipalRule(),
					Message: `principal does not match valid pattern ^(\d{12}|\*)$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudwatch_event_permission" "foo" {
	principal = "*"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudwatchEventPermissionInvalidPrincipalRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
