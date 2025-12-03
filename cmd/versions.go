package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pmady/kube-dependency-checker/pkg/compatibility"
	"github.com/spf13/cobra"
)

var (
	componentName    string
	versionsK8sVer   string
	showAllVersions  bool
)

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List compatible versions for a component",
	Long: `List compatible versions for a specific component across Kubernetes versions.

Examples:
  # List etcd versions for Kubernetes 1.30
  kube-dependency-checker versions --component etcd --k8s-version 1.30

  # List CoreDNS versions across all supported Kubernetes versions
  kube-dependency-checker versions --component coredns --all

  # List all supported Kubernetes versions
  kube-dependency-checker versions --list-k8s`,
	RunE: runVersions,
}

var listK8s bool

func init() {
	rootCmd.AddCommand(versionsCmd)
	versionsCmd.Flags().StringVar(&componentName, "component", "", "Component name (etcd, coredns, containerd, etc.)")
	versionsCmd.Flags().StringVar(&versionsK8sVer, "k8s-version", "", "Kubernetes version")
	versionsCmd.Flags().BoolVar(&showAllVersions, "all", false, "Show versions for all supported Kubernetes versions")
	versionsCmd.Flags().BoolVar(&listK8s, "list-k8s", false, "List all supported Kubernetes versions")
}

func runVersions(cmd *cobra.Command, args []string) error {
	// List supported K8s versions
	if listK8s {
		versions := compatibility.GetSupportedVersions()
		sort.Sort(sort.Reverse(sort.StringSlice(versions)))
		fmt.Println("Supported Kubernetes versions:")
		for _, v := range versions {
			fmt.Printf("  - %s\n", v)
		}
		return nil
	}

	// Validate component flag
	if componentName == "" {
		return fmt.Errorf("--component flag is required")
	}

	componentName = strings.ToLower(componentName)

	// Show versions for all K8s versions
	if showAllVersions {
		return showAllK8sVersions(componentName)
	}

	// Show version for specific K8s version
	if versionsK8sVer == "" {
		return fmt.Errorf("either --k8s-version or --all flag is required")
	}

	versionsK8sVer = strings.TrimPrefix(versionsK8sVer, "v")

	info, ok := compatibility.GetComponentInfo(versionsK8sVer, componentName)
	if !ok {
		return fmt.Errorf("component '%s' not found for Kubernetes %s", componentName, versionsK8sVer)
	}

	fmt.Printf("\n%s compatibility for Kubernetes %s:\n", info.Name, versionsK8sVer)
	fmt.Println(strings.Repeat("-", 50))

	if info.Version != "" {
		fmt.Printf("  Version:     %s\n", info.Version)
	}
	if info.Recommended != "" {
		fmt.Printf("  Recommended: %s\n", info.Recommended)
	}
	if info.MinVersion != "" {
		fmt.Printf("  Min Version: %s\n", info.MinVersion)
	}
	if info.MaxVersion != "" {
		fmt.Printf("  Max Version: %s\n", info.MaxVersion)
	}
	if info.SkewPolicy != "" {
		fmt.Printf("  Skew Policy: %s\n", info.SkewPolicy)
	}
	if info.Notes != "" {
		fmt.Printf("  Notes:       %s\n", info.Notes)
	}
	fmt.Println()

	return nil
}

func showAllK8sVersions(component string) error {
	versions := compatibility.GetSupportedVersions()
	sort.Sort(sort.Reverse(sort.StringSlice(versions)))

	fmt.Printf("\n%s versions across Kubernetes releases:\n", strings.Title(component))
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("%-15s %-15s %-15s\n", "K8S VERSION", "VERSION", "RECOMMENDED")
	fmt.Println(strings.Repeat("-", 60))

	for _, k8sVer := range versions {
		info, ok := compatibility.GetComponentInfo(k8sVer, component)
		if !ok {
			continue
		}

		version := info.Version
		if version == "" && info.SkewPolicy != "" {
			version = "(skew policy)"
		}

		recommended := info.Recommended
		if recommended == "" {
			recommended = "-"
		}

		fmt.Printf("%-15s %-15s %-15s\n", k8sVer, version, recommended)
	}
	fmt.Println()

	return nil
}
