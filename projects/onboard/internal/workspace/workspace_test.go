package workspace

import (
	"os"
	"path/filepath"
	"testing"
)

func TestClonedIsTrueWhenTheCheckoutHasAGitDir(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "acme", "api", ".git"), 0o755); err != nil {
		t.Fatal(err)
	}

	if !Cloned(root, "acme", "api") {
		t.Fatal("Cloned() = false, want true")
	}
}

func TestClonedIsFalseForAPlainDirectory(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "acme", "api"), 0o755); err != nil {
		t.Fatal(err)
	}

	if Cloned(root, "acme", "api") {
		t.Fatal("Cloned() = true, want false")
	}
}

func TestLastLineReturnsTheFinalLineOfGitOutput(t *testing.T) {
	got := lastLine("Cloning into 'api'...\nfatal: repository not found\n", "exit status 128")

	if got != "fatal: repository not found" {
		t.Fatalf("lastLine() = %q, want %q", got, "fatal: repository not found")
	}
}

func TestLastLineFallsBackWhenThereIsNoOutput(t *testing.T) {
	got := lastLine("", "exit status 128")

	if got != "exit status 128" {
		t.Fatalf("lastLine() = %q, want %q", got, "exit status 128")
	}
}
