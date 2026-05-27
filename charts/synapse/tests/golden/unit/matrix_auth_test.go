package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestGoldenMatrixAuthDeployment covers the MAS deployment with default values:
// correct image, probes on internal port, config checksum annotation, volume mount.
func TestGoldenMatrixAuthDeployment(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-deployment",
		Templates:      []string{"templates/matrix-authentication-deployment.yaml"},
	})
}

// TestGoldenMatrixAuthDeploymentDisabled verifies that the deployment renders
// nothing when matrixAuthentication is disabled (template is fully guarded by
// {{- if .Values.matrixAuthentication.enabled }}).
func TestGoldenMatrixAuthDeploymentDisabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-deployment-disabled",
		Templates:      []string{"templates/matrix-authentication-deployment.yaml"},
		SetValues:      map[string]string{"matrixAuthentication.enabled": "false"},
		AllowEmpty:     true,
	})
}

// TestGoldenMatrixAuthSecret covers the MAS config secret with default values:
// HTTP listeners, database config, matrix homeserver block, configYaml merged in.
func TestGoldenMatrixAuthSecret(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-secret",
		Templates:      []string{"templates/matrix-authentication-secret.yaml"},
	})
}
