package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Version information set at build time
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"

	// Global flags
	outputFormat string
)

var rootCmd = &cobra.Command{
	Use:   "kube-dependency-checker",
	Short: "Check Kubernetes component version compatibility",
	Long: `kube-dependency-checker is a CLI tool that helps you verify
Kubernetes component version compatibility.

It checks:
- Version skew policy compliance for core components
- Dependency component versions (etcd, CoreDNS, etc.)
- Upgrade path recommendations

Examples:
  # Check compatibility for Kubernetes 1.30
  kube-dependency-checker check --k8s-version 1.30

  # Show upgrade path from 1.28 to 1.30
  kube-dependency-checker upgrade --from 1.28 --to 1.30

  # List compatible etcd versions for Kubernetes 1.30
  kube-dependency-checker versions --component etcd --k8s-version 1.30`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Output format (table, json, yaml)")
}

