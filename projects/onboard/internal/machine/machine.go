// Package machine sets the machine itself up: it fetches the dotfiles repo and
// runs its bootstrap, which installs tools and links config.
package machine

import (
	"os"
	"os/exec"
	"path/filepath"
)

const script = `set -euo pipefail
repo=$1 dir=$2 work=$3
if [ -d "$dir/.git" ]; then
  git -C "$dir" pull --ff-only || echo "Could not update $dir, carrying on with what is there."
else
  gh repo clone "$repo" "$dir"
fi
if [ "$work" = work ]; then git -C "$dir" submodule update --init; fi
"$dir/bootstrap.sh"
printf '\nPress enter to continue. '
read -r _
`

// HasDotfiles reports whether the dotfiles repo is already checked out at dir.
func HasDotfiles(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, ".git"))
	return err == nil
}

// Bootstrap returns the command that clones repo into dir, or updates the
// checkout already there, and runs its bootstrap.sh. With work set it fetches the repo's submodules first, which
// is where the work-only config lives.
func Bootstrap(repo, dir string, work bool) *exec.Cmd {
	mode := "personal"
	if work {
		mode = "work"
	}
	return exec.Command("bash", "-c", script, "bootstrap", repo, dir, mode)
}
