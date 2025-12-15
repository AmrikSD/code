load("@rules_java//java:defs.bzl", "java_test")

package(default_visibility = ["//visibility:public"])

java_test(
    name = "LeetcodeTest",
    srcs = ["LeetcodeTest.java"],
    deps = [
        "@maven//:org_junit_jupiter_junit_jupiter_api",
        "@maven//:org_junit_platform_junit_platform_runner",
    ],
    runtime_deps = [
        "@maven//:org_junit_jupiter_junit_jupiter_engine",
        "@maven//:org_junit_platform_junit_platform_console",
        "@maven//:org_junit_platform_junit_platform_launcher",
        "@maven//:org_junit_platform_junit_platform_reporting",
    ],
)

java_test(
    name = "TwoSumTest",
    srcs = ["TwoSumTest.java"],
    deps = [
        "//projects/Leetcode/Java/src/main/java/uk/co/amrik/leetcode:lc-library",
        "@maven//:org_junit_jupiter_junit_jupiter_api",
        "@maven//:org_assertj_assertj_core",
        "@maven//:org_junit_platform_junit_platform_runner",
    ],
    runtime_deps = [
        "@maven//:org_junit_jupiter_junit_jupiter_engine",
        "@maven//:org_junit_platform_junit_platform_console",
        "@maven//:org_junit_platform_junit_platform_launcher",
        "@maven//:org_junit_platform_junit_platform_reporting",
    ],
)

java_test(
    name = "MergeAlternatelyTest",
    srcs = ["MergeAlternatelyTest.java"],
    deps = [
        "//projects/Leetcode/Java/src/main/java/uk/co/amrik/leetcode:lc-library",
        "@maven//:org_junit_jupiter_junit_jupiter_api",
        "@maven//:org_assertj_assertj_core",
        "@maven//:org_junit_platform_junit_platform_runner",
    ],
    runtime_deps = [
        "@maven//:org_junit_jupiter_junit_jupiter_engine",
        "@maven//:org_junit_platform_junit_platform_console",
        "@maven//:org_junit_platform_junit_platform_launcher",
        "@maven//:org_junit_platform_junit_platform_reporting",
    ],
)

java_test(
    name = "MoveZerosTest",
    srcs = ["MoveZerosTest.java"],
    deps = [
        "//projects/Leetcode/Java/src/main/java/uk/co/amrik/leetcode:lc-library",
        "@maven//:org_junit_jupiter_junit_jupiter_api",
        "@maven//:org_assertj_assertj_core",
        "@maven//:org_junit_platform_junit_platform_runner",
    ],
    runtime_deps = [
        "@maven//:org_junit_jupiter_junit_jupiter_engine",
        "@maven//:org_junit_platform_junit_platform_console",
        "@maven//:org_junit_platform_junit_platform_launcher",
        "@maven//:org_junit_platform_junit_platform_reporting",
    ],
)

java_test(
    name = "IsSubsequenceTest",
    srcs = ["IsSubsequenceTest.java"],
    deps = [
        "//projects/Leetcode/Java/src/main/java/uk/co/amrik/leetcode:lc-library",
        "@maven//:org_junit_jupiter_junit_jupiter_api",
        "@maven//:org_assertj_assertj_core",
        "@maven//:org_junit_platform_junit_platform_runner",
    ],
    runtime_deps = [
        "@maven//:org_junit_jupiter_junit_jupiter_engine",
        "@maven//:org_junit_platform_junit_platform_console",
        "@maven//:org_junit_platform_junit_platform_launcher",
        "@maven//:org_junit_platform_junit_platform_reporting",
    ],
)

filegroup(
    name = "srcs",
    srcs = glob(["**/*.java"]) + ["BUILD"],
    visibility = ["//projects/Leetcode/Java:__pkg__"],
)
