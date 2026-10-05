// Package tui is the terminal UI: tick owners and repos, then watch them clone.
package tui

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/amriksd/code/projects/onboard/internal/github"
	"github.com/amriksd/code/projects/onboard/internal/picker"
	"github.com/amriksd/code/projects/onboard/internal/workspace"
)

const workers = 6

var (
	title   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("141"))
	bold    = lipgloss.NewStyle().Bold(true)
	dim     = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	good    = lipgloss.NewStyle().Foreground(lipgloss.Color("84"))
	bad     = lipgloss.NewStyle().Foreground(lipgloss.Color("203"))
	pointer = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	spinner = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
)

// Config is what the UI needs from outside.
type Config struct {
	// Root is the directory repos are cloned under, as <Root>/<owner>/<repo>.
	Root string
	// Owners limits the picker to these owners, all ticked. When empty the
	// signed-in user and their organisations are listed, with only the user ticked.
	Owners []string
	GH     github.Runner
	Clone  func(root, owner, repo, url string) error
}

// Result is the outcome of cloning one repo.
type Result struct {
	Repo picker.Selection
	Err  error
}

type stage int

const (
	picking stage = iota
	cloning
	done
)

type (
	ownersMsg struct {
		logins []string
		err    error
	}
	reposMsg struct {
		owner string
		repos []github.Repo
		err   error
	}
	clonedMsg struct {
		job int
		err error
	}
	registeredMsg struct{ err error }
	tickMsg       struct{}
)

// Model is the Bubble Tea model for the whole flow.
type Model struct {
	cfg         Config
	picker      *picker.Picker
	stage       stage
	width       int
	height      int
	offset      int
	frame       int
	discovering bool
	ticked      map[string]bool
	err         error

	jobs        []picker.Selection
	next        int
	active      map[int]bool
	Results     []Result
	RegisterErr error
}

// New builds the model in its initial, loading state.
func New(cfg Config) Model {
	if cfg.Clone == nil {
		cfg.Clone = workspace.Clone
	}
	return Model{cfg: cfg, picker: &picker.Picker{}, height: 24, width: 80, ticked: map[string]bool{}, active: map[int]bool{}}
}

// Run shows the UI and returns the model as it was when the user left.
func Run(cfg Config) (Model, error) {
	final, err := tea.NewProgram(New(cfg), tea.WithAltScreen()).Run()
	if err != nil {
		return Model{}, err
	}
	return final.(Model), nil
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tick(), func() tea.Msg {
		if len(m.cfg.Owners) > 0 {
			return ownersMsg{logins: m.cfg.Owners}
		}
		logins, err := github.Owners(m.cfg.GH)
		return ownersMsg{logins: logins, err: err}
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.scroll()
	case tickMsg:
		m.frame++
		return m, tick()
	case ownersMsg:
		return m.gotOwners(msg)
	case reposMsg:
		m.gotRepos(msg)
	case clonedMsg:
		return m.gotClone(msg)
	case registeredMsg:
		m.RegisterErr = msg.err
		m.stage = done
	case tea.KeyMsg:
		return m.key(msg.String())
	}
	return m, nil
}

func (m Model) gotOwners(msg ownersMsg) (tea.Model, tea.Cmd) {
	if msg.err != nil {
		m.err = msg.err
		return m, nil
	}
	var cmds []tea.Cmd
	for i, login := range msg.logins {
		m.picker.Owners = append(m.picker.Owners, &picker.Owner{Login: login, Loading: true})
		m.ticked[login] = len(m.cfg.Owners) > 0 || i == 0
		cmds = append(cmds, func() tea.Msg {
			repos, err := github.Repos(m.cfg.GH, login)
			return reposMsg{owner: login, repos: repos, err: err}
		})
	}
	return m, tea.Batch(cmds...)
}

func (m *Model) gotRepos(msg reposMsg) {
	for _, owner := range m.picker.Owners {
		if owner.Login != msg.owner {
			continue
		}
		owner.Loading = false
		owner.Err = msg.err
		for _, r := range msg.repos {
			owner.Repos = append(owner.Repos, picker.Repo{
				Name:   r.Name,
				SSHURL: r.SSHURL,
				Cloned: workspace.Cloned(m.cfg.Root, owner.Login, r.Name),
			})
		}
		if m.ticked[owner.Login] {
			owner.SetAll(true)
		}
	}
}

