// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsCloudfrontDistributionInvalidHTTPVersionRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudfront_distribution" "foo" {
	http_version = "http1.2"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudfrontDistributionInvalidHTTPVersionRule(),
					Message: `http_version is not a valid value`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudfront_distribution" "foo" {
	http_version = "http2"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudfrontDistributionInvalidHTTPVersionRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
