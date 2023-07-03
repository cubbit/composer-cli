load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "input",
    srcs = ["input.go"],
    importpath = "github.com/cubbit/cubbit/client/cli/src/input",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_alecaivazis_survey_v2//:survey",
        "@org_golang_x_term//:term",
    ],
)
