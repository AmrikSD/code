package tui

import (
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

	next, _ := m.key("enter")

	if got := next.(Model).stage; got != picking {
		t.Fatalf("stage = %v, want picking", got)
	}
}

func TestEnterWithATickedRepoStartsCloning(t *testing.T) {
	m := New(Config{Root: t.TempDir()})
	m.picker.Owners = []*picker.Owner{{Login: "acme", Repos: []picker.Repo{{Name: "api", Selected: true}}}}

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
