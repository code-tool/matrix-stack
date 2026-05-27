package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestGoldenSynapsePodMonitorDisabled covers the default state: podMonitor disabled.
func TestGoldenSynapsePodMonitorDisabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-podmonitor-disabled",
		Templates:      []string{"templates/synapse-podmonitor.yaml"},
		AllowEmpty:     true,
	})
}

// TestGoldenSynapsePodMonitorEnabled verifies the PodMonitor is rendered
// when synapse.podMonitor.enabled is true.
func TestGoldenSynapsePodMonitorEnabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-podmonitor-enabled",
		Templates:      []string{"templates/synapse-podmonitor.yaml"},
		SetValues:      map[string]string{"synapse.podMonitor.enabled": "true"},
	})
}

// TestGoldenSynapseResourceQuota covers the default ResourceQuota (enabled: true).
func TestGoldenSynapseResourceQuota(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-resource-quota",
		Templates:      []string{"templates/synapse-resource-quota.yaml"},
	})
}

// TestGoldenSynapseResourceQuotaDisabled verifies nothing is rendered
// when resourceQuota.enabled is false.
func TestGoldenSynapseResourceQuotaDisabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-resource-quota-disabled",
		Templates:      []string{"templates/synapse-resource-quota.yaml"},
		SetValues:      map[string]string{"resourceQuota.enabled": "false"},
		AllowEmpty:     true,
	})
}

// TestGoldenPgbouncerPDB covers the default PDB (minAvailable: 1).
func TestGoldenPgbouncerPDB(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-pgbouncer-pdb",
		Templates:      []string{"templates/pgbouncer-deployment-pdb.yaml"},
	})
}

// TestGoldenPgbouncerPDBMaxUnavailable verifies that maxUnavailable is rendered
// instead of minAvailable when explicitly set.
func TestGoldenPgbouncerPDBMaxUnavailable(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-pgbouncer-pdb-max-unavailable",
		Templates:      []string{"templates/pgbouncer-deployment-pdb.yaml"},
		SetValues:      map[string]string{"synapse.pgbouncer.podDisruptionBudget.maxUnavailable": "1"},
	})
}

// TestGoldenPgbouncerPDBDisabled verifies nothing is rendered
// when the pgbouncer PDB is disabled.
func TestGoldenPgbouncerPDBDisabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-pgbouncer-pdb-disabled",
		Templates:      []string{"templates/pgbouncer-deployment-pdb.yaml"},
		SetValues:      map[string]string{"synapse.pgbouncer.podDisruptionBudget.enabled": "false"},
		AllowEmpty:     true,
	})
}
