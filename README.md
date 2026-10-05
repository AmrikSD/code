# code

My repo, for everything: homelab, personal projects, and whatever else I end up
poking at. All of it backed by [Bazel](https://bazel.build), or at least heading
that way.

## Why a monorepo, why Bazel

Mostly because I wanted to learn Bazel properly, and the best way to do that is
to make it build a pile of unrelated things in different languages. One repo
also means one CI pipeline, one Renovate config and one place to look when I
forget where I put something.

This is not how I'd set up a repo at work. Bazel pays for itself with many
engineers, a big dependency graph and slow builds that benefit from caching and
hermeticity. For a handful of small projects owned by one person, it's a steep
learning curve and a lot of BUILD file upkeep for wins that `go build` and
`mvn` would give me for free. I'm happy paying that cost here because the
learning is the point.

## [Projects](./projects)

| Project | What it is | Stack |
|---------|------------|-------|
| [amrik.co.uk](./projects/amrik.co.uk) | Personal site and blog | Jekyll |
| [Rolodex](./projects/Rolodex) | Document indexer, inspired by Paperless-ngx | Java, Postgres |
| [vibe](./projects/vibe) | CLI that starts work on a Jira ticket | Go |
| [onboard](./projects/onboard) | TUI that clones my repos onto a new machine | Go |
| [csv-to-pdf](./projects/csv-to-pdf) | Turns scuffed CSVs into decent looking PDFs | HTML, JS |
| [Leetcode](./projects/Leetcode) | Practice problems | Go, Java |

## [Infra](./infra)

Infrastructure as code for the homelab and everything the projects run on.

| Directory | What it is |
|-----------|------------|
| [aws](./infra/aws) | AWS accounts and S3, in Terraform |
| [cloudflare](./infra/cloudflare) | DNS, email and Pages for amrik.co.uk, in Terraform |
| [unifi](./infra/unifi) | Home network config, in Terraform |
| [containers](./infra/containers) | Upstream images mirrored into my own registry, built with Bazel |

Secrets are committed encrypted with [sops](https://github.com/getsops/sops).

## Building

```sh
bazel build //...
bazel test //...
bazel run //:gazelle
```
