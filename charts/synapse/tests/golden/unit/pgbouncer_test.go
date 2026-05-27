package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestGoldenPgbouncerDeployment(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-pgbouncer-deployment",
		Templates:      []string{"templates/pgbouncer-deployment.yaml"},
	})
}

func TestGoldenPgbouncerSecret(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-pgbouncer-secret",
		Templates:      []string{"templates/pgbouncer-secret.yaml"},
	})
}

func TestGoldenPgbouncerService(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-pgbouncer-service",
		Templates:      []string{"templates/pgbouncer-service.yaml"},
	})
}

func TestGoldenPgbouncerServiceSessionAffinity(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-pgbouncer-service-session-affinity",
		Templates:      []string{"templates/pgbouncer-service.yaml"},
		SetValues:      map[string]string{"synapse.pgbouncer.sessionAffinityTimeoutSeconds": "3600"},
	})
}

func TestGoldenPgbouncerConfigmap(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-configmap-namespace",
		GoldenFileName: "test-pgbouncer-configmap",
		Templates:      []string{"templates/pgbouncer-configmap.yaml"},
	})
}