func (m Model) gotClone(msg clonedMsg) (tea.Model, tea.Cmd) {
	delete(m.active, msg.job)
	m.Results = append(m.Results, Result{Repo: m.jobs[msg.job], Err: msg.err})
	if m.next < len(m.jobs) {
		cmd := m.clone(m.next)
		m.next++
		return m, cmd
	}
	if len(m.Results) < len(m.jobs) {
		return m, nil
	}
	return m, func() tea.Msg {
		return registeredMsg{err: workspace.Register(m.cfg.Root, owners(m.jobs))}
	}
}

func (m Model) key(key string) (tea.Model, tea.Cmd) {
	if key == "ctrl+c" {
		return m, tea.Quit
	}
	switch m.stage {
	case cloning:
		return m, nil
	case done:
		if key == "q" || key == "enter" || key == "esc" {
			return m, tea.Quit
		}
		return m, nil
	}

	switch key {
	case "q", "esc":
		return m, tea.Quit
	case "up", "k":
		m.picker.Move(-1)
	case "down", "j":
		m.picker.Move(1)
	case "pgup":
		m.picker.Move(-m.page())
	case "pgdown":
		m.picker.Move(m.page())
	case "home", "g":
		m.picker.Move(-len(m.picker.Rows()))
	case "end", "G":
		m.picker.Move(len(m.picker.Rows()))
	case " ", "x":
		m.picker.Toggle()
	case "right", "l", "tab":
		m.picker.Expand()
	case "left", "h":
		m.picker.Collapse()
	case "a":
		m.picker.SetAll(true)
	case "n":
		m.picker.SetAll(false)
	case "enter":
		return m.start()
	}
	m.scroll()
	return m, nil
}

func (m Model) start() (tea.Model, tea.Cmd) {
	m.jobs = m.picker.Selected()
	if len(m.jobs) == 0 {
		return m, nil
	}
	m.stage = cloning
	var cmds []tea.Cmd
	for m.next < len(m.jobs) && m.next < workers {
		cmds = append(cmds, m.clone(m.next))
		m.next++
	}
	return m, tea.Batch(cmds...)
}

func (m Model) clone(job int) tea.Cmd {
	m.active[job] = true
	repo := m.jobs[job]
	return func() tea.Msg {
		return clonedMsg{job: job, err: m.cfg.Clone(m.cfg.Root, repo.Owner, repo.Name, repo.SSHURL)}
	}
}

func (m Model) page() int {
	if rows := m.height - 7; rows > 3 {
		return rows
	}
	return 3
}

func (m *Model) scroll() {
	if m.picker.Cursor < m.offset {
		m.offset = m.picker.Cursor
	}
	if m.picker.Cursor >= m.offset+m.page() {
		m.offset = m.picker.Cursor - m.page() + 1
	}
}

func (m Model) View() string {
	switch {
	case m.err != nil:
		return fmt.Sprintf("\n  %s\n\n  %s\n", bad.Render(m.err.Error()), dim.Render("q quit"))
	case m.stage == cloning:
		return m.viewCloning()
	case m.stage == done:
		return m.viewDone()
	default:
		return m.viewPicking()
	}
}

func (m Model) viewPicking() string {
	var b strings.Builder
	fmt.Fprintf(&b, "\n  %s %s\n\n", title.Render("Clone repositories into"), tilde(m.cfg.Root))

	if len(m.picker.Owners) == 0 {
		fmt.Fprintf(&b, "  %s Looking up your GitHub account and organisations\n", m.spin())
		return b.String()
	}

	rows := m.picker.Rows()
	end := m.offset + m.page()
	if end > len(rows) {
		end = len(rows)
	}
	for i := m.offset; i < end; i++ {
		cursor := "  "
		if i == m.picker.Cursor {
			cursor = pointer.Render("❯ ")
		}
		b.WriteString("  " + cursor + m.viewRow(rows[i]) + "\n")
	}

	fmt.Fprintf(&b, "\n  %s\n  %s\n",
		bold.Render(fmt.Sprintf("%d to clone", len(m.picker.Selected()))),
		dim.Render("space tick · → open · ← close · a all · n none · enter clone · q quit"))
	return b.String()
}

