load("@rules_java//java:defs.bzl", "java_binary", "java_library")

package(default_visibility = ["//visibility:public"])

COMMON_DEPS = [
    "@maven//:com_google_inject_guice",
    "@maven//:com_google_inject_extensions_guice_servlet",
    "@maven//:org_slf4j_slf4j_api",
    "@maven//:org_eclipse_jetty_jetty_server",
    "@maven//:org_eclipse_jetty_ee10_jetty_ee10_servlet",
    "@maven//:jakarta_servlet_jakarta_servlet_api",
]

RUNTIME_DEPS = [
    "@maven//:org_eclipse_jetty_jetty_slf4j_impl",
]

java_library(
    name = "health",
    srcs = glob(
        ["*.java"],
    ),
    deps = [
    ] + COMMON_DEPS,
)

filegroup(
    name = "srcs",
    srcs = ["BUILD"] + glob(["**/*.java"]),
)
