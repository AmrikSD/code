load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    importpath = "github.com/amriksd/code/projects/go-docs/cmd",
    visibility = ["//visibility:private"],
    deps = ["@com_github_labstack_echo_v4//:echo"],
)

go_binary(
    name = "cmd",
    data = ["//projects/go-docs:public_files"],
    embed = [":cmd_lib"],
    visibility = ["//visibility:public"],
)
