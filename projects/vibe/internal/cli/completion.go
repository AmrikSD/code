package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/amriksd/code/projects/vibe/internal/jira"
)

const completionUsage = `vibe completion - shell completion

Prints a completion script for your shell. Once loaded, tab-completing the
<jira-key> argument offers the open tickets assigned to you (with status and
summary), and subcommands and flags complete too.

Usage:
  vibe completion zsh
  vibe completion bash

Install:
  # zsh (~/.zshrc)
  source <(vibe completion zsh)

  # bash (~/.bashrc)
  source <(vibe completion bash)

Ticket results are cached for a couple of minutes so repeated tabs are
instant. Claiming a ticket with "vibe queue" clears the cache.
`

// assignedJQL is what the <jira-key> completion lists.
const assignedJQL = "assignee = currentUser() AND statusCategory != Done"

const (
	completionCacheTTL   = 2 * time.Minute
	completionCacheLimit = 50
)

// candidate is one completion suggestion.
type candidate struct {
	value string
	desc  string
}

var (
	subcommands = []candidate{
		{"queue", "work the service desk queue"},
		{"completion", "print a shell completion script"},
	}
	shells = []candidate{
		{"zsh", ""},
		{"bash", ""},
	}
	rootFlags = []candidate{
		{"--dry-run", "show what would happen without executing"},
		{"--no-tmux", "skip tmux window creation"},
		{"--version", "print version and exit"},
		{"--help", "show help"},
	}
	queueFlags = []candidate{
		{"--list", "print the queue and exit"},
		{"--next", "take the oldest unassigned To Do ticket"},
		{"--no-claim", "don't assign or move the ticket"},
		{"--project", "service desk project (default SUP)"},
		{"--jql", "override the queue query"},
		{"--limit", "maximum tickets to list"},
		{"--dry-run", "show what would happen without executing"},
		{"--no-tmux", "skip tmux window creation"},
		{"--help", "show help"},
	}
)

func runCompletion(argv []string) int {
	if len(argv) == 0 || argv[0] == "--help" || argv[0] == "-h" {
		fmt.Fprint(os.Stderr, completionUsage)
		if len(argv) == 0 {
			return 1
		}
		return 0
	}
	switch argv[0] {
	case "zsh":
		fmt.Print(zshCompletion)
	case "bash":
		fmt.Print(bashCompletion)
	default:
		fmt.Fprintf(os.Stderr, "error: unsupported shell %q (want zsh or bash)\n", argv[0])
		return 1
	}
	return 0
}

// runComplete implements the hidden "vibe __complete" command that the shell
// scripts call. It receives every word after "vibe" up to and including the
// word being completed (which may be empty) and prints one candidate per
// line as "value<TAB>description".
func runComplete(words []string) int {
	for _, c := range completeCandidates(words, assignedTickets) {
		fmt.Printf("%s\t%s\n", c.value, c.desc)
	}
	return 0
}

// completeCandidates works out what position on the command line is being
// completed and returns the matching candidates. tickets is only called when
// a Jira key is wanted, so the shell never waits on Jira for flag completion.
func completeCandidates(words []string, tickets func() []candidate) []candidate {
	if len(words) == 0 {
		words = []string{""}
	}
	cur := words[len(words)-1]
	prev := words[:len(words)-1]
	wantFlag := strings.HasPrefix(cur, "-")

	// The first non-flag word decides the context.
	positional := ""
	for _, w := range prev {
		if !strings.HasPrefix(w, "-") {
			positional = w
			break
		}
	}

	var out []candidate
	switch positional {
	case "":
		if wantFlag {
			out = rootFlags
		} else {
			out = append(out, subcommands...)
			out = append(out, tickets()...)
		}
	case "queue":
		if wantFlag {
			out = queueFlags
		}
	case "completion":
		if !wantFlag && len(prev) == 1 {
			out = shells
		}
	default:
		// A Jira key is already there; only flags are left.
		if wantFlag {
			out = rootFlags
		}
	}

	var matched []candidate
	for _, c := range out {
		if strings.HasPrefix(strings.ToUpper(c.value), strings.ToUpper(cur)) {
			matched = append(matched, c)
		}
	}
	return matched
}

