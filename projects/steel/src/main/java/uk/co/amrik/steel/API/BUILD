load("@rules_java//java:defs.bzl", "java_library")

package(default_visibility = ["//visibility:public"])

filegroup(
    name = "api_base",
    srcs = glob(["**/*.java"]),
)


java_library(
    name = "API",
    srcs = [
        ":api_base",
    ],
    deps = [
        "@maven//:jakarta_platform_jakarta_jakartaee_web_api",
    ]
)
