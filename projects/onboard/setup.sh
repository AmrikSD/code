#!/usr/bin/env bash
# The first thing to run on a new machine. Installs Nix if it is missing, then
# starts the onboarding flow with the few tools it needs.
set -eo pipefail

cd "$(dirname "$0")/../.."

nix_profile=/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh
if ! command -v nix >/dev/null; then
  if [ ! -e "$nix_profile" ]; then
    curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install
  fi
  # shellcheck disable=SC1090
  . "$nix_profile"
fi

exec nix shell nixpkgs#bazelisk nixpkgs#gh nixpkgs#git --command \
  bazelisk run //projects/onboard/cmd/onboard -- "$@"
