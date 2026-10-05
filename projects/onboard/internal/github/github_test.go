package github

import (
	"reflect"
	"testing"
)

func TestOwnersListsTheUserBeforeTheirOrgs(t *testing.T) {
	run := func(args ...string) ([]byte, error) {
		if args[1] == "user" {
			return []byte("amrik\n"), nil
		}
		return []byte("acme\nglobex\n"), nil
	}

	got, err := Owners(run)

	want := []string{"amrik", "acme", "globex"}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("Owners() = %v, %v, want %v", got, err, want)
	}
}

func TestParseReposSortsByNameIgnoringCase(t *testing.T) {
	out := []byte(`[
		{"name": "web", "sshUrl": "git@github.com:acme/web.git"},
		{"name": "API", "sshUrl": "git@github.com:acme/API.git"},
		{"name": "docs", "sshUrl": "git@github.com:acme/docs.git"}
	]`)

	got, err := parseRepos(out)

	want := []Repo{
		{Name: "API", SSHURL: "git@github.com:acme/API.git"},
		{Name: "docs", SSHURL: "git@github.com:acme/docs.git"},
		{Name: "web", SSHURL: "git@github.com:acme/web.git"},
	}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("parseRepos() = %+v, %v, want %+v", got, err, want)
	}
}

func TestParseReposRejectsOutputThatIsNotJSON(t *testing.T) {
	_, err := parseRepos([]byte("not json"))

	if err == nil {
		t.Fatal("parseRepos() error = nil, want an error")
	}
}
