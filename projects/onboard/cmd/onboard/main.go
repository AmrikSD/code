// Command onboard clones your GitHub repositories onto a new machine. It shows
// your account and organisations with a tick box each, and clones whatever is
// ticked into <dir>/<owner>/<repo>.
package main

import (
	"flag"
	"fmt"
	"os"
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
	yes := flag.Bool("yes", false, "skip the picker and clone everything: the -owner list, or your own account")
	flag.Parse()

	if _, err := github.GH("auth", "status"); err != nil {
		fmt.Fprintln(os.Stderr, "Not signed in to GitHub. Run: gh auth login")
		return 1
	}

	var owners []string
	if *owner != "" {
		owners = strings.Split(*owner, ",")
	}

	if *yes {
		return cloneAll(*dir, owners)
	}

	model, err := tui.Run(tui.Config{Root: *dir, Owners: owners, GH: github.GH})
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
