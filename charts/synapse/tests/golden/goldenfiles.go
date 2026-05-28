package golden

import (
	"flag"
	"os"
	"regexp"
	"strings"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/suite"
)

var update = flag.Bool("update-golden", true, "update golden test output files")

type TemplateGoldenTest struct {
	suite.Suite
	ChartPath      string
	Release        string
	Namespace      string
	GoldenFileName string
	Templates      []string
	IgnoredLines   []string
	ValuesFiles    []string
	SetValues      map[string]string
	// AllowEmpty allows templates guarded by an `{{- if }}` block to render
	// nothing. Helm returns "could not find template" in that case, which is
	// treated as an empty string so the golden file captures the disabled state.
	AllowEmpty bool
}

func (s *TemplateGoldenTest) TestContainerGoldenTestDefaults() {
	options := &helm.Options{
		KubectlOptions: k8s.NewKubectlOptions("", "", s.Namespace),
		SetValues:      s.SetValues,
		ValuesFiles:    s.ValuesFiles,
	}

	var output string
	if s.AllowEmpty {
		var err error
		output, err = helm.RenderTemplateE(s.T(), options, s.ChartPath, s.Release, s.Templates)
		if err != nil {
			s.Require().True(
				strings.Contains(err.Error(), "could not find template"),
				"unexpected helm error: %v", err,
			)
			output = ""
		}
	} else {
		output = helm.RenderTemplate(s.T(), options, s.ChartPath, s.Release, s.Templates)
	}

	s.IgnoredLines = append(s.IgnoredLines, `\s+helm.sh/chart:\s+.*`)
	bytes := []byte(output)
	for _, ignoredLine := range s.IgnoredLines {
		regex := regexp.MustCompile(ignoredLine)
		bytes = regex.ReplaceAll(bytes, []byte(""))
	}
	output = string(bytes)

	goldenFile := "../fixtures/" + s.GoldenFileName + ".golden.yaml"

	if *update {
		err := os.WriteFile(goldenFile, bytes, 0644)
		s.Require().NoError(err, "Golden file was not writable")
	}

	expected, err := os.ReadFile(goldenFile)

	// then
	s.Require().NoError(err, "Golden file doesn't exist or was not readable")
	s.Require().Equal(string(expected), output)
}
