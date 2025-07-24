load("@rules_java//java:defs.bzl", "java_binary", "java_library")

package(default_visibility = ["//visibility:public"])

COMMON_DEPS = [
    "//projects/Rolodex/src/main/java/uk/co/amrik/rolodex/services",
    "//projects/Rolodex/src/main/java/uk/co/amrik/rolodex/services:watch",
    "//projects/Rolodex/src/main/java/uk/co/amrik/rolodex/services:health",
    "@maven//:com_google_inject_extensions_guice_servlet",
    "@maven//:com_google_inject_guice",
    "@maven//:jakarta_servlet_jakarta_servlet_api",
    "@maven//:org_eclipse_jetty_ee10_jetty_ee10_servlet",
    "@maven//:org_eclipse_jetty_jetty_server",
    "@maven//:org_slf4j_slf4j_api",
]

RUNTIME_DEPS = [
    "@maven//:org_eclipse_jetty_jetty_slf4j_impl",
]

TEST_DEPS = []

java_binary(
    name = "Rolodex",
    srcs = [
        ":srcs",
    ],
    main_class = "uk.co.amrik.rolodex.Rolodex",
    deps = [
        ":rolodex",
    ] + RUNTIME_DEPS + COMMON_DEPS,
)

java_library(
    name = "rolodex",
    srcs = glob(
        ["*.java"],
    ),
    resources = ["//projects/Rolodex/src/main/resources"],
    deps = [
    ] + COMMON_DEPS,
)

filegroup(
    name = "srcs",
    srcs = ["BUILD"] + glob(["**/*.java"]),
)
