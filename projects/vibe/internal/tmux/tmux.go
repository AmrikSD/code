package tmux

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// InsideTmux returns true if the current process is running inside tmux.
func InsideTmux() bool {
	return os.Getenv("TMUX") != ""
}

// HasWindow checks if a tmux window with the given name exists
// in the current session.
func HasWindow(name string) bool {
	cmd := exec.Command("tmux", "list-windows", "-F", "#{window_name}")
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line == name {
			return true
		}
	}
	return false
}

// SwitchToWindow switches to an existing tmux window by name.
func SwitchToWindow(name string) error {
	cmd := exec.Command("tmux", "select-window", "-t", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("switching to tmux window %q: %w", name, err)
	}
	return nil
}

// NewWindow creates a new tmux window with the given name and working directory,
// then runs the specified command inside it. The window is appended to the end
// of the current session's window list so it opens as a new tab on the right.
func NewWindow(name, workdir string, command []string) error {
	args := []string{"new-window", "-a", "-n", name}
	if workdir != "" {
		args = append(args, "-c", workdir)
	}
	if len(command) > 0 {
		// Join command parts into a single shell string for tmux.
		args = append(args, strings.Join(command, " "))
	}

	cmd := exec.Command("tmux", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("creating tmux window %q: %w", name, err)
	}
	return nil
}

// SendKeys sends keystrokes to a tmux window by name.
// This is useful for sending a command to an existing window.
func SendKeys(windowName string, keys string) error {
	cmd := exec.Command("tmux", "send-keys", "-t", windowName, keys, "Enter")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("sending keys to tmux window %q: %w", windowName, err)
	}
	return nil
}
