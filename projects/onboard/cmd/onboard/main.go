// Command onboard sets up a new machine. It fetches the dotfiles repo and runs
// its bootstrap, then shows your GitHub account and organisations with a tick
// box each and clones whatever is ticked into <dir>/<owner>/<repo>.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/amriksd/code/projects/onboard/internal/github"
	"github.com/amriksd/code/projects/onboard/internal/tui"
	"github.com/amriksd/code/projects/onboard/internal/workspace"
)

func main() {
	os.Exit(run())
}

func run() int {
	home, _ := os.UserHomeDir()
	dir := flag.String("dir", filepath.Join(home, "code"), "directory to clone into, as <dir>/<owner>/<repo>")
	owner := flag.String("owner", "", "comma separated owners to offer, instead of your account and all your organisations")
	yes := flag.Bool("yes", false, "skip the setup and picker and clone everything: the -owner list, or your own account")
	dotfiles := flag.String("dotfiles", "AmrikSD/.dotfiles", "dotfiles repo whose bootstrap.sh sets the machine up")
	dotfilesDir := flag.String("dotfiles-dir", filepath.Join(home, ".dotfiles"), "where the dotfiles repo lives")
	flag.Parse()

	if !signedIn() {
		return 1
	}

	var owners []string
	if *owner != "" {
		owners = strings.Split(*owner, ",")
	}

	if *yes {
		return cloneAll(*dir, owners)
	}

	model, err := tui.Run(tui.Config{Root: *dir, Owners: owners, GH: github.GH, Dotfiles: *dotfiles, DotfilesDir: *dotfilesDir})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	failed := model.Failed()
	fmt.Printf("Cloned %d of %d into %s\n", len(model.Results)-len(failed), len(model.Results), *dir)
	for _, r := range failed {
		fmt.Printf("  %s/%s: %v\n", r.Repo.Owner, r.Repo.Name, r.Err)
	}
	if len(failed) > 0 {
		return 1
	}
	return 0
}

// signedIn makes sure gh is signed in, walking through gh's own login when it
// is not. Repos are cloned over SSH with the key the dotfiles set up, so gh is
// told not to generate one for this machine.
func signedIn() bool {
	if _, err := github.GH("auth", "status"); err == nil {
		return true
	}
	login := exec.Command("gh", "auth", "login", "--git-protocol", "ssh", "--skip-ssh-key", "--web")
	login.Stdin, login.Stdout, login.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := login.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Could not sign in to GitHub:", err)
		return false
	}
	return true
}

func cloneAll(dir string, owners []string) int {
	if len(owners) == 0 {
		all, err := github.Owners(github.GH)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		owners = all[:1]
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	failed := 0
	slots := make(chan struct{}, 6)
	for _, owner := range owners {
		repos, err := github.Repos(github.GH, owner)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		for _, repo := range repos {
			if workspace.Cloned(dir, owner, repo.Name) {
				continue
			}
			wg.Add(1)
			slots <- struct{}{}
			go func() {
				defer wg.Done()
				err := workspace.Clone(dir, owner, repo.Name, repo.SSHURL)
				<-slots
				mu.Lock()
				defer mu.Unlock()
				if err != nil {
					failed++
					fmt.Printf("failed  %s/%s: %v\n", owner, repo.Name, err)
					return
				}
				fmt.Printf("cloned  %s/%s\n", owner, repo.Name)
			}()
		}
	}
	wg.Wait()

	if err := workspace.Register(dir, owners); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if failed > 0 {
		return 1
	}
	return 0
}
