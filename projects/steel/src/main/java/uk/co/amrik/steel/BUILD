load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")
load("@rules_java//java:defs.bzl", "java_binary", "java_library")

package(default_visibility = ["//visibility:public"])

COMMON_DEPS = [
    "//projects/steel/src/main/java/uk/co/amrik/steel/order",
    "//projects/steel/src/main/java/uk/co/amrik/steel/user",
    "//projects/steel/src/main/java/uk/co/amrik/steel/persistence",
    "//projects/steel/src/main/java/uk/co/amrik/steel/permissions",
    "//projects/steel/src/main/java/uk/co/amrik/steel/middleware",
    "@maven//:com_google_inject_guice",
    "@maven//:com_google_inject_extensions_guice_servlet",
    "@maven//:org_jboss_resteasy_resteasy_client",
    "@maven//:org_jboss_resteasy_resteasy_jaxb_provider",
    "@maven//:jakarta_platform_jakarta_jakartaee_web_api",
    "@maven//:org_eclipse_jetty_jetty_util",
    "@maven//:org_glassfish_jersey_core_jersey_server",
    "@maven//:org_glassfish_jersey_containers_jersey_container_servlet_core",
    "@maven//:org_jboss_resteasy_resteasy_servlet_initializer",
    "@maven//:org_eclipse_jetty_jetty_server",
    "@maven//:org_eclipse_jetty_jetty_servlet",
]

java_binary(
    name = "ProjectSteel",
    srcs = [":srcs"],
    main_class = "uk.co.amrik.steel.Steel",
    deps = [":steel_lib"] + COMMON_DEPS,
)

java_library(
    name = "steel_lib",
    srcs = glob(["*.java"]),
    deps = COMMON_DEPS,
)

filegroup(
    name = "srcs",
    srcs = ["BUILD"] + glob(["**/*.java"]),
)
