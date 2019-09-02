// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsDatasyncTaskInvalidSourceLocationArnRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_datasync_task" "foo" {
	source_location_arn = "arn:aws:datasync:us-east-2:111222333444:task/task-08de6e6697796f026"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsDatasyncTaskInvalidSourceLocationArnRule(),
					Message: `source_location_arn does not match valid pattern ^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_datasync_task" "foo" {
	source_location_arn = "arn:aws:datasync:us-east-2:111222333444:location/loc-07db7abfc326c50fb"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsDatasyncTaskInvalidSourceLocationArnRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
