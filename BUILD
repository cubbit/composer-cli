load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//coordinator:bazel/service.bzl", "go_service")

go_library(
    name = "cli_lib",
    srcs = ["main.go"],
    importpath = "github.com/cubbit/cubbit/client/cli",
    visibility = ["//visibility:private"],
    deps = [
        "//client/cli/src/actions",
        "@com_github_urfave_cli_v2//:cli",
    ],
)

go_service(
    name = "cli",
    library = ":cli_lib",
)
