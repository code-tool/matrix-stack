package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestGoldenWellKnownService(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-well-known-service",
		Templates:      []string{"templates/well-known-service.yaml"},
	})
}

func TestGoldenWellKnownIngress(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-well-known-ingress",
		Templates:      []string{"templates/well-known-ingress.yaml"},
	})
}

// TestGoldenWellKnownIngressWithTLS verifies secretName is rendered in the TLS block.
func TestGoldenWellKnownIngressWithTLS(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-well-known-ingress-tls",
		Templates:      []string{"templates/well-known-ingress.yaml"},
		SetValues:      map[string]string{"ingress.secretName": "synapse-tls"},
	})
}

// TestGoldenWellKnownDeployment covers the default deployment with openid-configuration
// volume mount (openid_configuration: true by default).
func TestGoldenWellKnownDeployment(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-well-known-deployment",
		Templates:      []string{"templates/well-known-deployment.yaml"},
	})
}

// TestGoldenWellKnownDeploymentNoOpenID verifies that the openid-configuration
// volume mount is absent when openid_configuration is disabled.
func TestGoldenWellKnownDeploymentNoOpenID(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-well-known-deployment-no-openid",
		Templates:      []string{"templates/well-known-deployment.yaml"},
		SetValues:      map[string]string{"openid_configuration": "false"},
	})
}
