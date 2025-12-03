package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// CheckResult represents the result of a compatibility check
type CheckResult struct {
	K8sVersion string            `json:"k8sVersion" yaml:"k8sVersion"`
	Components []ComponentResult `json:"components" yaml:"components"`
	Summary    Summary           `json:"summary" yaml:"summary"`
}

// ComponentResult represents the check result for a single component
type ComponentResult struct {
	Name        string `json:"name" yaml:"name"`
	Current     string `json:"current,omitempty" yaml:"current,omitempty"`
	Required    string `json:"required" yaml:"required"`
	Recommended string `json:"recommended,omitempty" yaml:"recommended,omitempty"`
	Status      string `json:"status" yaml:"status"` // compatible, incompatible, unknown
	SkewPolicy  string `json:"skewPolicy,omitempty" yaml:"skewPolicy,omitempty"`
	Notes       string `json:"notes,omitempty" yaml:"notes,omitempty"`
}

// Summary provides an overview of the check results
type Summary struct {
	TotalComponents      int `json:"totalComponents" yaml:"totalComponents"`
	CompatibleComponents int `json:"compatibleComponents" yaml:"compatibleComponents"`
	IncompatibleCount    int `json:"incompatibleCount" yaml:"incompatibleCount"`
	UnknownCount         int `json:"unknownCount" yaml:"unknownCount"`
}

// Formatter interface for different output formats
type Formatter interface {
	Format(result *CheckResult) error
}

// TableFormatter outputs results as a table
type TableFormatter struct {
	Writer io.Writer
}

// JSONFormatter outputs results as JSON
type JSONFormatter struct {
	Writer io.Writer
}

// YAMLFormatter outputs results as YAML
type YAMLFormatter struct {
	Writer io.Writer
}

// NewFormatter creates a new formatter based on the format string
func NewFormatter(format string) Formatter {
	switch strings.ToLower(format) {
	case "json":
		return &JSONFormatter{Writer: os.Stdout}
	case "yaml":
		return &YAMLFormatter{Writer: os.Stdout}
	default:
		return &TableFormatter{Writer: os.Stdout}
	}
}

// Format outputs the result as a table
func (f *TableFormatter) Format(result *CheckResult) error {
	// Header
	fmt.Fprintf(f.Writer, "\n")
	fmt.Fprintf(f.Writer, "Kubernetes Version: %s\n", result.K8sVersion)
	fmt.Fprintf(f.Writer, "%s\n\n", strings.Repeat("=", 60))

	// Table header
	fmt.Fprintf(f.Writer, "%-28s %-15s %-15s %-12s\n", "COMPONENT", "REQUIRED", "RECOMMENDED", "STATUS")
	fmt.Fprintf(f.Writer, "%s\n", strings.Repeat("-", 70))

	// Component rows
	for _, c := range result.Components {
		status := formatStatus(c.Status)
		required := c.Required
		if required == "" {
			required = c.SkewPolicy
		}
		recommended := c.Recommended
		if recommended == "" {
			recommended = "-"
		}
		fmt.Fprintf(f.Writer, "%-28s %-15s %-15s %s\n", c.Name, required, recommended, status)
	}

	// Summary
	fmt.Fprintf(f.Writer, "\n%s\n", strings.Repeat("-", 70))
	fmt.Fprintf(f.Writer, "Summary: %d components checked\n", result.Summary.TotalComponents)
	if result.Summary.IncompatibleCount > 0 {
		fmt.Fprintf(f.Writer, "  ⚠️  %d incompatible\n", result.Summary.IncompatibleCount)
	}
	if result.Summary.UnknownCount > 0 {
		fmt.Fprintf(f.Writer, "  ❓ %d unknown\n", result.Summary.UnknownCount)
	}
	if result.Summary.CompatibleComponents == result.Summary.TotalComponents {
		fmt.Fprintf(f.Writer, "  ✅ All components compatible\n")
	}
	fmt.Fprintf(f.Writer, "\n")

	return nil
}

func formatStatus(status string) string {
	switch status {
	case "compatible":
		return "✅ Compatible"
	case "incompatible":
		return "❌ Incompatible"
	default:
		return "❓ Unknown"
	}
}

// Format outputs the result as JSON
func (f *JSONFormatter) Format(result *CheckResult) error {
	encoder := json.NewEncoder(f.Writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

// Format outputs the result as YAML
func (f *YAMLFormatter) Format(result *CheckResult) error {
	encoder := yaml.NewEncoder(f.Writer)
	encoder.SetIndent(2)
	return encoder.Encode(result)
}
