package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestGoldenSynapseHPA covers all autoscalingWorkers HPAs with default values
// (HPA enabled, KEDA disabled for all workers).
func TestGoldenSynapseHPA(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-hpa",
		Templates:      []string{"templates/synapse-hpa.yaml"},
	})
}

// TestGoldenSynapseSecret covers the full homeserver.yaml config secret for all workers.
func TestGoldenSynapseSecret(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-secret",
		Templates:      []string{"templates/synapse-secret.yaml"},
	})
}

// TestGoldenSynapseSecretDisablePgBouncerForStreamWriters verifies that stream writer
// workers connect directly to postgres (not pgbouncer) when the flag is enabled.
func TestGoldenSynapseSecretDisablePgBouncerForStreamWriters(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-secret-no-pgbouncer-stream-writers",
		Templates:      []string{"templates/synapse-secret.yaml"},
		SetValues:      map[string]string{"synapse.disablePgBouncerForStreamWriters": "true"},
	})
}

// TestGoldenSynapseService covers all worker services with default values.
func TestGoldenSynapseService(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-service",
		Templates:      []string{"templates/synapse-service.yaml"},
	})
}

// TestGoldenSynapseWorkersDeployment covers all autoscalingWorkers deployments.
func TestGoldenSynapseWorkersDeployment(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-workers-deployment",
		Templates:      []string{"templates/synapse-workers-deployment.yaml"},
	})
}

// TestGoldenSynapseWorkersPDB covers PDBs for deployScalingWorkers and singletonWorkers.
// deployScalingWorkers with replicas>1 get minAvailable, replicas=1 get maxUnavailable:1.
// singletonWorkers always get maxUnavailable:1.
func TestGoldenSynapseWorkersPDB(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-workers-pdb",
		Templates:      []string{"templates/synapse-workers-pdb.yaml"},
	})
}

// TestGoldenSynapseWorkersPDBDisabled verifies nothing is rendered when the PDB is disabled.
func TestGoldenSynapseWorkersPDBDisabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-workers-pdb-disabled",
		Templates:      []string{"templates/synapse-workers-pdb.yaml"},
		SetValues:      map[string]string{"synapse.deployScalingWorkersPdb.enabled": "false"},
		AllowEmpty:     true,
	})
}

// TestGoldenSynapseWorkersStatefulset covers all deployScalingWorkers and singletonWorkers
// StatefulSets with default values.
func TestGoldenSynapseWorkersStatefulset(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-workers-statefulset",
		Templates:      []string{"templates/synapse-workers-statefulset.yaml"},
	})
}
