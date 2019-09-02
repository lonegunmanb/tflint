// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsAcmCertificateInvalidCertificateChainRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_acm_certificate" "foo" {
	certificate_chain = <<TEXT
-----BEGIN PRIVATE KEY-----
MIIEpAIBAAKCAQEAuY6k+pUBFm8jXhzBHiycUTe8D5nyV31kiS+Nr2DzPDcvDcVL
VBPHmGr7CUEqdj3MAcVWGrYFTzXi97GjxsXxqZCLmj4EDtZ+2vjvt1W/xgDXt4PU
V/96dvNAx3ZvqaaUaevMoYrzK7c541yQnmzosfinYBwTC1KsduxVpqnISfr0O+MJ
rg0dWOX4evTQ5+/ZAfGSYlKWVkNAj64KIvCsVlPBU9CZ0MQbxOgifNT9v9S7Sf/5
4GrMrsd+490euEOGS4E1/VbQF/gZ6MvTYfcE2+s244escztWzeCMGhb3PstheYGb
3Lepud6fw37JCCPlrh+hnOpx6zf4P1ePVaPt2QIDAQABAoIBAGN8EAXtV3zwrzSp
E/0ai+Cbki+HKUAxEXLf1QX/Y8mYCJlIex+jzzJvwRHwJ1TnwvX8GDMP/f6+9GY6
joVm4S85OS/EKibOZ4r9RoCz77K4Bu/0NSfM6JrXxpZqcGmzzwSPENJXjhKVFOtK
WJsn5wZsO0izJJ7Af4jvIujNRH4sq8oO0T5OHNlJprrR8KY7+7IJrAd+JyCWxuX3
uToqrQALpZjSLi0+60+UZIVH+F4yJtkar1MS+K5bAAtX6gxaq3embIOD5KCooHpf
+CXBuWhUpFIqdWWDKJunnbriuXPBVVx1BZBtALUMPEYZ+gR22l2s/ck0BGixsdcN
VwDluuUCgYEA9k64okPnwUAQXyoHnA09zSJ0acr4aX/lqixY+fx0yn5Rs46WuFl0
OkUSZ3YMoj5MM/hmrMgz2v7bleVDrsZYUM4//9bYVxRIf1/gHm6PRJDnAMDEiz7y
crJCD1HTk/4CTcufbVnyyDt7NQe+qL9wRbTUMIojdZ1PsThv4zOuNhMCgYEAwNvu
LXhUJ7Wc92xFwRbmhNw8nQu0YgwVkGiD5QMY/6FJhcGdLa/Ih3NvyZI+Q/x/qzTh
udCHHlMgAftx5uA7IpZxQo0ak5a/T46X1EvRqcPZvOyDliBeYhhAXflyNfusmm5X
cExL6g3aUXxa7ue255xkKo2Q1VH+b9YeP+pZeeMCgYEA2VqEkiTMWfvXtrLXPj1t
IR3bjxQe/LJxkCdMaWYABkVMgeA9XvcJmvYjFIvXAEFra50zthuBryqhyfgkLxI3
Ey++yFzmUonCpCyOESzNXttkDoUNrDdjKhXmN7Ckvf80N0SOLqhml43t3tEzzaQK
RmkZqq/sNLkafzBnhB6yCGMCgYEAqtFhiXadmzpZ2DBXLCobbSwgp7zZPUqUwv4/
bFUtDCYQF9+gVvnuNELDjZbxfYgkkEDbeZhARVS88eSDQ0nyNrVnhdmy42xO8KlM
w2WQQ7xLm/Ekr5Dl6B6wzEuHpFbQz0vSOI3rY1h3uVras+Yac9RqR+JxmO/x256b
1mK8c58CgYAA2zWV1m34tdbUmhc6ELNznbPxaM4sAe7mVooSxs02pTN5ToxMpkyn
hLL+1AMBOhv57k6YDTzYf5adR38023TOrf6NuD9G5s0KeaFM8c7SrBt8BHA6MW9C
7kJKP9RgFcXPFH6Z57rKykRaaDR6M+ELrhd488tOmqlwS9tnnWuoBQ==
-----END PRIVATE KEY-----
TEXT
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsAcmCertificateInvalidCertificateChainRule(),
					Message: `certificate_chain does not match valid pattern ^(-{5}BEGIN CERTIFICATE-{5}\x{000D}?\x{000A}([A-Za-z0-9/+]{64}\x{000D}?\x{000A})*[A-Za-z0-9/+]{1,64}={0,2}\x{000D}?\x{000A}-{5}END CERTIFICATE-{5}\x{000D}?\x{000A})*-{5}BEGIN CERTIFICATE-{5}\x{000D}?\x{000A}([A-Za-z0-9/+]{64}\x{000D}?\x{000A})*[A-Za-z0-9/+]{1,64}={0,2}\x{000D}?\x{000A}-{5}END CERTIFICATE-{5}(\x{000D}?\x{000A})?$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_acm_certificate" "foo" {
	certificate_chain = <<TEXT
-----BEGIN CERTIFICATE-----
MIIEWjCCA0KgAwIBAgIQBUNA0KLEzIER+qg3fUbgbzANBgkqhkiG9w0BAQsFADBa
MQswCQYDVQQGEwJJRTESMBAGA1UEChMJQmFsdGltb3JlMRMwEQYDVQQLEwpDeWJl
clRydXN0MSIwIAYDVQQDExlCYWx0aW1vcmUgQ3liZXJUcnVzdCBSb290MB4XDTE2
MTExNTEyMDMzMVoXDTI1MDUxMDEyMDAwMFowWjELMAkGA1UEBhMCSlAxIzAhBgNV
BAoTGkN5YmVydHJ1c3QgSmFwYW4gQ28uLCBMdGQuMSYwJAYDVQQDEx1DeWJlcnRy
dXN0IEphcGFuIFB1YmxpYyBDQSBHMzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAJRWo0VEVKpgZL+4V59O29R5aF8TBfQ/zSXdPF5Ydxyd5p/jMknvAjo0
U41S5eM5Zh/nM2G2J8YkVVAnAmXwsIxBjTBeR1uCb8ecoyhDbVh7yBWYTiVvy3Yn
WwssLLWYI+eLfP13GsRSul0Z7nghTSGa2RJ8MxVrGsmB6traV7fVL84fS/y0M+Cg
yZQnuydAtpDbrJ51phErSRktw8JDBwm7PW6Io+OKxdKG9mVbNMOfTALlCbosxnZm
69F2JfQwE/tYYKhY41FvSwgEYY2sqTAvUkGjIsEzWat7WfmTZ0vJiXVS7ylJNJMc
nJNznBnOXBjNTAknwT/1Sez04t9Lr48CAwEAAaOCARowggEWMB0GA1UdDgQWBBRz
qAhTKbgV+5mA5cU32Pg5e6QTBjAfBgNVHSMEGDAWgBTlnVkwgkdYzKz6CFQ2hns6
tQRN8DASBgNVHRMBAf8ECDAGAQH/AgEAMA4GA1UdDwEB/wQEAwIBhjA0BggrBgEF
BQcBAQQoMCYwJAYIKwYBBQUHMAGGGGh0dHA6Ly9vY3NwLmRpZ2ljZXJ0LmNvbTA6
BgNVHR8EMzAxMC+gLaArhilodHRwOi8vY3JsMy5kaWdpY2VydC5jb20vT21uaXJv
b3QyMDI1LmNybDA+BgNVHSAENzA1MDMGBWeBDAECMCowKAYIKwYBBQUHAgEWHGh0
dHBzOi8vd3d3LmRpZ2ljZXJ0LmNvbS9DUFMwDQYJKoZIhvcNAQELBQADggEBABqR
cPGWxhjLeEP8gzZ29OHVpL8EFWNvpnEbpnAWa1uYmFNf0fsSoibP04UCB0Ye21KA
jPNuTE2OWpgxlTVvl99sIJFKhXCIn6OWCLQpbjFMu8Ro5rdzg4ekedsihWN9NMwU
+K8DqO+oukIYjF2zPnnwcSykKaUhdi0BqUtmzJhHkFjoslJllOXiJ4DQkbpjiQpR
yjPfliI1wzAcOM3xxswHAxOq8BVHKErKYUqHkHgHFZ6Ycts+sShKqVf+Zdx9zYJV
pItiSs0DDCCFapxWACouIVY+xPYG8EMUWXz0gK8SAwNHXLRxBtjNWhCGOiSN+ihp
277L8LbKp8eA8OlOMiU=
-----END CERTIFICATE-----
TEXT
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsAcmCertificateInvalidCertificateChainRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
