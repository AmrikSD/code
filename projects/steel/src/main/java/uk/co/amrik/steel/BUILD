load("@rules_java//java:defs.bzl", "java_binary", "java_library")

package(default_visibility = ["//visibility:public"])

java_binary(
    name = "ProjectSteel",
    main_class = "uk.co.amrik.steel.Steel",
    srcs = [":srcs"],
    deps = [
        ":steel",
        "//projects/steel/src/main/java/uk/co/amrik/steel/order",
        "//projects/steel/src/main/java/uk/co/amrik/steel/user",
        "@maven//:com_google_inject_guice",
        "@maven//:com_google_inject_extensions_guice_servlet",
        "@maven//:org_jboss_resteasy_resteasy_client",
        "@maven//:org_jboss_resteasy_resteasy_servlet_initializer",
        "@maven//:org_jboss_resteasy_resteasy_jaxb_provider",
        "@maven//:jakarta_platform_jakarta_jakartaee_web_api",
        "@maven//:org_eclipse_jetty_jetty_server",
        "@maven//:org_eclipse_jetty_jetty_servlet",
        "@maven//:org_eclipse_jetty_jetty_util",
        "@maven//:org_glassfish_jersey_core_jersey_server",
        "@maven//:org_glassfish_jersey_containers_jersey_container_servlet_core",
        "@maven//:org_glassfish_jersey_inject_jersey_hk2"
    ]
)

java_library(
    name = "steel",
    srcs = glob(
        ["*.java"]
    ),
    deps = [
        "//projects/steel/src/main/java/uk/co/amrik/steel/order",
        "//projects/steel/src/main/java/uk/co/amrik/steel/user",
        "@maven//:com_google_inject_guice",
        "@maven//:com_google_inject_extensions_guice_servlet",
        "@maven//:org_jboss_resteasy_resteasy_client",
        "@maven//:org_jboss_resteasy_resteasy_servlet_initializer",
        "@maven//:org_jboss_resteasy_resteasy_jaxb_provider",
        "@maven//:jakarta_platform_jakarta_jakartaee_web_api",
        "@maven//:org_eclipse_jetty_jetty_server",
        "@maven//:org_eclipse_jetty_jetty_servlet",
        "@maven//:org_eclipse_jetty_jetty_util",
        "@maven//:org_glassfish_jersey_core_jersey_server",
        "@maven//:org_glassfish_jersey_containers_jersey_container_servlet_core",
        "@maven//:org_glassfish_jersey_inject_jersey_hk2"
    ]
)

filegroup(
    name = "srcs",
    srcs = ["BUILD"] + glob(["**/*.java"]),
)
