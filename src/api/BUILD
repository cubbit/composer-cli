load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "api",
    srcs = [
        "authentication_client.go",
        "response.go",
        "response_extractors.go",
    ],
    importpath = "github.com/cubbit/cubbit/client/cli/src/api",
    visibility = ["//visibility:public"],
    deps = [
        "//client/cli/src/request_utils",
        "//client/cli/utils",
    ],
)
