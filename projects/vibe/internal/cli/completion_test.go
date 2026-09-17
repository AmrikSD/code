package cli

import (
	"reflect"
	"testing"
	"time"

	"github.com/amriksd/code/projects/vibe/internal/jira"
)

func TestCompleteCandidates(t *testing.T) {
	tickets := func() []candidate {
		return []candidate{
			{"ENG-4107", "To Do: bump graph-service"},
			{"SUP-17", "In Progress: printer on fire"},
		}
	}
	values := func(cs []candidate) []string {
		out := make([]string, 0, len(cs))
		for _, c := range cs {
			out = append(out, c.value)
		}
		return out
	}

	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{"empty line", nil, []string{"queue", "completion", "ENG-4107", "SUP-17"}},
		{"first word empty", []string{""}, []string{"queue", "completion", "ENG-4107", "SUP-17"}},
		{"ticket prefix", []string{"SUP"}, []string{"SUP-17"}},
		{"ticket prefix is case-insensitive", []string{"eng-4"}, []string{"ENG-4107"}},
		{"after a root flag", []string{"--dry-run", ""}, []string{"queue", "completion", "ENG-4107", "SUP-17"}},
		{"root flags", []string{"--"}, values(rootFlags)},
		{"root flag prefix", []string{"--n"}, []string{"--no-tmux"}},
		{"flags after a key", []string{"ENG-4107", "--d"}, []string{"--dry-run"}},
		{"nothing else after a key", []string{"ENG-4107", ""}, nil},
		{"queue flags", []string{"queue", "--"}, values(queueFlags)},
		{"queue takes no positionals", []string{"queue", ""}, nil},
		{"completion shells", []string{"completion", ""}, []string{"zsh", "bash"}},
		{"completion shell prefix", []string{"completion", "z"}, []string{"zsh"}},
		{"completion takes one shell", []string{"completion", "zsh", ""}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := values(completeCandidates(tt.words, tickets))
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("completeCandidates(%q) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}

func TestCompleteCandidatesDoesNotFetchTicketsForFlags(t *testing.T) {
	called := false
	tickets := func() []candidate {
		called = true
		return nil
	}
	for _, words := range [][]string{{"--"}, {"queue", "--l"}, {"ENG-1", "--"}, {"completion", ""}} {
		completeCandidates(words, tickets)
	}
	if called {
		t.Fatal("tickets fetched while completing something that is not a Jira key")
	}
}

func TestSortForCompletion(t *testing.T) {
	day := func(n int) time.Time { return time.Date(2026, 9, n, 0, 0, 0, 0, time.UTC) }
	issues := []jira.Issue{
		{Key: "ENG-1", Status: "To Do", Updated: day(10)},
		{Key: "ENG-2", Status: "In Progress", Updated: day(1)},
		{Key: "ENG-3", Status: "Review", Updated: day(12)},
		{Key: "ENG-4", Status: "In Progress", Updated: day(5)},
		{Key: "ENG-5", Status: "To Do", Updated: day(11)},
	}
	sortForCompletion(issues)
	var got []string
	for _, is := range issues {
		got = append(got, is.Key)
	}
	want := []string{"ENG-4", "ENG-2", "ENG-3", "ENG-5", "ENG-1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("order = %v, want %v", got, want)
	}
}
