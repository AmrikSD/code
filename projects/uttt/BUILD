load("@rules_go//go:def.bzl", "go_binary", "go_library", )

go_library(
    name = "app_lib",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "app",
    embed = [":app_lib"],
    visibility = ["//visibility:public"],
    out = "uttt"
)