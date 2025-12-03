package version

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Version represents a semantic version
type Version struct {
	Major int
	Minor int
	Patch int
	Raw   string
}

// Parse parses a version string into a Version struct
// Accepts formats: "1.30", "v1.30", "1.30.0", "v1.30.0"
func Parse(s string) (*Version, error) {
	s = strings.TrimPrefix(s, "v")
	s = strings.TrimSpace(s)

	// Match version pattern
	re := regexp.MustCompile(`^(\d+)\.(\d+)(?:\.(\d+))?$`)
	matches := re.FindStringSubmatch(s)
	if matches == nil {
		return nil, fmt.Errorf("invalid version format: %s", s)
	}

	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch := 0
	if matches[3] != "" {
		patch, _ = strconv.Atoi(matches[3])
	}

	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
		Raw:   s,
	}, nil
}

// String returns the version as a string
func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// ShortString returns the version without patch
func (v *Version) ShortString() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

// Compare compares two versions
// Returns -1 if v < other, 0 if v == other, 1 if v > other
func (v *Version) Compare(other *Version) int {
	if v.Major != other.Major {
		if v.Major < other.Major {
			return -1
		}
		return 1
	}
	if v.Minor != other.Minor {
		if v.Minor < other.Minor {
			return -1
		}
		return 1
	}
	if v.Patch != other.Patch {
		if v.Patch < other.Patch {
			return -1
		}
		return 1
	}
	return 0
}

// MinorDiff returns the difference in minor versions
func (v *Version) MinorDiff(other *Version) int {
	diff := v.Minor - other.Minor
	if diff < 0 {
		return -diff
	}
	return diff
}

// IsNewerThan returns true if v is newer than other
func (v *Version) IsNewerThan(other *Version) bool {
	return v.Compare(other) > 0
}

// IsOlderThan returns true if v is older than other
func (v *Version) IsOlderThan(other *Version) bool {
	return v.Compare(other) < 0
}

// IsCompatibleWithAPIServer checks if this version is compatible with the given API server version
// based on Kubernetes version skew policy
func (v *Version) IsCompatibleWithAPIServer(apiServer *Version, component string, maxSkew int) bool {
	// Component must not be newer than API server
	if v.IsNewerThan(apiServer) {
		return false
	}

	// Check minor version difference
	diff := apiServer.Minor - v.Minor
	return diff <= maxSkew
}
