load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(name = "gazelle")

go_library(
    name = "cli_lib",
    srcs = ["main.go"],
    importpath = "github.com/cubbit/cubbit/client/cli",
    visibility = ["//visibility:private"],
    deps = ["//src/cmd"],
)

go_binary(
    name = "cli",
    embed = [":cli_lib"],
    visibility = ["//visibility:public"],
)

gazelle(
    name = "gazelle-update-repos",
    command = "update-repos",
    extra_args = [
        "-from_file=go.mod",
        "-to_macro=third_party/go/repositories.bzl%go_repositories",
        "-prune=true",
        "-build_file_proto_mode=disable",
    ],
)
