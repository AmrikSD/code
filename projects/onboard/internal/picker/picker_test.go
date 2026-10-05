package picker

import (
	"reflect"
	"testing"
)

func TestToggleOnOwnerSelectsEveryRepoUnderIt(t *testing.T) {
	p := &Picker{Owners: []*Owner{
		{Login: "acme", Repos: []Repo{{Name: "api"}, {Name: "web"}}},
	}}

	p.Toggle()

	want := []Repo{{Name: "api", Selected: true}, {Name: "web", Selected: true}}
	if !reflect.DeepEqual(p.Owners[0].Repos, want) {
		t.Fatalf("repos = %+v, want %+v", p.Owners[0].Repos, want)
	}
}

func TestToggleOnFullySelectedOwnerClearsEveryRepoUnderIt(t *testing.T) {
	p := &Picker{Owners: []*Owner{
		{Login: "acme", Repos: []Repo{{Name: "api", Selected: true}, {Name: "web", Selected: true}}},
	}}

	p.Toggle()

	want := []Repo{{Name: "api"}, {Name: "web"}}
	if !reflect.DeepEqual(p.Owners[0].Repos, want) {
		t.Fatalf("repos = %+v, want %+v", p.Owners[0].Repos, want)
	}
}

func TestToggleOnPartlySelectedOwnerSelectsTheRest(t *testing.T) {
	p := &Picker{Owners: []*Owner{
		{Login: "acme", Repos: []Repo{{Name: "api", Selected: true}, {Name: "web"}}},
	}}

	p.Toggle()

	if got := p.Owners[0].State(); got != All {
		t.Fatalf("State() = %v, want All", got)
	}
}

func TestToggleOnRepoFlipsOnlyThatRepo(t *testing.T) {
	p := &Picker{
		Owners: []*Owner{{Login: "acme", Expanded: true, Repos: []Repo{{Name: "api"}, {Name: "web"}}}},
		Cursor: 2,
	}

	p.Toggle()

	want := []Repo{{Name: "api"}, {Name: "web", Selected: true}}
	if !reflect.DeepEqual(p.Owners[0].Repos, want) {
		t.Fatalf("repos = %+v, want %+v", p.Owners[0].Repos, want)
	}
}

func TestClonedReposAreNeverSelected(t *testing.T) {
	p := &Picker{Owners: []*Owner{
		{Login: "acme", Repos: []Repo{{Name: "api", Cloned: true}, {Name: "web"}}},
	}}

	p.Toggle()

	want := []Selection{{Owner: "acme", Name: "web"}}
	if got := p.Selected(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Selected() = %+v, want %+v", got, want)
	}
}

func TestStateIgnoresClonedRepos(t *testing.T) {
	o := &Owner{Repos: []Repo{{Name: "api", Cloned: true}, {Name: "web", Selected: true}}}

	if got := o.State(); got != All {
		t.Fatalf("State() = %v, want All", got)
	}
}

func TestStateIsSomeWhenOnlyPartIsSelected(t *testing.T) {
	o := &Owner{Repos: []Repo{{Name: "api", Selected: true}, {Name: "web"}}}

	if got := o.State(); got != Some {
		t.Fatalf("State() = %v, want Some", got)
	}
}

func TestRowsOnlyIncludeReposOfExpandedOwners(t *testing.T) {
	p := &Picker{Owners: []*Owner{
		{Login: "acme", Expanded: true, Repos: []Repo{{Name: "api"}, {Name: "web"}}},
		{Login: "globex", Repos: []Repo{{Name: "site"}}},
	}}

	want := []Row{{Owner: 0, Repo: -1}, {Owner: 0, Repo: 0}, {Owner: 0, Repo: 1}, {Owner: 1, Repo: -1}}
	if got := p.Rows(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Rows() = %+v, want %+v", got, want)
	}
}

func TestMoveStopsAtTheLastRow(t *testing.T) {
	p := &Picker{Owners: []*Owner{{Login: "acme"}, {Login: "globex"}}}

	p.Move(5)

	if p.Cursor != 1 {
		t.Fatalf("Cursor = %d, want 1", p.Cursor)
	}
}

func TestCollapseFromRepoRowMovesCursorToItsOwner(t *testing.T) {
	p := &Picker{
		Owners: []*Owner{
			{Login: "acme"},
			{Login: "globex", Expanded: true, Repos: []Repo{{Name: "site"}, {Name: "shop"}}},
		},
		Cursor: 3,
	}

	p.Collapse()

	if p.Cursor != 1 || p.Owners[1].Expanded {
		t.Fatalf("Cursor = %d, Expanded = %v, want 1 and false", p.Cursor, p.Owners[1].Expanded)
	}
}

func TestSelectedListsChosenReposInDisplayOrder(t *testing.T) {
	p := &Picker{Owners: []*Owner{
		{Login: "acme", Repos: []Repo{{Name: "api", SSHURL: "git@github.com:acme/api.git", Selected: true}, {Name: "web"}}},
		{Login: "globex", Repos: []Repo{{Name: "site", SSHURL: "git@github.com:globex/site.git", Selected: true}}},
	}}

	want := []Selection{
		{Owner: "acme", Name: "api", SSHURL: "git@github.com:acme/api.git"},
		{Owner: "globex", Name: "site", SSHURL: "git@github.com:globex/site.git"},
	}
	if got := p.Selected(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Selected() = %+v, want %+v", got, want)
	}
}
