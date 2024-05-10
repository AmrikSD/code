load("@rules_java//java:defs.bzl", "java_test")

package(default_visibility = ["//visibility:public"])

java_test(
    name = "LeetcodeTest",
    srcs = ["LeetcodeTest.java"],
    deps = [
        "@maven//:junit_junit"
    ]
)

java_test(
    name = "TwoSumTest",
    srcs = ["TwoSumTest.java"],

    deps = [
        "//projects/Leetcode/src/main/java/uk/co/amrik/leetcode:twosum-lib",
        "@maven//:junit_junit",
        "@maven//:org_assertj_assertj_core"
    ]
)

java_test(
    name = "MergeAlternatelyTest",
    srcs = ["MergeAlternatelyTest.java"],

    deps = [
        "//projects/Leetcode/src/main/java/uk/co/amrik/leetcode:mergealternately-lib",
        "@maven//:junit_junit",
        "@maven//:org_assertj_assertj_core"
    ]
)

filegroup(
    name = "srcs",
    srcs = glob(["**/*.java"]) + ["BUILD"],
    visibility = ["//projects/Leetcode:__pkg__"],
)
