package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestGoldenSynapseHTTPRouteDisabled verifies nothing is rendered when the
// Gateway API HTTPRoute is disabled (the default).
func TestGoldenSynapseHTTPRouteDisabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-httproute-disabled",
		Templates:      []string{"templates/synapse-httproute.yaml"},
		AllowEmpty:     true,
	})
}

// TestGoldenSynapseHTTPRoute covers the HTTPRoute that routes all traffic to
// synapse-client-reader-envoy (path: /, pathType: PathPrefix).
func TestGoldenSynapseHTTPRoute(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-synapse-httproute",
		Templates:      []string{"templates/synapse-httproute.yaml"},
		SetValues: map[string]string{
			"gatewayApi.enabled":                 "true",
			"gatewayApi.parentRefs[0].name":      "matrix-gateway",
			"gatewayApi.parentRefs[0].namespace": "gateway-system",
		},
	})
}

// TestGoldenAdminHTTPRouteDisabled verifies nothing is rendered when the
// Gateway API HTTPRoute is disabled (the default).
func TestGoldenAdminHTTPRouteDisabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-admin-httproute-disabled",
		Templates:      []string{"templates/admin-httproute.yaml"},
		AllowEmpty:     true,
	})
}

// TestGoldenAdminHTTPRoute covers the HTTPRoute that routes to the admin service.
func TestGoldenAdminHTTPRoute(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-admin-httproute",
		Templates:      []string{"templates/admin-httproute.yaml"},
		SetValues: map[string]string{
			"admin.gatewayApi.enabled":                 "true",
			"admin.gatewayApi.parentRefs[0].name":      "matrix-gateway",
			"admin.gatewayApi.parentRefs[0].namespace": "gateway-system",
		},
	})
}

// TestGoldenMatrixAuthHTTPRouteDisabled verifies nothing is rendered when the
// Gateway API HTTPRoute is disabled (the default).
func TestGoldenMatrixAuthHTTPRouteDisabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-httproute-disabled",
		Templates:      []string{"templates/matrix-authentication-httproute.yaml"},
		AllowEmpty:     true,
	})
}

// TestGoldenMatrixAuthHTTPRoute covers the HTTPRoute that routes to the
// matrix-authentication service.
func TestGoldenMatrixAuthHTTPRoute(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-matrix-auth-httproute",
		Templates:      []string{"templates/matrix-authentication-httproute.yaml"},
		SetValues: map[string]string{
			"matrixAuthentication.gatewayApi.enabled":                 "true",
			"matrixAuthentication.gatewayApi.parentRefs[0].name":      "matrix-gateway",
			"matrixAuthentication.gatewayApi.parentRefs[0].namespace": "gateway-system",
		},
	})
}

// TestGoldenWellKnownHTTPRouteDisabled verifies nothing is rendered when the
// Gateway API HTTPRoute is disabled (the default).
func TestGoldenWellKnownHTTPRouteDisabled(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-well-known-httproute-disabled",
		Templates:      []string{"templates/well-known-httproute.yaml"},
		AllowEmpty:     true,
	})
}

// TestGoldenWellKnownHTTPRoute covers the HTTPRoute that routes the
// /.well-known/matrix path to the well-known service.
func TestGoldenWellKnownHTTPRoute(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-well-known-httproute",
		Templates:      []string{"templates/well-known-httproute.yaml"},
		SetValues: map[string]string{
			"wellKnown.gatewayApi.enabled":                 "true",
			"wellKnown.gatewayApi.parentRefs[0].name":      "matrix-gateway",
			"wellKnown.gatewayApi.parentRefs[0].namespace": "gateway-system",
		},
	})
}
