load("@contrib_rules_jvm//java:defs.bzl", "java_junit5_test")

package(default_visibility = ["//visibility:public"])

java_junit5_test(
    name = "OrderServiceTest",
    srcs = glob(["**/*.java"]),
    test_class = "uk.co.amrik.steel.order.api.v1.OrderServiceTest",
    deps = [
        "//projects/steel/src/main/java/uk/co/amrik/steel/order",
        "//projects/steel/src/main/java/uk/co/amrik/steel/persistence",
        "@maven//:com_google_inject_guice",
        "@maven//:org_assertj_assertj_core",
        "@maven//:org_junit_platform_junit_platform_launcher",
        "@maven//:org_junit_platform_junit_platform_reporting",
        "@maven//:org_junit_jupiter_junit_jupiter_api",
        "@maven//:org_junit_jupiter_junit_jupiter_engine",
        "@maven//:org_flywaydb_flyway_core",
        "@maven//:org_testcontainers_postgresql",
        "@maven//:org_testcontainers_testcontainers",
        "@maven//:org_mockito_mockito_core",
        "@maven//:org_mockito_mockito_junit_jupiter",
    ],
)
