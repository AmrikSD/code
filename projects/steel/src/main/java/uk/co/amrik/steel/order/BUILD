load("@bazel_tools//tools/jdk:default_java_toolchain.bzl", "default_java_toolchain")
load("@rules_java//java:defs.bzl", "java_library")

default_java_toolchain(
    name = "my_java_toolchain",
    source_version = "21",
    target_version = "21",
)

package(default_visibility = ["//visibility:public"])

filegroup(
    name = "order_root",
    srcs = glob(["*.java"]),
)

filegroup(
    name = "order_api_v1",
    srcs = glob(["api/v1/**/*.java"]),
)

java_library(
    name = "order",
    srcs = [
        ":order_api_v1",
        ":order_root",
    ],
    deps = [
        "//projects/steel/src/main/java/uk/co/amrik/steel/API",
        "//projects/steel/src/main/java/uk/co/amrik/steel/persistence",
        "@maven//:com_google_inject_extensions_guice_servlet",
        "@maven//:com_google_inject_guice",
        "@maven//:jakarta_platform_jakarta_jakartaee_web_api",
        "@maven//:org_glassfish_jersey_containers_jersey_container_servlet_core",
        "@maven//:org_glassfish_jersey_core_jersey_server",
        "@maven//:org_glassfish_jersey_inject_jersey_hk2",
        "@maven//:org_glassfish_jersey_media_jersey_media_json_jackson",
        "@maven//:org_jooq_jooq",
    ],
)

filegroup(
    name = "order_test_srcs",
    srcs = glob(["**/*.java"]),
)

java_library(
    name = "order_test_lib",
    srcs = [":order_test_srcs"],
    deps = [
        "//projects/steel/src/main/java/uk/co/amrik/steel/API",
        "//projects/steel/src/main/java/uk/co/amrik/steel/order",
        "//projects/steel/src/main/java/uk/co/amrik/steel/persistence",
        "@maven//:com_google_inject_guice",
        "@maven//:org_assertj_assertj_core",
        "@maven//:org_jooq_jooq",
        "@maven//:org_junit_jupiter_junit_jupiter_api",
        "@maven//:org_junit_jupiter_junit_jupiter_engine",
    ],
)
