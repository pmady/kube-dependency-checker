package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pmady/kube-dependency-checker/pkg/compatibility"
	"github.com/pmady/kube-dependency-checker/pkg/output"
	"github.com/spf13/cobra"
)

var (
	k8sVersion string
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check component compatibility for a Kubernetes version",
	Long: `Check component compatibility for a specific Kubernetes version.

This command shows the required and recommended versions for all
components that are compatible with the specified Kubernetes version.

Examples:
  # Check compatibility for Kubernetes 1.30
  kube-dependency-checker check --k8s-version 1.30

  # Output as JSON
  kube-dependency-checker check --k8s-version 1.30 -o json

  # Output as YAML
  kube-dependency-checker check --k8s-version 1.30 -o yaml`,
	RunE: runCheck,
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVar(&k8sVersion, "k8s-version", "", "Kubernetes version to check (e.g., 1.30)")
	checkCmd.MarkFlagRequired("k8s-version")
}

func runCheck(cmd *cobra.Command, args []string) error {
	// Normalize version (remove 'v' prefix if present)
	k8sVersion = strings.TrimPrefix(k8sVersion, "v")

	// Get compatibility matrix for the specified version
	matrix, ok := compatibility.GetMatrix(k8sVersion)
	if !ok {
		supportedVersions := compatibility.GetSupportedVersions()
		sort.Sort(sort.Reverse(sort.StringSlice(supportedVersions)))
		return fmt.Errorf("unsupported Kubernetes version: %s\nSupported versions: %s",
			k8sVersion, strings.Join(supportedVersions, ", "))
	}

	// Build the result
	result := &output.CheckResult{
		K8sVersion: k8sVersion,
		Components: make([]output.ComponentResult, 0),
	}

	// Define component order for consistent output
	componentOrder := []string{
		"etcd",
		"coredns",
		"containerd",
		"kubelet",
		"kube-proxy",
		"kube-controller-manager",
		"kube-scheduler",
		"kubectl",
	}

	for _, compName := range componentOrder {
		info, exists := matrix.Components[compName]
		if !exists {
			continue
		}

		compResult := output.ComponentResult{
			Name:        info.Name,
			Required:    info.Version,
			Recommended: info.Recommended,
			Status:      "compatible", // Default to compatible when showing requirements
			SkewPolicy:  info.SkewPolicy,
			Notes:       info.Notes,
		}

		// For skew policy components, show the policy instead of version
		if info.SkewPolicy != "" {
			compResult.Required = ""
		}

		result.Components = append(result.Components, compResult)
	}

	// Calculate summary
	result.Summary = output.Summary{
		TotalComponents:      len(result.Components),
		CompatibleComponents: len(result.Components),
		IncompatibleCount:    0,
		UnknownCount:         0,
	}

	// Output the result
	formatter := output.NewFormatter(outputFormat)
	return formatter.Format(result)
}
