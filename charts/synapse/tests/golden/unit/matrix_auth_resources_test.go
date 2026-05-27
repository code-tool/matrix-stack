package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestGoldenMatrixAuthService(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-service",
		Templates:      []string{"templates/matrix-authentication-service.yaml"},
	})
}

// TestGoldenMatrixAuthPDB covers the default PDB (minAvailable: 1).
func TestGoldenMatrixAuthPDB(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-pdb",
		Templates:      []string{"templates/matrix-authentication-pdb.yaml"},
	})
}

// TestGoldenMatrixAuthPDBMaxUnavailable verifies that maxUnavailable is rendered
// instead of minAvailable when explicitly set.
func TestGoldenMatrixAuthPDBMaxUnavailable(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-pdb-max-unavailable",
		Templates:      []string{"templates/matrix-authentication-pdb.yaml"},
		SetValues:      map[string]string{"matrixAuthentication.podDisruptionBudget.maxUnavailable": "1"},
	})
}

// TestGoldenMatrixAuthHPA covers the default state: autoscaling disabled → no HPA rendered.
func TestGoldenMatrixAuthHPA(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-hpa-disabled",
		Templates:      []string{"templates/matrix-authentication-hpa.yaml"},
		AllowEmpty:     true,
	})
}

// TestGoldenMatrixAuthHPAEnabled verifies HPA is rendered with CPU and memory targets.
func TestGoldenMatrixAuthHPAEnabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-hpa-enabled",
		Templates:      []string{"templates/matrix-authentication-hpa.yaml"},
		SetValues: map[string]string{
			"matrixAuthentication.autoscaling.enabled":                        "true",
			"matrixAuthentication.autoscaling.minReplicas":                    "2",
			"matrixAuthentication.autoscaling.maxReplicas":                    "5",
			"matrixAuthentication.autoscaling.targetCPUUtilizationPercentage": "80",
		},
	})
}

// TestGoldenMatrixAuthIngress covers the default ingress (no TLS secret).
func TestGoldenMatrixAuthIngress(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-ingress",
		Templates:      []string{"templates/matrix-authentication-ingress.yaml"},
	})
}

// TestGoldenMatrixAuthIngressWithTLS verifies secretName is rendered in the TLS block.
func TestGoldenMatrixAuthIngressWithTLS(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-ingress-tls",
		Templates:      []string{"templates/matrix-authentication-ingress.yaml"},
		SetValues:      map[string]string{"matrixAuthentication.ingress.secretName": "mas-tls"},
	})
}

// TestGoldenMatrixAuthJobConfigSync covers the default job with ArgoCD annotations.
func TestGoldenMatrixAuthJobConfigSync(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-job-config-sync",
		Templates:      []string{"templates/matrix-authentication-job-config-sync.yaml"},
	})
}

// TestGoldenMatrixAuthJobConfigSyncHelm verifies helm.sh/hook annotations are used
// when argocd mode is disabled.
func TestGoldenMatrixAuthJobConfigSyncHelm(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-job-config-sync-helm",
		Templates:      []string{"templates/matrix-authentication-job-config-sync.yaml"},
		SetValues:      map[string]string{"argocd": "false"},
	})
}

// TestGoldenMatrixAuthJobConfigSyncWithPrune verifies the --prune flag is added
// to the config sync command when configSyncPrune is enabled.
func TestGoldenMatrixAuthJobConfigSyncWithPrune(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-job-config-sync-prune",
		Templates:      []string{"templates/matrix-authentication-job-config-sync.yaml"},
		SetValues:      map[string]string{"matrixAuthentication.configSyncPrune": "true"},
	})
}

// TestGoldenMatrixAuthJobDbMigration covers the DB migration job with ArgoCD annotations.
func TestGoldenMatrixAuthJobDbMigration(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-job-db-migration",
		Templates:      []string{"templates/matrix-authentication-job-db-migration.yaml"},
	})
}

// TestGoldenMatrixAuthJobDbMigrationHelm verifies helm.sh/hook annotations when
// argocd mode is disabled.
func TestGoldenMatrixAuthJobDbMigrationHelm(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-job-db-migration-helm",
		Templates:      []string{"templates/matrix-authentication-job-db-migration.yaml"},
		SetValues:      map[string]string{"argocd": "false"},
	})
}
