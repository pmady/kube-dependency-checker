package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pmady/kube-dependency-checker/pkg/compatibility"
	"github.com/spf13/cobra"
)

var (
	fromVersion string
	toVersion   string
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Show upgrade path between Kubernetes versions",
	Long: `Show the upgrade path and component version changes between two Kubernetes versions.

This command helps you plan upgrades by showing:
- The recommended upgrade path (one minor version at a time)
- Component version changes at each step
- Any breaking changes or deprecations

Examples:
  # Show upgrade path from 1.28 to 1.30
  kube-dependency-checker upgrade --from 1.28 --to 1.30

  # Show upgrade path from 1.29 to 1.32
  kube-dependency-checker upgrade --from 1.29 --to 1.32`,
	RunE: runUpgrade,
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
	upgradeCmd.Flags().StringVar(&fromVersion, "from", "", "Starting Kubernetes version")
	upgradeCmd.Flags().StringVar(&toVersion, "to", "", "Target Kubernetes version")
	upgradeCmd.MarkFlagRequired("from")
	upgradeCmd.MarkFlagRequired("to")
}

func runUpgrade(cmd *cobra.Command, args []string) error {
	fromVersion = strings.TrimPrefix(fromVersion, "v")
	toVersion = strings.TrimPrefix(toVersion, "v")

	// Validate versions
	fromMatrix, ok := compatibility.GetMatrix(fromVersion)
	if !ok {
		return fmt.Errorf("unsupported source version: %s", fromVersion)
	}

	toMatrix, ok := compatibility.GetMatrix(toVersion)
	if !ok {
		return fmt.Errorf("unsupported target version: %s", toVersion)
	}

	// Parse minor versions
	fromMinor := parseMinorVersion(fromVersion)
	toMinor := parseMinorVersion(toVersion)

	if fromMinor >= toMinor {
		return fmt.Errorf("target version must be newer than source version")
	}

	// Generate upgrade path
	fmt.Printf("\n")
	fmt.Printf("Upgrade Path: %s → %s\n", fromVersion, toVersion)
	fmt.Printf("%s\n\n", strings.Repeat("=", 60))

	// Show step-by-step upgrade path
	fmt.Println("📋 Recommended Upgrade Steps:")
	fmt.Println(strings.Repeat("-", 60))

	steps := toMinor - fromMinor
	for i := 0; i < steps; i++ {
		stepFrom := fmt.Sprintf("1.%d", fromMinor+i)
		stepTo := fmt.Sprintf("1.%d", fromMinor+i+1)
		fmt.Printf("  Step %d: %s → %s\n", i+1, stepFrom, stepTo)
	}

	fmt.Println()
	fmt.Println("⚠️  Note: Kubernetes supports upgrading one minor version at a time.")
	fmt.Println()

	// Show component changes
	fmt.Println("📦 Component Version Changes:")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("%-25s %-15s %-15s\n", "COMPONENT", fromVersion, toVersion)
	fmt.Println(strings.Repeat("-", 60))

	// Components to compare
	components := []string{"etcd", "coredns", "containerd"}

	for _, comp := range components {
		fromInfo := fromMatrix.Components[comp]
		toInfo := toMatrix.Components[comp]

		fromVer := fromInfo.Recommended
		if fromVer == "" {
			fromVer = fromInfo.Version
		}

		toVer := toInfo.Recommended
		if toVer == "" {
			toVer = toInfo.Version
		}

		change := ""
		if fromVer != toVer {
			change = " ⬆️"
		}

		fmt.Printf("%-25s %-15s %-15s%s\n", fromInfo.Name, fromVer, toVer, change)
	}

	fmt.Println()

	// Show skew policy reminders
	fmt.Println("📌 Version Skew Policy Reminders:")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("  • Upgrade kube-apiserver first")
	fmt.Println("  • Then upgrade kube-controller-manager, kube-scheduler")
	fmt.Println("  • Finally upgrade kubelet on all nodes")
	fmt.Println("  • kubelet can be up to 3 minor versions older than kube-apiserver")
	fmt.Println()

	return nil
}

func parseMinorVersion(version string) int {
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		minor, _ := strconv.Atoi(parts[1])
		return minor
	}
	return 0
}

// Helper to get sorted versions
func getSortedVersions() []string {
	versions := compatibility.GetSupportedVersions()
	sort.Sort(sort.Reverse(sort.StringSlice(versions)))
	return versions
}
