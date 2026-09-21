package cli

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/amriksd/code/projects/vibe/internal/jira"
)

const queueUsage = `vibe queue - work the service desk queue

Lists open tickets in the service desk queue (oldest first), lets you pick
one, claims it (assigns to you and moves it to an active state), and starts a
vibe session on it with the support-desk prompt.

Usage:
  vibe queue [flags]

Flags:
  --list          Print the queue and exit
  --next          Skip the picker: take the oldest unassigned "To Do" ticket
  --no-claim      Don't assign or transition the ticket
  --claim-states  Preferred transition states, in order (default "In Progress,Start,In review")
  --project KEY   Service desk project (default "SUP")
  --jql QUERY     Override the queue query entirely
  --limit N       Maximum tickets to list (default 50)
  --dry-run       Show what would happen without executing
  --no-tmux       Skip tmux window creation, run opencode in current shell
  --help          Show this help message

Examples:
  vibe queue                  # list, pick, claim, go
  vibe queue --list           # just look
  vibe queue --next           # grab the oldest unclaimed ticket
  vibe queue --next --dry-run
  vibe queue --jql 'project = SUP AND assignee = currentUser() AND statusCategory != Done'
`

const (
	defaultDeskProject = "SUP"
	defaultClaimStates = "In Progress,Start,In review"
	statusToDo         = "To Do"
	statusInProgress   = "In Progress"
)

type queueParams struct {
	List        bool
	Next        bool
	NoClaim     bool
	ClaimStates string
	Project     string
	JQL         string
	Limit       int
	DryRun      bool
	NoTmux      bool
}

func runQueue(argv []string) int {
	fs := flag.NewFlagSet("vibe queue", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprint(os.Stderr, queueUsage)
	}

	var params queueParams
	fs.BoolVar(&params.List, "list", false, "Print the queue and exit")
	fs.BoolVar(&params.Next, "next", false, "Take the oldest unassigned To Do ticket")
	fs.BoolVar(&params.NoClaim, "no-claim", false, "Don't assign/transition the ticket")
	fs.StringVar(&params.ClaimStates, "claim-states", defaultClaimStates, "Preferred transition states, in order")
	fs.StringVar(&params.Project, "project", defaultDeskProject, "Service desk project key")
	fs.StringVar(&params.JQL, "jql", "", "Override the queue query")
	fs.IntVar(&params.Limit, "limit", 50, "Maximum tickets to list")
	fs.BoolVar(&params.DryRun, "dry-run", false, "Show what would happen without executing")
	fs.BoolVar(&params.NoTmux, "no-tmux", false, "Skip tmux window creation")

	if err := fs.Parse(reorderArgs(argv)); err != nil {
		return 1
	}
	if fs.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "error: unexpected argument %q\n", fs.Arg(0))
		fs.Usage()
		return 1
	}

	params.Project = strings.ToUpper(params.Project)
	if params.JQL == "" {
		params.JQL = fmt.Sprintf("project = %s AND statusCategory != Done", params.Project)
	}
	if params.Limit <= 0 || params.Limit > 100 {
		params.Limit = 100
	}
	claimStates := parseStatePreferences(params.ClaimStates)

	issues, err := jira.Search(params.JQL, params.Limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return 1
	}
	if len(issues) == 0 {
		fmt.Println("queue is empty 🎉")
		return 0
	}

	printQueue(issues)
	if params.List {
		return 0
	}

	var picked *jira.Issue
	if params.Next {
		picked = nextUnclaimed(issues)
		if picked == nil {
			fmt.Println("nothing unclaimed in the queue; use the picker or --jql")
			return 0
		}
		fmt.Printf("\nnext up: %s\n", picked.Key)
	} else {
		picked = pickIssue(issues)
		if picked == nil {
			return 0
		}
	}

	if !params.NoClaim {
		if err := claim(picked, params.DryRun, claimStates); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			return 1
		}
	}

	return launch(picked.Key, params.DryRun, params.NoTmux)
}

func printQueue(issues []jira.Issue) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "#\tKEY\tSTATUS\tASSIGNEE\tREPORTER\tAGE\tSUMMARY")
	for i, is := range issues {
		assignee := is.Assignee
		if assignee == "" {
			assignee = "-"
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%s\n",
			i+1, is.Key, is.Status, assignee, is.Reporter, age(is.Created), oneLine(is.Summary, 72))
	}
	w.Flush()
}

// nextUnclaimed returns the oldest ticket nobody has picked up yet.
func nextUnclaimed(issues []jira.Issue) *jira.Issue {
	for i := range issues {
		if issues[i].Assignee == "" && issues[i].Status == statusToDo {
			return &issues[i]
		}
	}
	return nil
}