func (m Model) viewRow(row picker.Row) string {
	owner := m.picker.Owners[row.Owner]
	if row.Repo >= 0 {
		repo := owner.Repos[row.Repo]
		switch {
		case repo.Cloned:
			return "      " + dim.Render(" ✓  "+repo.Name+"  already cloned")
		case repo.Selected:
			return "      " + good.Render("[x]") + " " + repo.Name
		default:
			return "      [ ] " + repo.Name
		}
	}

	arrow := "▸"
	if owner.Expanded {
		arrow = "▾"
	}
	box := "[ ]"
	switch owner.State() {
	case picker.All:
		box = good.Render("[x]")
	case picker.Some:
		box = good.Render("[-]")
	}
	line := fmt.Sprintf("%s %s %s", arrow, box, bold.Render(owner.Login))
	switch {
	case owner.Loading:
		return line + "  " + dim.Render(m.spin()+" loading")
	case owner.Err != nil:
		return line + "  " + bad.Render(owner.Err.Error())
	}
	selected, cloned := owner.Count()
	return line + "  " + dim.Render(fmt.Sprintf("%d of %d ticked · %d already cloned", selected, len(owner.Repos)-cloned, cloned))
}

func (m Model) viewCloning() string {
	var b strings.Builder
	finished, total := len(m.Results), len(m.jobs)
	fmt.Fprintf(&b, "\n  %s\n\n", title.Render(fmt.Sprintf("Cloning %d of %d", finished, total)))

	const width = 40
	filled := width * finished / total
	fmt.Fprintf(&b, "  %s%s\n\n", good.Render(strings.Repeat("█", filled)), dim.Render(strings.Repeat("░", width-filled)))

	room := m.height - 8 - len(m.active)
	if room < 3 {
		room = 3
	}
	first := 0
	if len(m.Results) > room {
		first = len(m.Results) - room
	}
	for _, r := range m.Results[first:] {
		b.WriteString("  " + viewResult(r) + "\n")
	}
	for job := range m.jobs {
		if m.active[job] {
			fmt.Fprintf(&b, "  %s %s/%s\n", m.spin(), m.jobs[job].Owner, m.jobs[job].Name)
		}
	}
	return b.String()
}

func (m Model) viewDone() string {
	var b strings.Builder
	failed := m.Failed()
	fmt.Fprintf(&b, "\n  %s %s\n\n", title.Render(fmt.Sprintf("Cloned %d of %d into", len(m.Results)-len(failed), len(m.Results))), tilde(m.cfg.Root))
	for _, r := range failed {
		b.WriteString("  " + viewResult(r) + "\n")
	}
	if m.RegisterErr != nil {
		b.WriteString("  " + bad.Render(m.RegisterErr.Error()) + "\n")
	}
	b.WriteString("\n  " + dim.Render("q quit") + "\n")
	return b.String()
}

// Failed lists the repos that could not be cloned.
func (m Model) Failed() []Result {
	var failed []Result
	for _, r := range m.Results {
		if r.Err != nil {
			failed = append(failed, r)
		}
	}
	return failed
}

func viewResult(r Result) string {
	name := r.Repo.Owner + "/" + r.Repo.Name
	if r.Err != nil {
		return bad.Render("✗ "+name) + "  " + dim.Render(r.Err.Error())
	}
	return good.Render("✓") + " " + name
}

func (m Model) spin() string {
	return spinner[m.frame%len(spinner)]
}

func tilde(path string) string {
	if home, err := os.UserHomeDir(); err == nil && strings.HasPrefix(path, home) {
		return "~" + strings.TrimPrefix(path, home)
	}
	return path
}

func tick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(time.Time) tea.Msg { return tickMsg{} })
}

func owners(jobs []picker.Selection) []string {
	var out []string
	seen := map[string]bool{}
	for _, j := range jobs {
		if !seen[j.Owner] {
			seen[j.Owner] = true
			out = append(out, j.Owner)
		}
	}
	return out
}
