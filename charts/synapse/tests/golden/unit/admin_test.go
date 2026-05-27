package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestGoldenAdminDeployment(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-admin-deployment",
		Templates:      []string{"templates/admin-deployment.yaml"},
	})
}

func TestGoldenAdminService(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-admin-service",
		Templates:      []string{"templates/admin-service.yaml"},
	})
}

// TestGoldenAdminIngress covers the default ingress (enabled, no TLS secret).
func TestGoldenAdminIngress(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-admin-ingress",
		Templates:      []string{"templates/admin-ingress.yaml"},
	})
}

// TestGoldenAdminIngressWithTLS verifies secretName is rendered in the TLS block.
func TestGoldenAdminIngressWithTLS(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-admin-ingress-tls",
		Templates:      []string{"templates/admin-ingress.yaml"},
		SetValues:      map[string]string{"admin.ingress.secretName": "admin-tls"},
	})
}

// TestGoldenAdminIngressDisabled verifies nothing is rendered when ingress is disabled.
func TestGoldenAdminIngressDisabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-admin-ingress-disabled",
		Templates:      []string{"templates/admin-ingress.yaml"},
		SetValues:      map[string]string{"admin.ingress.enabled": "false"},
		AllowEmpty:     true,
	})
}