// pickIssue prompts on stdin for a row number or issue key. Returns nil if
// the user quits.
func pickIssue(issues []jira.Issue) *jira.Issue {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\npick [1-%d, key, or q]: ", len(issues))
		line, err := reader.ReadString('\n')
		if err != nil && line == "" {
			fmt.Println()
			return nil
		}
		in := strings.TrimSpace(line)
		switch {
		case in == "" || strings.EqualFold(in, "q"):
			return nil
		}
		if n, err := strconv.Atoi(in); err == nil {
			if n >= 1 && n <= len(issues) {
				return &issues[n-1]
			}
			fmt.Printf("out of range\n")
			continue
		}
		key := strings.ToUpper(in)
		for i := range issues {
			if issues[i].Key == key {
				return &issues[i]
			}
		}
		fmt.Printf("%q is not in the queue\n", in)
	}
}

// claim assigns the ticket to the current user and moves it to an active state,
// skipping whichever steps are already done.
func claim(is *jira.Issue, dryRun bool, preferredStates []string) error {
	me, err := jira.Me()
	if err != nil {
		return err
	}

	needsAssign := is.Assignee == ""
	needsMove := is.Status == statusToDo

	if dryRun {
		if needsAssign {
			fmt.Printf("[dry-run] jira:       assign %s to %s\n", is.Key, me)
		}
		if needsMove {
			if len(preferredStates) > 0 {
				fmt.Printf("[dry-run] jira:       move %s to %q\n", is.Key, preferredStates[0])
			}
		}
		if !needsAssign && !needsMove {
			fmt.Printf("[dry-run] jira:       %s already claimed (%s, %s)\n", is.Key, is.Assignee, is.Status)
		}
		return nil
	}

	if is.Assignee != "" && !strings.EqualFold(is.Assignee, me) {
		fmt.Printf("note: %s is assigned to %s; leaving assignment alone\n", is.Key, is.Assignee)
	}
	if needsAssign {
		if err := jira.Assign(is.Key, me); err != nil {
			return err
		}
		fmt.Printf("assigned %s to %s\n", is.Key, me)
		invalidateCompletionCache()
	}
	if needsMove {
		toState, err := moveWithFallback(jira.Move, is.Key, preferredStates)
		if err != nil {
			return err
		}
		fmt.Printf("moved %s to %q\n", is.Key, toState)
	}
	return nil
}

func parseStatePreferences(csv string) []string {
	parts := strings.Split(csv, ",")
	out := make([]string, 0, len(parts))
	seen := map[string]bool{}
	for _, part := range parts {
		state := strings.TrimSpace(part)
		if state == "" {
			continue
		}
		key := strings.ToUpper(state)
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, state)
	}
	if len(out) == 0 {
		return []string{statusInProgress}
	}
	return out
}

func moveWithFallback(move func(string, string) error, key string, preferredStates []string) (string, error) {
	if len(preferredStates) == 0 {
		preferredStates = []string{statusInProgress}
	}
	target := preferredStates[0]
	if err := move(key, target); err == nil {
		return target, nil
	} else {
		states := availableStates(err)
		fallback, ok := pickPreferredState(preferredStates, states)
		if !ok || strings.EqualFold(fallback, target) {
			return "", err
		}
		if err2 := move(key, fallback); err2 != nil {
			return "", errors.Join(err, err2)
		}
		return fallback, nil
	}
}

var availableStateRE = regexp.MustCompile(`'([^']+)'`)

func availableStates(err error) []string {
	if err == nil {
		return nil
	}
	matches := availableStateRE.FindAllStringSubmatch(err.Error(), -1)
	out := make([]string, 0, len(matches))
	for _, m := range matches {
		if len(m) < 2 {
			continue
		}
		state := strings.TrimSpace(m[1])
		if state != "" {
			out = append(out, state)
		}
	}
	return out
}

func pickPreferredState(preferredStates, available []string) (string, bool) {
	if len(preferredStates) == 0 || len(available) == 0 {
		return "", false
	}
	for _, want := range preferredStates {
		for _, have := range available {
			if strings.EqualFold(have, want) {
				return have, true
			}
		}
	}
	return "", false
}

func age(t time.Time) string {
	if t.IsZero() {
		return "?"
	}
	d := time.Since(t)
	switch {
	case d < time.Hour:
		return fmt.Sprintf("%dm", int(d.Minutes()))
	case d < 48*time.Hour:
		return fmt.Sprintf("%dh", int(d.Hours()))
	default:
		return fmt.Sprintf("%dd", int(d.Hours()/24))
	}
}

// oneLine collapses whitespace and truncates to max runes.
func oneLine(s string, max int) string {
	s = strings.Join(strings.Fields(s), " ")
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max-1]) + "…"
}
