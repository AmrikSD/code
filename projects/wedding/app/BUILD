load("@aspect_rules_jest//jest:defs.bzl", "jest_test")
load("@aspect_rules_ts//ts:defs.bzl", "ts_project")

SRCS = [
    "styles/global.css",
    "Accordion.tsx",
    "Countdown.tsx",
    "favicon.ico",
    "section.tsx",

    "accordion.module.css",
    "countdown.module.css",
    "section.module.css",
]

ts_project(
    name = "app",
    srcs = SRCS,
    allow_js = True,
    no_emit = True,
    preserve_jsx = True,
    resolve_json_module = True,
    transpiler = "tsc",
    tsconfig = "//projects/wedding:tsconfig",
    visibility = ["//projects/wedding:__subpackages__"],
    deps = [
        "//projects/wedding:node_modules/@types/react",
        "//projects/wedding:node_modules/@types/react-dom",
        "//projects/wedding:node_modules/next",
    ],
)
