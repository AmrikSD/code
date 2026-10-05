package machine

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestHasDotfilesIsTrueForACheckout(t *testing.T) {
	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}

	if !HasDotfiles(dir) {
		t.Fatal("HasDotfiles() = false, want true")
	}
}

func TestHasDotfilesIsFalseForAMissingDirectory(t *testing.T) {
	if HasDotfiles(filepath.Join(t.TempDir(), "missing")) {
		t.Fatal("HasDotfiles() = true, want false")
	}
}

func TestBootstrapPassesRepoDirAndWorkModeToTheScript(t *testing.T) {
	cmd := Bootstrap("amrik/dotfiles", "/home/amrik/.dotfiles", true)

	want := []string{"bash", "-c", script, "bootstrap", "amrik/dotfiles", "/home/amrik/.dotfiles", "work"}
	if !reflect.DeepEqual(cmd.Args, want) {
		t.Fatalf("Args = %q, want %q", cmd.Args, want)
	}
}

func TestBootstrapWithoutWorkRunsInPersonalMode(t *testing.T) {
	cmd := Bootstrap("amrik/dotfiles", "/home/amrik/.dotfiles", false)

	if got := cmd.Args[len(cmd.Args)-1]; got != "personal" {
		t.Fatalf("mode = %q, want %q", got, "personal")
	}
}
