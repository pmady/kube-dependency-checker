package compatibility

// ComponentInfo holds version compatibility information for a component
type ComponentInfo struct {
	Name           string
	Version        string
	MinVersion     string
	MaxVersion     string
	Recommended    string
	SkewPolicy     string
	MaxMinorSkew   int
	CanBeNewer     bool // kubectl can be newer than API server
	Notes          string
}

// K8sVersionMatrix holds all component compatibility info for a K8s version
type K8sVersionMatrix struct {
	K8sVersion string
	Components map[string]ComponentInfo
}

// CompatibilityMatrix holds the full compatibility data
var CompatibilityMatrix = map[string]K8sVersionMatrix{
	"1.33": {
		K8sVersion: "1.33",
		Components: map[string]ComponentInfo{
			"etcd": {
				Name:        "etcd",
				Version:     "3.5.x",
				MinVersion:  "3.5.0",
				MaxVersion:  "3.5.99",
				Recommended: "3.5.15",
				Notes:       "etcd 3.5.x is required for Kubernetes 1.33",
			},
			"coredns": {
				Name:        "CoreDNS",
				Version:     "1.12.0",
				MinVersion:  "1.11.0",
				MaxVersion:  "1.12.99",
				Recommended: "1.12.0",
				Notes:       "Installed by kubeadm",
			},
			"containerd": {
				Name:        "containerd",
				Version:     "1.7.x",
				MinVersion:  "1.7.0",
				MaxVersion:  "2.0.99",
				Recommended: "1.7.22",
				Notes:       "containerd 1.7+ or 2.0+ supported",
			},
			"kubelet": {
				Name:         "kubelet",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-proxy": {
				Name:         "kube-proxy",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-controller-manager": {
				Name:         "kube-controller-manager",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kube-scheduler": {
				Name:         "kube-scheduler",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kubectl": {
				Name:         "kubectl",
				SkewPolicy:   "Within 1 minor version (older or newer)",
				MaxMinorSkew: 1,
				CanBeNewer:   true,
			},
		},
	},
	"1.32": {
		K8sVersion: "1.32",
		Components: map[string]ComponentInfo{
			"etcd": {
				Name:        "etcd",
				Version:     "3.5.x",
				MinVersion:  "3.5.0",
				MaxVersion:  "3.5.99",
				Recommended: "3.5.15",
				Notes:       "etcd 3.5.x is required for Kubernetes 1.32",
			},
			"coredns": {
				Name:        "CoreDNS",
				Version:     "1.11.3",
				MinVersion:  "1.10.0",
				MaxVersion:  "1.11.99",
				Recommended: "1.11.3",
				Notes:       "Installed by kubeadm",
			},
			"containerd": {
				Name:        "containerd",
				Version:     "1.7.x",
				MinVersion:  "1.6.0",
				MaxVersion:  "2.0.99",
				Recommended: "1.7.22",
				Notes:       "containerd 1.6+ supported",
			},
			"kubelet": {
				Name:         "kubelet",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-proxy": {
				Name:         "kube-proxy",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-controller-manager": {
				Name:         "kube-controller-manager",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kube-scheduler": {
				Name:         "kube-scheduler",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kubectl": {
				Name:         "kubectl",
				SkewPolicy:   "Within 1 minor version (older or newer)",
				MaxMinorSkew: 1,
				CanBeNewer:   true,
			},
		},
	},
	"1.31": {
		K8sVersion: "1.31",
		Components: map[string]ComponentInfo{
			"etcd": {
				Name:        "etcd",
				Version:     "3.5.x",
				MinVersion:  "3.5.0",
				MaxVersion:  "3.5.99",
				Recommended: "3.5.12",
				Notes:       "etcd 3.5.x is required for Kubernetes 1.31",
			},
			"coredns": {
				Name:        "CoreDNS",
				Version:     "1.11.3",
				MinVersion:  "1.10.0",
				MaxVersion:  "1.11.99",
				Recommended: "1.11.3",
				Notes:       "Installed by kubeadm",
			},
			"containerd": {
				Name:        "containerd",
				Version:     "1.7.x",
				MinVersion:  "1.6.0",
				MaxVersion:  "1.7.99",
				Recommended: "1.7.20",
				Notes:       "containerd 1.6+ supported",
			},
			"kubelet": {
				Name:         "kubelet",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-proxy": {
				Name:         "kube-proxy",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-controller-manager": {
				Name:         "kube-controller-manager",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kube-scheduler": {
				Name:         "kube-scheduler",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kubectl": {
				Name:         "kubectl",
				SkewPolicy:   "Within 1 minor version (older or newer)",
				MaxMinorSkew: 1,
				CanBeNewer:   true,
			},
		},
	},
	"1.30": {
		K8sVersion: "1.30",
		Components: map[string]ComponentInfo{
			"etcd": {
				Name:        "etcd",
				Version:     "3.5.x",
				MinVersion:  "3.5.0",
				MaxVersion:  "3.5.99",
				Recommended: "3.5.12",
				Notes:       "etcd 3.5.x is required for Kubernetes 1.30",
			},
			"coredns": {
				Name:        "CoreDNS",
				Version:     "1.11.1",
				MinVersion:  "1.10.0",
				MaxVersion:  "1.11.99",
				Recommended: "1.11.1",
				Notes:       "Installed by kubeadm",
			},
			"containerd": {
				Name:        "containerd",
				Version:     "1.7.x",
				MinVersion:  "1.6.0",
				MaxVersion:  "1.7.99",
				Recommended: "1.7.16",
				Notes:       "containerd 1.6+ supported",
			},
			"kubelet": {
				Name:         "kubelet",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-proxy": {
				Name:         "kube-proxy",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-controller-manager": {
				Name:         "kube-controller-manager",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kube-scheduler": {
				Name:         "kube-scheduler",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kubectl": {
				Name:         "kubectl",
				SkewPolicy:   "Within 1 minor version (older or newer)",
				MaxMinorSkew: 1,
				CanBeNewer:   true,
			},
		},
	},
	"1.29": {
		K8sVersion: "1.29",
		Components: map[string]ComponentInfo{
			"etcd": {
				Name:        "etcd",
				Version:     "3.5.x",
				MinVersion:  "3.5.0",
				MaxVersion:  "3.5.99",
				Recommended: "3.5.10",
				Notes:       "etcd 3.5.x is required for Kubernetes 1.29",
			},
			"coredns": {
				Name:        "CoreDNS",
				Version:     "1.11.1",
				MinVersion:  "1.9.0",
				MaxVersion:  "1.11.99",
				Recommended: "1.11.1",
				Notes:       "Installed by kubeadm",
			},
			"containerd": {
				Name:        "containerd",
				Version:     "1.7.x",
				MinVersion:  "1.6.0",
				MaxVersion:  "1.7.99",
				Recommended: "1.7.13",
				Notes:       "containerd 1.6+ supported",
			},
			"kubelet": {
				Name:         "kubelet",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-proxy": {
				Name:         "kube-proxy",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-controller-manager": {
				Name:         "kube-controller-manager",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kube-scheduler": {
				Name:         "kube-scheduler",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kubectl": {
				Name:         "kubectl",
				SkewPolicy:   "Within 1 minor version (older or newer)",
				MaxMinorSkew: 1,
				CanBeNewer:   true,
			},
		},
	},
	"1.28": {
		K8sVersion: "1.28",
		Components: map[string]ComponentInfo{
			"etcd": {
				Name:        "etcd",
				Version:     "3.5.x",
				MinVersion:  "3.5.0",
				MaxVersion:  "3.5.99",
				Recommended: "3.5.9",
				Notes:       "etcd 3.5.x is required for Kubernetes 1.28",
			},
			"coredns": {
				Name:        "CoreDNS",
				Version:     "1.10.1",
				MinVersion:  "1.9.0",
				MaxVersion:  "1.10.99",
				Recommended: "1.10.1",
				Notes:       "Installed by kubeadm",
			},
			"containerd": {
				Name:        "containerd",
				Version:     "1.7.x",
				MinVersion:  "1.6.0",
				MaxVersion:  "1.7.99",
				Recommended: "1.7.8",
				Notes:       "containerd 1.6+ supported",
			},
			"kubelet": {
				Name:         "kubelet",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-proxy": {
				Name:         "kube-proxy",
				SkewPolicy:   "Up to 3 minor versions older than kube-apiserver",
				MaxMinorSkew: 3,
				CanBeNewer:   false,
			},
			"kube-controller-manager": {
				Name:         "kube-controller-manager",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kube-scheduler": {
				Name:         "kube-scheduler",
				SkewPolicy:   "Up to 1 minor version older than kube-apiserver",
				MaxMinorSkew: 1,
				CanBeNewer:   false,
			},
			"kubectl": {
				Name:         "kubectl",
				SkewPolicy:   "Within 1 minor version (older or newer)",
				MaxMinorSkew: 1,
				CanBeNewer:   true,
			},
		},
	},
}

// GetMatrix returns the compatibility matrix for a given K8s version
func GetMatrix(k8sVersion string) (*K8sVersionMatrix, bool) {
	matrix, ok := CompatibilityMatrix[k8sVersion]
	if !ok {
		return nil, false
	}
	return &matrix, true
}

// GetSupportedVersions returns all supported K8s versions
func GetSupportedVersions() []string {
	versions := make([]string, 0, len(CompatibilityMatrix))
	for v := range CompatibilityMatrix {
		versions = append(versions, v)
	}
	return versions
}

// GetComponentInfo returns component info for a specific K8s version
func GetComponentInfo(k8sVersion, component string) (*ComponentInfo, bool) {
	matrix, ok := GetMatrix(k8sVersion)
	if !ok {
		return nil, false
	}
	info, ok := matrix.Components[component]
	if !ok {
		return nil, false
	}
	return &info, true
}
