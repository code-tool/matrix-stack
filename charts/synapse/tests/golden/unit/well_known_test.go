package golden

import (
	"path/filepath"
	"testing"
	"tests/golden"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestGoldenWellKnownDefault covers the default well-known config:
// MAS authentication block present (msc2965 enabled by default),
// no identity_server, no RTC foci block.
func TestGoldenWellKnownDefault(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-well-known-default",
		Templates:      []string{"templates/well-known-configmap.yaml"},
	})
}

// TestGoldenWellKnownIdentityServer verifies that identity_server block
// is rendered when identity_server_vector is enabled.
func TestGoldenWellKnownIdentityServer(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-well-known-identity-server",
		Templates:      []string{"templates/well-known-configmap.yaml"},
		SetValues:      map[string]string{"identity_server_vector": "true"},
	})
}

// TestGoldenWellKnownMsc3266 verifies that RTC foci block (for Element Call via LiveKit)
// is rendered when msc3266 is enabled. Also tests the livekit ingress host substitution.
func TestGoldenWellKnownMsc3266(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../..")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "golden-file-test",
		Namespace:      "test-namespace",
		GoldenFileName: "test-well-known-msc3266",
		Templates:      []string{"templates/well-known-configmap.yaml"},
		SetValues: map[string]string{
			"experimentalFeatures.msc3266.enabled": "true",
			"livekitServer.ingress.host":           "livekit.example.com",
		},
	})
}
