load("@rules_java//java:defs.bzl", "java_library")

package(default_visibility = ["//visibility:public"])

filegroup(
    name = "order_root",
    srcs = glob(["*.java"]),
)

filegroup(
    name = "order_api_v1",
    srcs = glob(["api/v1/**/*.java"]),
)

filegroup(
    name = "order_model",
    srcs = glob(["model/**/*.java"]),
)


java_library(
    name = "order",
    srcs = [
        ":order_root",
        ":order_model",
        ":order_api_v1"
    ],
    deps = [
        "//projects/steel/src/main/java/uk/co/amrik/steel/persistence:persistence",
        "@maven//:com_google_inject_guice",
        "@maven//:com_google_inject_extensions_guice_servlet",
        "@maven//:jakarta_platform_jakarta_jakartaee_web_api",
        "@maven//:org_glassfish_jersey_core_jersey_server",
        "@maven//:org_glassfish_jersey_containers_jersey_container_servlet_core",
        "@maven//:org_glassfish_jersey_inject_jersey_hk2",
        "@maven//:org_glassfish_jersey_media_jersey_media_json_jackson",
        "@maven//:org_jooq_jooq"
    ]
)