// assignedTickets returns the caller's open tickets as candidates, reading a
// short-lived cache first so consecutive tabs don't each hit Jira. Errors are
// swallowed: a completion that fails just offers nothing.
func assignedTickets() []candidate {
	if cached, ok := readCompletionCache(); ok {
		return cached
	}
	issues, err := jira.Search(assignedJQL, completionCacheLimit)
	if err != nil {
		return nil
	}
	sortForCompletion(issues)
	out := make([]candidate, 0, len(issues))
	for _, is := range issues {
		out = append(out, candidate{is.Key, is.Status + ": " + oneLine(is.Summary, 60)})
	}
	writeCompletionCache(out)
	return out
}

// sortForCompletion puts In Progress tickets first, then the most recently
// updated. The shell scripts are told to keep this order.
func sortForCompletion(issues []jira.Issue) {
	sort.SliceStable(issues, func(i, j int) bool {
		a, b := issues[i], issues[j]
		if (a.Status == statusInProgress) != (b.Status == statusInProgress) {
			return a.Status == statusInProgress
		}
		return a.Updated.After(b.Updated)
	})
}

func completionCachePath() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "vibe", "completion-tickets"), nil
}

func readCompletionCache() ([]candidate, bool) {
	path, err := completionCachePath()
	if err != nil {
		return nil, false
	}
	info, err := os.Stat(path)
	if err != nil || time.Since(info.ModTime()) > completionCacheTTL {
		return nil, false
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, false
	}
	defer f.Close()

	var out []candidate
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		value, desc, _ := strings.Cut(sc.Text(), "\t")
		if value != "" {
			out = append(out, candidate{value, desc})
		}
	}
	return out, sc.Err() == nil
}

func writeCompletionCache(cs []candidate) {
	path, err := completionCachePath()
	if err != nil {
		return
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return
	}
	var b strings.Builder
	for _, c := range cs {
		fmt.Fprintf(&b, "%s\t%s\n", c.value, c.desc)
	}
	// Write then rename so a concurrent completion never reads a torn file.
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, []byte(b.String()), 0o644); err != nil {
		return
	}
	_ = os.Rename(tmp, path)
}

// invalidateCompletionCache drops the cached ticket list, e.g. after claiming
// a ticket so it shows up on the next tab.
func invalidateCompletionCache() {
	path, err := completionCachePath()
	if err != nil {
		return
	}
	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return
	}
}

const zshCompletion = `#compdef vibe
# zsh completion for vibe. Load with:  source <(vibe completion zsh)

if (( ! ${+functions[compdef]} )); then
    autoload -Uz compinit && compinit
fi

_vibe() {
    local -a lines completions
    local line value desc
    # Every word after "vibe" up to and including the one being completed.
    lines=("${(@f)$(vibe __complete "${(@)words[2,CURRENT]}" 2>/dev/null)}")
    for line in "${lines[@]}"; do
        [[ -z "$line" ]] && continue
        value="${line%%$'\t'*}"
        desc="${line#*$'\t'}"
        [[ "$desc" == "$line" ]] && desc=""
        completions+=("${value//:/\\:}:${desc}")
    done
    # -o nosort keeps vibe's order: In Progress first, then most recently updated.
    _describe -t vibe 'vibe' completions -o nosort
}

compdef _vibe vibe
`

const bashCompletion = `# bash completion for vibe. Load with:  source <(vibe completion bash)

_vibe() {
    local cur="${COMP_WORDS[COMP_CWORD]}"
    local -a args=("${COMP_WORDS[@]:1:COMP_CWORD}")
    local values
    values="$(vibe __complete "${args[@]}" 2>/dev/null | cut -f1)"
    COMPREPLY=($(compgen -W "${values}" -- "${cur}"))
}

# bash >= 4.4 can keep vibe's order (In Progress first, then most recently updated).
complete -o nosort -F _vibe vibe 2>/dev/null || complete -F _vibe vibe
`
