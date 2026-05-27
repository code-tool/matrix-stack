package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestGoldenIngressDefault covers the default routing configuration:
// matrixAuthentication enabled, msc4306 disabled.
// Verifies that MAS login/logout routes point to matrix-authentication service
// and threadSubscriptions routes are absent.
func TestGoldenIngressDefault(t *testing.T) {
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

// TestGoldenIngressWithoutMAS covers routing when matrixAuthentication is disabled:
// clientReaderRegister routes (login, register) must appear and point to
// synapse-client-reader-envoy instead of matrix-authentication.
func TestGoldenIngressWithoutMAS(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-ingress-no-mas",
		Templates:      []string{"templates/synapse-ingress.yaml"},
		SetValues:      map[string]string{"matrixAuthentication.enabled": "false"},
	})
}

// TestGoldenIngressWithMsc4306 covers routing when thread_subscriptions MSC is enabled:
// threadSubscriptionsRoutes must appear and point to synapse-thread-subscriptions service.
func TestGoldenIngressWithMsc4306(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-ingress-msc4306",
		Templates:      []string{"templates/synapse-ingress.yaml"},
		SetValues:      map[string]string{"experimentalFeatures.msc4306.enabled": "true"},
	})
}
