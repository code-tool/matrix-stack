package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestGoldenEnvoyConfigmap covers the envoy configmap that embeds
// scripts/envoy.yaml and scripts/synapse.lua from chart files.
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
