load("@aspect_rules_ts//ts:defs.bzl", "ts_project")

SRCS = [
    "_app.tsx",
    "_document.tsx",
    "index.tsx",
    "location.tsx",
    "rsvp.tsx",
    "start.module.css",
    "start.tsx",
]

ts_project(
    name = "pages",
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
    ]
)
