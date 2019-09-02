// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsEc2FleetInvalidExcessCapacityTerminationPolicyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ec2_fleet" "foo" {
	excess_capacity_termination_policy = "remain"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsEc2FleetInvalidExcessCapacityTerminationPolicyRule(),
					Message: `excess_capacity_termination_policy is not a valid value`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ec2_fleet" "foo" {
	excess_capacity_termination_policy = "termination"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsEc2FleetInvalidExcessCapacityTerminationPolicyRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
