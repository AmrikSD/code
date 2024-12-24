load("@rules_java//java:defs.bzl", "java_library")

package(default_visibility = ["//visibility:public"])

filegroup(
    name = "user_root",
    srcs = glob(["*.java"]),
)

filegroup(
    name = "user_api_v1",
    srcs = glob(["api/v1/**/*.java"]),
)

java_library(
    name = "user",
    srcs = [
        ":user_api_v1",
        ":user_root",
    ],
    javacopts = [
        "-source",
        "21",
        "-target",
        "21",
    ],
    deps = [
        "//projects/steel/src/main/java/uk/co/amrik/steel/API",
        "//projects/steel/src/main/java/uk/co/amrik/steel/middleware",
        "//projects/steel/src/main/java/uk/co/amrik/steel/permissions",
        "//projects/steel/src/main/java/uk/co/amrik/steel/persistence",
        "@maven//:com_fasterxml_jackson_datatype_jackson_datatype_jdk8",
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
