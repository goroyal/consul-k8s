package connectinject

import (
	"testing"

	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/kubernetes/fake"
)

func TestRun_FlagValidation(t *testing.T) {
	cases := []struct {
		name   string
		flags  []string
		expErr string
	}{
		{
			flags:  []string{},
			expErr: "-consul-k8s-image must be set",
		},
		{
			flags:  []string{"-consul-k8s-image", "foo", "-ca-file", "bar"},
			expErr: "Error reading Consul's CA cert file \"bar\"",
		},
		{
			flags:  []string{"-consul-k8s-image", "foo", "-default-sidecar-proxy-cpu-limit=unparseable"},
			expErr: "-default-sidecar-proxy-cpu-limit is invalid",
		},
		{
			flags:  []string{"-consul-k8s-image", "foo", "-default-sidecar-proxy-cpu-request=unparseable"},
			expErr: "-default-sidecar-proxy-cpu-request is invalid",
		},
		{
			flags:  []string{"-consul-k8s-image", "foo", "-default-sidecar-proxy-memory-limit=unparseable"},
			expErr: "-default-sidecar-proxy-memory-limit is invalid",
		},
		{
			flags:  []string{"-consul-k8s-image", "foo", "-default-sidecar-proxy-memory-request=unparseable"},
			expErr: "-default-sidecar-proxy-memory-request is invalid",
		},
	}

	for _, c := range cases {
		t.Run(c.expErr, func(t *testing.T) {
			k8sClient := fake.NewSimpleClientset()
			ui := cli.NewMockUi()
			cmd := Command{
				UI:        ui,
				clientset: k8sClient,
			}
			code := cmd.Run(c.flags)
			require.Equal(t, 1, code)
			require.Contains(t, ui.ErrorWriter.String(), c.expErr)
		})
	}
}
