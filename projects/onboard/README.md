# onboard

Sets up a new machine from nothing.

## Setup

```sh
git clone https://github.com/AmrikSD/code ~/code/AmrikSD/code
~/code/AmrikSD/code/projects/onboard/setup.sh
```

`setup.sh` installs Nix if it is missing, then builds and starts the onboarding flow:

1. GitHub: signs in with `gh` in the browser if needed. No SSH key is created.
2. Tools and config: clones the dotfiles repo into `~/.dotfiles`, or updates it, and runs its `bootstrap.sh`. That installs every tool and app, links the config (including SSH through the 1Password agent), adds the Claude Code hooks and imports the GPG signing key. Tick "Work setup" on a work machine to fetch the work-only config too.
3. Repositories: lists your account and organisations with a tick box each. Ticking an owner ticks every repo under it, opening one lets you pick repos individually. Whatever is ticked is cloned into `~/code/<owner>/<repo>`.

Then open a new terminal.

On a network that intercepts TLS and sets `SSL_CERT_FILE`, as some company laptops do, `setup.sh` first builds a certificate store from that file and points Bazel at it in `~/.bazelrc`. Bazel cannot download anything there otherwise.

Before you start: install the 1Password app, sign in, and in its Developer settings turn on "Use the SSH agent" and "Integrate with 1Password CLI". Your SSH key, GPG key and the work secrets all come from it, so every machine uses the same keys. Without the SSH agent the repositories step cannot clone anything.

## Enjoy

Every step is safe to re-run, and repos that are already cloned are skipped. To run it again on a machine that is set up:

```sh
bazel run //projects/onboard/cmd/onboard
```

Day to day, tools are added, pinned and upgraded through the dotfiles repo. See `~/.dotfiles/NIX.md`.

## WIP

Known gaps, most useful first:

- The Mac-only parts of the dotfiles bootstrap have not had a real first run: copying apps into `~/Applications/Nix Apps`, pointing gpg at the Nix pinentry, the `docker compose` link and the 1Password steps.
- `setup.sh` has only been read, not run, since running it installs Nix.
- Apple Silicon and Linux only. Current nixpkgs has dropped Intel Macs.
- On ARM Linux, Node 26 has no prebuilt package and fails to compile, which fails the tool install. Swap `nodejs_26` for `nodejs_24` in `~/.dotfiles/tools/tools.nix`.
- No search in the picker. Individual repos in a big organisation mean scrolling.
