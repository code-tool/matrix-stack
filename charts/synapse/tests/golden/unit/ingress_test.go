package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestGoldenIngress verifies the simplified ingress that routes all traffic to
// synapse-client-reader-envoy (path: /, pathType: Prefix). All routing logic
// now lives in the envoy configmap — see envoy_test.go for routing variants.
func TestGoldenIngress(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-ingress-default",
		Templates:      []string{"templates/synapse-ingress.yaml"},
	})
}
