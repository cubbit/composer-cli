load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "actions",
    srcs = ["actions.go"],
    importpath = "github.com/cubbit/cubbit/client/cli/src/actions",
    visibility = ["//visibility:public"],
    deps = [
        "//client/cli/src/api",
        "//client/cli/src/configuration",
        "//client/cli/src/input",
        "@com_github_urfave_cli_v2//:cli",
    ],
)
