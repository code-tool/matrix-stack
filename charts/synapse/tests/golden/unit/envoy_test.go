package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestGoldenEnvoyConfigmap covers the default routing configuration:
// matrixAuthentication enabled, msc4306 disabled.
// Verifies that MAS login/logout routes point to httpd-matrix-auth cluster
// and threadSubscriptions routes are absent.
func TestGoldenEnvoyConfigmap(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-envoy-configmap",
		Templates:      []string{"templates/envoy-configmap.yaml"},
	})
}

// TestGoldenEnvoyConfigmapNoMAS covers routing when matrixAuthentication is
// disabled: clientReaderRegister routes (login$, register$, password_policy$)
// must appear and point to httpd-client-reader; httpd-matrix-auth cluster
// must be absent.
func TestGoldenEnvoyConfigmapNoMAS(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-envoy-configmap-no-mas",
		Templates:      []string{"templates/envoy-configmap.yaml"},
		SetValues:      map[string]string{"matrixAuthentication.enabled": "false"},
	})
}

// TestGoldenEnvoyConfigmapMsc4306 covers routing when thread_subscriptions MSC
// is enabled: threadSubscriptionsRoutes and httpd-thread-subscriptions cluster
// must appear.
func TestGoldenEnvoyConfigmapMsc4306(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-envoy-configmap-msc4306",
		Templates:      []string{"templates/envoy-configmap.yaml"},
		SetValues:      map[string]string{"experimentalFeatures.msc4306.enabled": "true"},
	})
}

func TestGoldenEnvoyDeployment(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-envoy-deployment",
		Templates:      []string{"templates/envoy-deployment.yaml"},
	})
}

func TestGoldenEnvoyService(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-envoy-service",
		Templates:      []string{"templates/envoy-service.yaml"},
	})
}

// TestGoldenEnvoyPDB covers the default PDB (minAvailable: 1).
func TestGoldenEnvoyPDB(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-envoy-pdb",
		Templates:      []string{"templates/envoy-pdb.yaml"},
	})
}

// TestGoldenEnvoyPDBMaxUnavailable verifies that maxUnavailable is rendered
// instead of minAvailable when explicitly set.
func TestGoldenEnvoyPDBMaxUnavailable(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-envoy-pdb-max-unavailable",
		Templates:      []string{"templates/envoy-pdb.yaml"},
		SetValues:      map[string]string{"envoyProxy.podDisruptionBudget.maxUnavailable": "1"},
	})
}

// TestGoldenEnvoyServiceMonitor covers the default ServiceMonitor (metrics enabled).
func TestGoldenEnvoyServiceMonitor(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-envoy-servicemonitor",
		Templates:      []string{"templates/envoy-servicemonitor.yaml"},
	})
}

// TestGoldenEnvoyServiceMonitorDisabled verifies nothing is rendered
// when envoyProxy.metrics is false.
func TestGoldenEnvoyServiceMonitorDisabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-envoy-servicemonitor-disabled",
		Templates:      []string{"templates/envoy-servicemonitor.yaml"},
		SetValues:      map[string]string{"envoyProxy.metrics": "false"},
		AllowEmpty:     true,
	})
}
