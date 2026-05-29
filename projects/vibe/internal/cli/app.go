package cli

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/amriksd/code/projects/vibe/internal/prompt"
	"github.com/amriksd/code/projects/vibe/internal/tmux"
	"github.com/amriksd/code/projects/vibe/internal/version"
)

const usage = `vibe - start work on a Jira ticket

Usage:
  vibe <jira-key>
  vibe [flags]

Arguments:
  <jira-key>    Jira issue key (e.g. ENG-3355)

Flags:
  --dry-run     Show what would happen without executing
  --no-tmux     Skip tmux window creation, run opencode in current shell
  --version     Print version and exit
  --help        Show this help message

Examples:
  vibe ENG-3355
  vibe ENG-3355 --dry-run
  vibe ENG-3355 --no-tmux
`

// Params holds the parsed CLI parameters.
type Params struct {
	JiraKey string
	DryRun  bool
	NoTmux  bool
}

// plan holds the execution plan.
type plan struct {
	jiraKey string
	useTmux bool
	prompt  string
}

func Run() int {
	fs := flag.NewFlagSet("vibe", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	showVersion := fs.Bool("version", false, "Print version and exit")
	dryRun := fs.Bool("dry-run", false, "Show what would happen without executing")
	noTmux := fs.Bool("no-tmux", false, "Skip tmux window creation")

	// Reorder os.Args so flags work in any position (e.g. "vibe ENG-3355 --dry-run").
	reordered := reorderArgs(os.Args[1:])

	if err := fs.Parse(reordered); err != nil {
		return 1
	}

	if *showVersion {
		fmt.Println(version.String())
		return 0
	}

	args := fs.Args()
	if len(args) == 0 {
		fs.Usage()
		return 1
	}

	jiraKey := strings.ToUpper(args[0])
	if !isValidJiraKey(jiraKey) {
		fmt.Fprintf(os.Stderr, "error: %q does not look like a valid Jira key (expected format: PROJ-123)\n", jiraKey)
		return 1
	}

	params := Params{
		JiraKey: jiraKey,
		DryRun:  *dryRun,
		NoTmux:  *noTmux,
	}

	return execute(params)
}

// reorderArgs moves flags before positional arguments so that Go's flag
// package can parse them regardless of position.
// e.g. ["ENG-3355", "--dry-run"] -> ["--dry-run", "ENG-3355"]
func reorderArgs(args []string) []string {
	var flags, positional []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		if strings.HasPrefix(a, "-") {
			flags = append(flags, a)
		} else {
			positional = append(positional, a)
		}
	}
	return append(flags, positional...)
}

func execute(params Params) int {
	p := &plan{
		jiraKey: params.JiraKey,
		useTmux: !params.NoTmux,
		prompt:  prompt.Bootstrap(params.JiraKey),
	}

	if params.DryRun {
		printDryRun(p)
		return 0
	}

	return run(p)
}

func printDryRun(p *plan) {
	fmt.Printf("[dry-run] Jira key:   %s\n", p.jiraKey)

	if p.useTmux {
		fmt.Printf("[dry-run] tmux:       create/switch to rightmost window %q\n", p.jiraKey)
	} else {
		fmt.Printf("[dry-run] tmux:       disabled\n")
	}

	fmt.Printf("[dry-run] opencode:   will launch with bootstrap prompt\n")
	fmt.Printf("\n[dry-run] prompt:\n%s\n", p.prompt)
}

func run(p *plan) int {
	// Build the opencode command.
	// Use the default interactive TUI command with a bootstrap prompt so the
	// tmux window stays open for the work session.
	opencodeArgs := []string{"--prompt", p.prompt}
	opencodeBin := "opencode"

	if p.useTmux {
		return runWithTmux(p.jiraKey, opencodeBin, opencodeArgs)
	}
	return runDirect(opencodeBin, opencodeArgs)
}

func runWithTmux(windowName, bin string, args []string) int {
	if !tmux.InsideTmux() {
		fmt.Fprintf(os.Stderr, "error: not inside a tmux session (use --no-tmux to skip)\n")
		return 1
	}

	// If the window already exists, switch to it.
	if tmux.HasWindow(windowName) {
		fmt.Printf("tmux window %q already exists, switching to it\n", windowName)
		if err := tmux.SwitchToWindow(windowName); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			return 1
		}
		return 0
	}

	// Build the full command to run in the new tmux window.
	cmdParts := buildShellCommand(bin, args)

	if err := tmux.NewWindow(windowName, "", cmdParts); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return 1
	}

	fmt.Printf("started work on %s in new rightmost tmux window %q\n", windowName, windowName)
	return 0
}

func runDirect(bin string, args []string) int {
	cmd := exec.Command(bin, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: running opencode: %s\n", err)
		return 1
	}
	return 0
}

// buildShellCommand constructs a shell-safe command string slice for tmux.
// The prompt needs to be properly quoted since it contains newlines and special chars.
func buildShellCommand(bin string, args []string) []string {
	parts := make([]string, 0, len(args)+1)
	parts = append(parts, bin)
	for _, a := range args {
		// Quote arguments that contain spaces, newlines, or special characters.
		if needsQuoting(a) {
			// Use single quotes, escaping any internal single quotes.
			escaped := strings.ReplaceAll(a, "'", "'\\''")
			parts = append(parts, "'"+escaped+"'")
		} else {
			parts = append(parts, a)
		}
	}
	return parts
}

func isValidJiraKey(key string) bool {
	parts := strings.SplitN(key, "-", 2)
	if len(parts) != 2 {
		return false
	}
	project := parts[0]
	number := parts[1]
	if len(project) == 0 || len(number) == 0 {
		return false
	}
	for _, r := range project {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	for _, r := range number {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// needsQuoting returns true if a string contains characters that need
// shell quoting (spaces, tabs, newlines, quotes, backslashes, dollar signs, backticks).
func needsQuoting(s string) bool {
	for _, r := range s {
		switch r {
		case ' ', '\t', '\n', '"', '\'', '\\', '$', '`':
			return true
		}
	}
	return false
}
