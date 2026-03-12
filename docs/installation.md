# Installation

This guide covers all available methods to install the Cubbit Composer CLI.

---

## Prerequisites

- **Pre-built binaries**: no dependencies required
- **Go install / build from source**: Go 1.24+ and Git

---

## Option 1: Pre-built Binaries (Recommended)

Download the latest release for your platform from the [releases page](https://github.com/cubbit/composer-cli/releases).

**Linux / macOS:**

```bash
# After downloading, make it executable and move to PATH
chmod +x cubbit
sudo mv cubbit /usr/local/bin/
```

**Windows:**

Download `cubbit.exe` and copy it to a directory in your `PATH`.

Available builds: `linux_amd64`, `linux_arm64`, `linux_armv6`, `linux_armv7`, `linux_386`, `darwin_amd64`, `darwin_arm64`, `windows_amd64`, `windows_386`.

---

## Option 2: Go Install

```bash
go install github.com/cubbit/composer-cli@latest
```

This compiles and installs the `cubbit` binary into `$GOPATH/bin`. Make sure that directory is in your system `PATH`.

---

## Option 3: Build from Source

```bash
git clone https://github.com/cubbit/composer-cli.git
cd composer-cli

# Build for your current platform
go build -o build/cubbit .

# Install to PATH (Linux/macOS)
sudo cp build/cubbit /usr/local/bin/
```

**Cross-compilation:**

```bash
# macOS (Intel)
env GOOS=darwin GOARCH=amd64 go build -o build/cubbit .

# macOS (Apple Silicon)
env GOOS=darwin GOARCH=arm64 go build -o build/cubbit .

# Windows
env GOOS=windows GOARCH=amd64 go build -o build/cubbit.exe .

# Linux ARM64
env GOOS=linux GOARCH=arm64 go build -o build/cubbit .
```

---

## Option 4: Build with Bazel

For developers using the Bazel build system:

```bash
git clone https://github.com/cubbit/composer-cli.git
cd composer-cli

# Build for current platform
bazel build //:cli

# Platform-specific builds
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:cli
bazel build --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64 //:cli
bazel build --platforms=@io_bazel_rules_go//go/toolchain:windows_amd64 //:cli

# Run directly without a separate build step
bazel run //:cli -- --help
bazel run //:cli -- --version

# Install the built binary
cp bazel-bin/cli_/cli /usr/local/bin/cubbit
```

Visit [bazel.build/install](https://bazel.build/install) for Bazel installation instructions if needed.

---

## Verify Installation

```bash
cubbit --version
```

You should see the installed version printed to your terminal.

---

## Next Steps

- [Getting Started](getting-started.md) — first login and initial setup
- [Configuration](configuration.md) — managing profiles and environments
