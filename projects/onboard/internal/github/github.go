// Package github lists the owners and repositories the signed-in user can
// clone, by shelling out to the gh CLI.
package github

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

// Repo is a repository as gh reports it.
type Repo struct {
	Name   string `json:"name"`
	SSHURL string `json:"sshUrl"`
}

// Runner runs gh with the given arguments and returns its standard output.
type Runner func(args ...string) ([]byte, error)

// GH runs the real gh CLI.
func GH(args ...string) ([]byte, error) {
	out, err := exec.Command("gh", args...).Output()
	if err != nil {
		if exit, ok := err.(*exec.ExitError); ok && len(exit.Stderr) > 0 {
			return nil, fmt.Errorf("gh %s: %s", args[0], strings.TrimSpace(string(exit.Stderr)))
		}
		return nil, fmt.Errorf("gh %s: %w", args[0], err)
	}
	return out, nil
}

// Owners returns the signed-in user followed by the organisations they belong to.
func Owners(run Runner) ([]string, error) {
	user, err := run("api", "user", "--jq", ".login")
	if err != nil {
		return nil, err
	}
	orgs, err := run("api", "user/orgs", "--paginate", "--jq", ".[].login")
	if err != nil {
		return nil, err
	}
	return append(strings.Fields(string(user)), strings.Fields(string(orgs))...), nil
}

// Repos lists an owner's repositories by name, leaving out archived ones.
func Repos(run Runner, owner string) ([]Repo, error) {
	out, err := run("repo", "list", owner, "--no-archived", "--limit", "5000", "--json", "name,sshUrl")
	if err != nil {
		return nil, err
	}
	return parseRepos(out)
}

func parseRepos(out []byte) ([]Repo, error) {
	var repos []Repo
	if err := json.Unmarshal(out, &repos); err != nil {
		return nil, fmt.Errorf("parsing gh repo list output: %w", err)
	}
	sort.Slice(repos, func(i, j int) bool {
		return strings.ToLower(repos[i].Name) < strings.ToLower(repos[j].Name)
	})
	return repos, nil
}
