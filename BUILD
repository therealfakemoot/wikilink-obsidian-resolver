load("@gazelle//:def.bzl", "gazelle")
load("@rules_go//go:def.bzl", "go_library", "go_test")

# gazelle:prefix github.com/therealfakemoot/wikilink-obsidian-resolver
gazelle(name = "gazelle")

go_library(
    name = "wikilink-obsidian-resolver",
    srcs = [
        "resolver.go",
        "version.go",
    ],
    importpath = "github.com/therealfakemoot/wikilink-obsidian-resolver",
    visibility = ["//visibility:public"],
    x_defs = {
        "Version": "{STABLE_STAMP_VERSION}",
        "Build": "{STABLE_STAMP_COMMIT}",
    },
    deps = [
        "@com_github_therealfakemoot_go_obsidian//:go-obsidian",
        "@dev_abhg_go_goldmark_wikilink//:wikilink",
        "@org_uber_go_zap//:zap",
    ],
)

go_test(
    name = "wikilink-obsidian-resolver_test",
    srcs = ["resolver_test.go"],
    data = glob(["testdata/**"]),
    deps = [
        ":wikilink-obsidian-resolver",
        "@com_github_stretchr_testify//assert",
        "@dev_abhg_go_goldmark_wikilink//:wikilink",
    ],
)
