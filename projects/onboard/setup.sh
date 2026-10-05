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

# Bazel's JVM ignores SSL_CERT_FILE. Where that is set because the network
# intercepts TLS, every Bazel download fails until the JVM is given the same
# certificates as a truststore.
if [ -n "${SSL_CERT_FILE:-}" ] && ! grep -qs 'javax.net.ssl.trustStore=' "$HOME/.bazelrc"; then
  store="$HOME/.cache/bazel-truststore.p12"
  if [ ! -e "$store" ]; then
    echo "Building a certificate store for Bazel from $SSL_CERT_FILE"
    mkdir -p "$(dirname "$store")"
    certs=$(mktemp -d)
    awk -v dir="$certs" '/BEGIN CERTIFICATE/ { if (out) close(out); out = dir "/" ++n ".pem" } out { print > out }' "$SSL_CERT_FILE"
    nix shell nixpkgs#jdk --command bash -c '
      for cert in "$1"/*.pem; do
        keytool -importcert -noprompt -alias "$(basename "$cert" .pem)" -file "$cert" \
          -keystore "$2" -storetype PKCS12 -storepass changeit >/dev/null 2>&1
      done' truststore "$certs" "$store"
  fi
  {
    echo "startup --host_jvm_args=-Djavax.net.ssl.trustStore=$store"
    echo "startup --host_jvm_args=-Djavax.net.ssl.trustStorePassword=changeit"
  } >> "$HOME/.bazelrc"
fi

exec nix shell nixpkgs#bazelisk nixpkgs#gh nixpkgs#git --command \
  bazelisk run //projects/onboard/cmd/onboard -- "$@"
