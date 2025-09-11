# Contributing to Cubbit CLI

This document outlines the contribution guidelines for the Cubbit CLI project. By contributing, you agree to abide by these guidelines and help maintain the quality and consistency of our codebase.

## Table of Contents

- [Development Setup](#development-setup)
- [Git Workflow](#git-workflow)
- [Code Style Guidelines](#code-style-guidelines)
- [Pull Request Process](#pull-request-process)
- [Issue Reporting](#issue-reporting)
- [Testing](#testing)
- [Code of Conduct](#code-of-conduct)

## Development Setup

### Prerequisites

- **Go**: Version 1.24.1 or later (latest stable version recommended)
- **Git**: Version 2.0 or later
- **Access**: Cubbit account for testing (see setup instructions below)

### Setup Instructions

#### For Internal Contributors

1. **Set up the repository** (follow internal setup procedures)

2. **Configure development environment** (use internal documentation)

3. **Set up CLI configuration**
   ```bash
   # Create .env.local with your endpoints
   cat > .env.local << EOF
   iam: https://your-iam-endpoint
   ch: https://your-ch-endpoint
   dash: https://your-dash-endpoint
   EOF
   ```

#### For External Contributors

1. **Fork and clone**
   ```bash
   git clone https://github.com/cubbit/composer-cli.git
   cd composer-cli
   ```

2. **Use your Cubbit account** for testing and development
   - Configure the CLI with your existing Cubbit credentials
   - Test against the production environment

### Build and Test

```bash
# Build locally
go build -o build/cubbit .

# Verify installation
./build/cubbit --version
./build/cubbit --help
```

## Git Workflow

### Commit Guidelines

- Use clear, descriptive commit messages
- Follow conventional commit format
- Keep commits focused on single changes
- Squash related commits before submitting PR

## Code Style Guidelines

### Go Standards

This project follows Go best practices and official Go Code Review Comments:

- **Package naming**: Lowercase, descriptive
- **Function naming**: PascalCase for exported, camelCase for unexported
- **Variable naming**: Descriptive names, avoid abbreviations
- **Error handling**: Always check errors, wrap with context

### Required Patterns

```go
// Error handling
if err != nil {
    return fmt.Errorf("%s: %w", constants.ErrorContext, err)
}

// Configuration loading
conf, err := configuration.LoadConfig()
if err != nil {
    return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
}

// Flag retrieval
if value, err = cmd.Flags().GetString("flag-name"); err != nil {
    return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
}
```

### CLI Architecture

- **Commands**: Use Cobra framework following established patterns
- **Flags**: Use persistent flags for common options
- **Output**: Support multiple formats (human, json, yaml, xml)
- **Errors**: Use constants from `constants/errors.go`

### Project Structure

- **Packages**: Logical separation of concerns
- **Files**: Descriptive names indicating functionality
- **Constants**: Centralized in `constants/` package
- **Utilities**: Organized in `utils/` package with single responsibility

## Pull Request Process

### Pre-Submission Checklist

- [ ] Code follows style guidelines
- [ ] Comprehensive error handling implemented
- [ ] Documentation updated (README, help text, comments)
- [ ] Testing with valid access
- [ ] Multiple output formats verified
- [ ] Branch rebased from latest develop

### PR Requirements

1. **Use appropriate template** based on change type (bugfix, feature, chore)
2. **Provide clear description** of changes and motivation
3. **Include testing details** and verification steps
4. **Update documentation** as needed

### Review Process

- Reviews are handled by maintainers as time permits
- Complex changes may require additional discussion
- Approval requires at least one maintainer approval

## Issue Reporting

### Bug Reports

Include the following information:

1. **Environment**
   - OS and version
   - Go version
   - CLI version (`cubbit --version`)

2. **Reproduction Steps**
   - Exact command sequence
   - Expected vs actual behavior
   - Complete error messages

3. **Context**
   - What you were trying to achieve
   - Recent changes to environment/config

### Feature Requests

Provide these details:

1. **Use Case**
   - Problem description
   - Expected benefits

2. **Proposed Solution**
   - Implementation approach

3. **Priority**
   - Impact on workflow
   - Estimated user adoption

## Testing

### Testing Approach

Due to infrastructure requirements, testing relies on:

1. **Real Environment Testing**
   - Internal contributors: Use development endpoints
   - External contributors: Use production with your account

2. **Validation Requirements**
   - Multi-format output compatibility
   - Error handling verification
   - Edge case testing

### Testing Checklist

```bash
# Basic functionality
./build/cubbit --help
./build/cubbit --version

# Output formats
./build/cubbit tenant list --output json
./build/cubbit tenant list --output yaml
./build/cubbit tenant list --output xml

# Error conditions
./build/cubbit tenant describe --id nonexistent
```

## Support

For assistance:

1. Check existing issues for similar questions
2. Review codebase for established patterns
3. Create issue with `question` label
4. Contact maintainers for complex technical questions

## Code of Conduct

This project adheres to our [Code of Conduct](CODE_OF_CONDUCT.md). All participants are expected to uphold these standards.

---

Thank you for contributing to Cubbit CLI! Your contributions help improve the tool for the entire Cubbit community.
