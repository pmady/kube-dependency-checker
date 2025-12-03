# Kube Dependency Checker - Design Document

## Overview

`kube-dependency-checker` is a cross-platform CLI tool that checks Kubernetes component version compatibility. It helps cluster administrators understand:

1. **Current compatibility** - Are your current component versions compatible with your K8s version?
2. **Upgrade planning** - What component versions are required for a future K8s version?
3. **Version skew validation** - Does your setup comply with Kubernetes version skew policy?

## Supported Components

### Core Kubernetes Components (Version Skew Policy)
| Component | Skew Rule |
|-----------|-----------|
| kube-apiserver | HA clusters: within 1 minor version |
| kubelet | Up to 3 minor versions older than kube-apiserver |
| kube-proxy | Up to 3 minor versions older than kube-apiserver |
| kube-controller-manager | Up to 1 minor version older than kube-apiserver |
| kube-scheduler | Up to 1 minor version older than kube-apiserver |
| cloud-controller-manager | Up to 1 minor version older than kube-apiserver |
| kubectl | Within 1 minor version (older or newer) |

### Dependency Components
| Component | Source |
|-----------|--------|
| etcd | Kubernetes release notes / kubeadm defaults |
| CoreDNS | [CoreDNS-k8s_version.md](https://github.com/coredns/deployment/blob/master/kubernetes/CoreDNS-k8s_version.md) |
| containerd | Kubernetes release notes |
| CNI plugins | Kubernetes release notes |
| metrics-server | GitHub releases compatibility |
| ingress-nginx | GitHub releases compatibility |

## CLI Usage

```bash
# Check current cluster compatibility
kube-dependency-checker check --kubeconfig ~/.kube/config

# Check compatibility for a specific K8s version
kube-dependency-checker check --k8s-version 1.30

# Check upgrade path from current to target version
kube-dependency-checker upgrade --from 1.28 --to 1.30

# List compatible versions for a component
kube-dependency-checker versions --component etcd --k8s-version 1.30

# Output formats
kube-dependency-checker check --k8s-version 1.30 --output json
kube-dependency-checker check --k8s-version 1.30 --output yaml
kube-dependency-checker check --k8s-version 1.30 --output table
```

## Architecture

```
cmd/
├── root.go           # Root command
├── check.go          # Check compatibility command
├── upgrade.go        # Upgrade path command
├── versions.go       # List versions command
└── completion.go     # Shell completion

pkg/
├── compatibility/
│   ├── matrix.go     # Version compatibility matrix
│   ├── skew.go       # Version skew policy logic
│   └── data/         # Embedded compatibility data
├── kubernetes/
│   ├── client.go     # K8s client for cluster inspection
│   └── versions.go   # Version parsing utilities
├── components/
│   ├── etcd.go       # etcd version compatibility
│   ├── coredns.go    # CoreDNS version compatibility
│   ├── containerd.go # containerd compatibility
│   └── cni.go        # CNI plugins compatibility
└── output/
    ├── table.go      # Table output formatter
    ├── json.go       # JSON output formatter
    └── yaml.go       # YAML output formatter
```

## Data Sources

### Embedded Data (Offline Mode)
- Compatibility matrix embedded at build time
- Updated with each release
- Covers last 5 supported K8s versions

### Online Mode (Optional)
- Fetch latest compatibility data from GitHub
- Check for newer component versions
- Validate against official Kubernetes releases

## Version Compatibility Matrix

### CoreDNS to Kubernetes Mapping
| Kubernetes | CoreDNS |
|------------|---------|
| v1.33 | v1.12.0 |
| v1.32 | v1.11.3 |
| v1.31 | v1.11.3 |
| v1.30 | v1.11.1 |
| v1.29 | v1.11.1 |
| v1.28 | v1.10.1 |
| v1.27 | v1.10.1 |
| v1.26 | v1.9.3 |
| v1.25 | v1.9.3 |

### etcd to Kubernetes Mapping
| Kubernetes | etcd |
|------------|------|
| v1.33 | v3.5.x |
| v1.32 | v3.5.x |
| v1.31 | v3.5.x |
| v1.30 | v3.5.x |
| v1.29 | v3.5.x |
| v1.28 | v3.5.x |

## Build & Release

### Cross-Platform Builds
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

### Release Automation
- GoReleaser for builds
- GitHub Actions for CI/CD
- Homebrew tap for macOS
- APT/YUM repos for Linux

## Future Enhancements

1. **Helm chart compatibility** - Check Helm chart requirements
2. **Operator compatibility** - Check operator version requirements
3. **Cloud provider specifics** - EKS, GKE, AKS version mappings
4. **Interactive mode** - TUI for exploring compatibility
5. **CI/CD integration** - GitHub Action for pipeline checks
