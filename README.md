# Kube Dependency Checker

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![CI](https://github.com/pmady/kube-dependency-checker/actions/workflows/ci.yml/badge.svg)](https://github.com/pmady/kube-dependency-checker/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/pmady/kube-dependency-checker)](https://goreportcard.com/report/github.com/pmady/kube-dependency-checker)

A CLI tool to check Kubernetes component version compatibility and plan upgrades.

## Overview

Kube Dependency Checker helps cluster administrators:

- **Check compatibility** - Verify component versions are compatible with your K8s version
- **Plan upgrades** - See what component versions are required for a target K8s version
- **Validate version skew** - Ensure compliance with Kubernetes version skew policy

## Features

- Check compatibility for Kubernetes versions 1.28 - 1.33
- Support for core components (kubelet, kube-proxy, etc.) and dependencies (etcd, CoreDNS, containerd)
- Upgrade path recommendations with step-by-step guidance
- Multiple output formats (table, JSON, YAML)
- Cross-platform support (Linux, macOS, Windows)

## Installation

### Homebrew (macOS/Linux)

```bash
brew install pmady/tap/kube-dependency-checker
```

### Go Install

```bash
go install github.com/pmady/kube-dependency-checker@latest
```

### Binary Download

Download the latest release from [GitHub Releases](https://github.com/pmady/kube-dependency-checker/releases).

### From Source

```bash
git clone https://github.com/pmady/kube-dependency-checker.git
cd kube-dependency-checker
make build
```

## Usage

### Check Compatibility

```bash
# Check component compatibility for Kubernetes 1.30
kube-dependency-checker check --k8s-version 1.30

# Output as JSON
kube-dependency-checker check --k8s-version 1.30 -o json

# Output as YAML
kube-dependency-checker check --k8s-version 1.30 -o yaml
```

### Plan Upgrades

```bash
# Show upgrade path from 1.28 to 1.30
kube-dependency-checker upgrade --from 1.28 --to 1.30
```

### List Component Versions

```bash
# List etcd versions for Kubernetes 1.30
kube-dependency-checker versions --component etcd --k8s-version 1.30

# List CoreDNS versions across all supported Kubernetes versions
kube-dependency-checker versions --component coredns --all

# List all supported Kubernetes versions
kube-dependency-checker versions --list-k8s
```

## Example Output

```
$ kube-dependency-checker check --k8s-version 1.30

Kubernetes Version: 1.30
============================================================

COMPONENT                    REQUIRED        RECOMMENDED     STATUS
----------------------------------------------------------------------
etcd                         3.5.x           3.5.12          ✅ Compatible
CoreDNS                      1.11.1          1.11.1          ✅ Compatible
containerd                   1.7.x           1.7.16          ✅ Compatible
kubelet                      Up to 3 minor versions older    ✅ Compatible
kube-proxy                   Up to 3 minor versions older    ✅ Compatible
kube-controller-manager      Up to 1 minor version older     ✅ Compatible
kube-scheduler               Up to 1 minor version older     ✅ Compatible
kubectl                      Within 1 minor version          ✅ Compatible

----------------------------------------------------------------------
Summary: 8 components checked
  ✅ All components compatible
```

## Documentation

- [Design Document](docs/DESIGN.md)
- [Contributing Guide](CONTRIBUTING.md)
- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Security Policy](SECURITY.md)

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

All commits must be signed off (DCO). See the contributing guide for instructions.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## DCO

All contributors must sign off their commits using the Developer Certificate of Origin (DCO). See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Maintainers

See [MAINTAINERS.md](MAINTAINERS.md) for the list of maintainers.
