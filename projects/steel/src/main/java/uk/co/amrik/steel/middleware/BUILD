load("@rules_java//java:defs.bzl", "java_library")

package(default_visibility = ["//visibility:public"])

filegroup(
    name = "middleware_annotations",
    srcs = glob(["annotations/**/*.java"]),
)
filegroup(
    name = "middleware_filters",
    srcs = glob(["filters/**/*.java"]),
)


java_library(
    name = "middleware",
    srcs = [
        ":middleware_annotations",
        ":middleware_filters",
    ],
    deps = [
        "@maven//:jakarta_platform_jakarta_jakartaee_web_api",
    ]
)
