// Package workspace manages the local checkouts, laid out as <root>/<owner>/<repo>.
package workspace

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Path is where a repo is checked out.
func Path(root, owner, repo string) string {
	return filepath.Join(root, owner, repo)
}

// Cloned reports whether a checkout of the repo already exists.
func Cloned(root, owner, repo string) bool {
	_, err := os.Stat(filepath.Join(Path(root, owner, repo), ".git"))
	return err == nil
}

// Clone checks the repo out into its place under root.
func Clone(root, owner, repo, url string) error {
	dest := Path(root, owner, repo)
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	cmd := exec.Command("git", "clone", "--quiet", url, dest)
	cmd.Env = append(os.Environ(), "GIT_TERMINAL_PROMPT=0")
	if os.Getenv("GIT_SSH_COMMAND") == "" {
		// A new machine has no known_hosts yet, and a prompt would hang behind the TUI.
		cmd.Env = append(cmd.Env, "GIT_SSH_COMMAND=ssh -o BatchMode=yes -o StrictHostKeyChecking=accept-new")
	}
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s", lastLine(string(out), err.Error()))
	}
	return nil
}

// Register adds each owner's checkouts to gita, grouped by owner. It does
// nothing when gita is not installed.
func Register(root string, owners []string) error {
	gita, err := exec.LookPath("gita")
	if err != nil {
		// Tools installed during this run are not on this process's PATH yet.
		home, _ := os.UserHomeDir()
		gita = filepath.Join(home, ".nix-profile", "bin", "gita")
		if _, err := os.Stat(gita); err != nil {
			return nil
		}
	}
	for _, owner := range owners {
		if out, err := exec.Command(gita, "add", "-a", filepath.Join(root, owner)).CombinedOutput(); err != nil {
			return fmt.Errorf("gita add %s: %s", owner, lastLine(string(out), err.Error()))
		}
	}
	return nil
}

func lastLine(out, fallback string) string {
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if last := strings.TrimSpace(lines[len(lines)-1]); last != "" {
		return last
	}
	return fallback
}
