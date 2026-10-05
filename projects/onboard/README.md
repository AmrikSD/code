# onboard

Sets up a new machine.

```sh
git clone https://github.com/AmrikSD/code ~/code/AmrikSD/code
~/code/AmrikSD/code/projects/onboard/setup.sh
```

`setup.sh` installs Nix if it is missing, then builds and starts the onboarding flow:

1. Signs in to GitHub with `gh` if needed. Pick SSH and let it upload a key.
2. Tools and config: clones the dotfiles repo into `~/.dotfiles` and runs its `bootstrap.sh`, which installs every tool and app and links the config. Tick "Work setup" to fetch the work-only config too.
3. Repositories: lists your account and organisations with a tick box each. Ticking an owner ticks every repo under it, opening one lets you pick repos individually. Whatever is ticked is cloned into `~/code/<owner>/<repo>`.

Every step is safe to re-run. Repos that are already cloned are skipped.

On a machine that is already set up, run it directly:

```sh
bazel run //projects/onboard/cmd/onboard
```
