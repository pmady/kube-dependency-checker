package compatibility

import (
	"testing"
)

func TestGetMatrix(t *testing.T) {
	tests := []struct {
		name       string
		k8sVersion string
		wantOk     bool
	}{
		{"valid version 1.30", "1.30", true},
		{"valid version 1.28", "1.28", true},
		{"valid version 1.33", "1.33", true},
		{"invalid version", "1.20", false},
		{"invalid format", "invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := GetMatrix(tt.k8sVersion)
			if ok != tt.wantOk {
				t.Errorf("GetMatrix() ok = %v, want %v", ok, tt.wantOk)
			}
		})
	}
}

func TestGetComponentInfo(t *testing.T) {
	tests := []struct {
		name       string
		k8sVersion string
		component  string
		wantOk     bool
	}{
		{"etcd for 1.30", "1.30", "etcd", true},
		{"coredns for 1.30", "1.30", "coredns", true},
		{"kubelet for 1.30", "1.30", "kubelet", true},
		{"invalid component", "1.30", "invalid", false},
		{"invalid version", "1.20", "etcd", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, ok := GetComponentInfo(tt.k8sVersion, tt.component)
			if ok != tt.wantOk {
				t.Errorf("GetComponentInfo() ok = %v, want %v", ok, tt.wantOk)
			}
			if ok && info == nil {
				t.Error("GetComponentInfo() returned nil info with ok=true")
			}
		})
	}
}

func TestGetSupportedVersions(t *testing.T) {
	versions := GetSupportedVersions()
	if len(versions) == 0 {
		t.Error("GetSupportedVersions() returned empty slice")
	}

	// Check that expected versions are present
	expectedVersions := []string{"1.28", "1.29", "1.30", "1.31", "1.32", "1.33"}
	for _, expected := range expectedVersions {
		found := false
		for _, v := range versions {
			if v == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("GetSupportedVersions() missing expected version %s", expected)
		}
	}
}

func TestComponentInfoFields(t *testing.T) {
	// Test that etcd has expected fields
	info, ok := GetComponentInfo("1.30", "etcd")
	if !ok {
		t.Fatal("Failed to get etcd info for 1.30")
	}

	if info.Name != "etcd" {
		t.Errorf("etcd Name = %s, want etcd", info.Name)
	}
	if info.Version == "" {
		t.Error("etcd Version is empty")
	}
	if info.Recommended == "" {
		t.Error("etcd Recommended is empty")
	}

	// Test that kubelet has skew policy
	kubeletInfo, ok := GetComponentInfo("1.30", "kubelet")
	if !ok {
		t.Fatal("Failed to get kubelet info for 1.30")
	}

	if kubeletInfo.SkewPolicy == "" {
		t.Error("kubelet SkewPolicy is empty")
	}
	if kubeletInfo.MaxMinorSkew != 3 {
		t.Errorf("kubelet MaxMinorSkew = %d, want 3", kubeletInfo.MaxMinorSkew)
	}
}
