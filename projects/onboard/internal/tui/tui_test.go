package tui

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/amriksd/code/projects/onboard/internal/github"
	"github.com/amriksd/code/projects/onboard/internal/picker"
)

func TestOnlyTheSignedInUserIsTickedByDefault(t *testing.T) {
	m := New(Config{Root: t.TempDir()})

	next, _ := m.gotOwners(ownersMsg{logins: []string{"amrik", "acme"}})
	m = next.(Model)
	m.gotRepos(reposMsg{owner: "amrik", repos: []github.Repo{{Name: "dotfiles"}}})
	m.gotRepos(reposMsg{owner: "acme", repos: []github.Repo{{Name: "api"}}})

	want := []picker.Selection{{Owner: "amrik", Name: "dotfiles"}}
	if got := m.picker.Selected(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Selected() = %+v, want %+v", got, want)
	}
}

func TestOwnersGivenUpFrontAreAllTicked(t *testing.T) {
	m := New(Config{Root: t.TempDir(), Owners: []string{"acme", "globex"}})

	next, _ := m.gotOwners(ownersMsg{logins: []string{"acme", "globex"}})
	m = next.(Model)
	m.gotRepos(reposMsg{owner: "acme", repos: []github.Repo{{Name: "api"}}})
	m.gotRepos(reposMsg{owner: "globex", repos: []github.Repo{{Name: "site"}}})

	want := []picker.Selection{{Owner: "acme", Name: "api"}, {Owner: "globex", Name: "site"}}
	if got := m.picker.Selected(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Selected() = %+v, want %+v", got, want)
	}
}

func TestEnterWithNothingTickedStaysOnThePicker(t *testing.T) {
	m := New(Config{Root: t.TempDir()})
	m.picker.Owners = []*picker.Owner{{Login: "acme", Repos: []picker.Repo{{Name: "api"}}}}
	m.stage = picking

	next, _ := m.key("enter")

	if got := next.(Model).stage; got != picking {
		t.Fatalf("stage = %v, want picking", got)
	}
}

func TestEnterWithATickedRepoStartsCloning(t *testing.T) {
	m := New(Config{Root: t.TempDir()})
	m.picker.Owners = []*picker.Owner{{Login: "acme", Repos: []picker.Repo{{Name: "api", Selected: true}}}}
	m.stage = picking

	next, _ := m.key("enter")

	if got := next.(Model).stage; got != cloning {
		t.Fatalf("stage = %v, want cloning", got)
	}
}

func TestAFailedCloneIsReportedAsFailed(t *testing.T) {
	m := New(Config{Root: t.TempDir()})
	m.jobs = []picker.Selection{{Owner: "acme", Name: "api"}, {Owner: "acme", Name: "web"}}
	m.next = 2

	next, _ := m.gotClone(clonedMsg{job: 0, err: errTest})
	next, _ = next.(Model).gotClone(clonedMsg{job: 1})

	want := []Result{{Repo: picker.Selection{Owner: "acme", Name: "api"}, Err: errTest}}
	if got := next.(Model).Failed(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Failed() = %+v, want %+v", got, want)
	}
}

func TestToolsAreTickedWhenTheDotfilesAreMissing(t *testing.T) {
	m := New(Config{DotfilesDir: filepath.Join(t.TempDir(), "missing")})

	if !m.setupTools {
		t.Fatal("setupTools = false, want true")
	}
}

func TestToolsAreNotTickedWhenTheDotfilesAreAlreadyThere(t *testing.T) {
	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}

	m := New(Config{DotfilesDir: dir})

	if m.setupTools {
		t.Fatal("setupTools = true, want false")
	}
}

func TestTickingWorkSetupAlsoTicksTools(t *testing.T) {
	m := New(Config{DotfilesDir: t.TempDir()})
	m.setupTools = false
	m.setupCursor = 1

	next, _ := m.key(" ")

	got := next.(Model)
	if !got.setupTools || !got.setupWork {
		t.Fatalf("setupTools = %v, setupWork = %v, want both true", got.setupTools, got.setupWork)
	}
}

func TestEnterOnSetupWithNothingTickedGoesToThePicker(t *testing.T) {
	m := New(Config{DotfilesDir: t.TempDir()})
	m.setupTools = false

	next, _ := m.key("enter")

	if got := next.(Model).stage; got != picking {
		t.Fatalf("stage = %v, want picking", got)
	}
}

func TestAFailedBootstrapStaysOnSetupWithTheError(t *testing.T) {
	m := New(Config{DotfilesDir: t.TempDir()})

	next, _ := m.Update(bootstrapMsg{err: errTest})

	got := next.(Model)
	if got.stage != setup || got.bootstrapErr != errTest {
		t.Fatalf("stage = %v, bootstrapErr = %v, want setup and %v", got.stage, got.bootstrapErr, errTest)
	}
}
