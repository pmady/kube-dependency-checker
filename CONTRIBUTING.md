# Contributing to Kube Dependency Checker

Thank you for your interest in contributing to Kube Dependency Checker! This document provides guidelines and information about contributing to this project.

## Developer Certificate of Origin (DCO)

This project uses the [Developer Certificate of Origin (DCO)](https://developercertificate.org/) to ensure that contributors have the right to submit their contributions.

By contributing to this project, you agree to the DCO, which states that you have the right to submit your contribution and that you agree to license it under the project's open source license.

### Signing Your Commits

All commits must be signed off with your real name and email address:

```bash
git commit -s -m "Your commit message"
```

This adds a `Signed-off-by` line to your commit message:

```
Signed-off-by: Your Name <your.email@example.com>
```

### Signing Past Commits

If you forgot to sign off on a commit, you can amend it:

```bash
# For the last commit
git commit --amend -s

# For multiple commits, use interactive rebase
git rebase -i HEAD~n  # where n is the number of commits
# Then mark commits as 'edit' and run: git commit --amend -s
```

## Code of Conduct

This project follows the [CNCF Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## How to Contribute

### Reporting Issues

Before creating an issue, please check if a similar issue already exists. When creating an issue, please include:

- A clear and descriptive title
- A detailed description of the problem
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Environment details (OS, Kubernetes version, etc.)

### Submitting Pull Requests

1. **Fork the repository** and create your branch from `main`.
2. **Make your changes** following the coding guidelines below.
3. **Add tests** for any new functionality.
4. **Ensure all tests pass** before submitting.
5. **Update documentation** if needed.
6. **Submit a pull request** with a clear description of your changes.

### Pull Request Guidelines

- Keep PRs focused on a single change
- Write clear, descriptive commit messages
- Reference any related issues in your PR description
- Ensure CI checks pass
- Request review from maintainers

## Development Setup

```bash
# Clone the repository
git clone https://github.com/pmady/kube-dependency-checker.git
cd kube-dependency-checker

# Install dependencies
# (Instructions will be added as the project develops)
```

## Coding Guidelines

- Follow the existing code style
- Write clear, self-documenting code
- Add comments for complex logic
- Keep functions small and focused
- Write unit tests for new functionality

## Commit Message Guidelines

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

Types:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

## Review Process

1. All submissions require review by at least one maintainer
2. Reviewers may request changes or ask questions
3. Once approved, a maintainer will merge the PR

## Getting Help

- Open an issue for bugs or feature requests
- Join our community channels for discussions
- Check existing documentation and issues first

## Recognition

Contributors will be recognized in our [CONTRIBUTORS](CONTRIBUTORS) file.

Thank you for contributing to Kube Dependency Checker!
