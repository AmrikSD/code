package cli

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestAvailableStatesParsesQuotedStates(t *testing.T) {
	err := errors.New("jira issue move: invalid transition state \"In Progress\"\nAvailable states for issue SUP-136: 'Start', 'In review'")

	got := availableStates(err)
	want := []string{"Start", "In review"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("availableStates() = %v, want %v", got, want)
	}
}

func TestParseStatePreferences(t *testing.T) {
	got := parseStatePreferences("In Progress, Start, In review, start")
	want := []string{"In Progress", "Start", "In review"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("parseStatePreferences() = %v, want %v", got, want)
	}

	got = parseStatePreferences(" , , ")
	want = []string{"In Progress"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("parseStatePreferences() empty = %v, want %v", got, want)
	}
}

func TestPickPreferredStatePrefersConfiguredOrder(t *testing.T) {
	state, ok := pickPreferredState([]string{"In Progress", "Start", "In review"}, []string{"Start", "In review"})
	if !ok {
		t.Fatal("pickPreferredState() returned no state")
	}
	if state != "Start" {
		t.Fatalf("pickPreferredState() = %q, want %q", state, "Start")
	}

	state, ok = pickPreferredState([]string{"In Progress", "Start", "In review"}, []string{"In review", "In Progress"})
	if !ok {
		t.Fatal("pickPreferredState() returned no state")
	}
	if state != "In Progress" {
		t.Fatalf("pickPreferredState() = %q, want %q", state, "In Progress")
	}
}

func TestMoveWithFallbackUsesAvailableState(t *testing.T) {
	var called []string
	move := func(_ string, status string) error {
		called = append(called, status)
		if status == "In Progress" {
			return errors.New("invalid transition state \"In Progress\". Available states for issue SUP-136: 'Start', 'In review'")
		}
		if status == "Start" {
			return nil
		}
		return fmt.Errorf("unexpected status %q", status)
	}

	got, err := moveWithFallback(move, "SUP-136", []string{"In Progress", "Start", "In review"})
	if err != nil {
		t.Fatalf("moveWithFallback() error = %v", err)
	}
	if got != "Start" {
		t.Fatalf("moveWithFallback() state = %q, want %q", got, "Start")
	}
	wantCalls := []string{"In Progress", "Start"}
	if !reflect.DeepEqual(called, wantCalls) {
		t.Fatalf("move calls = %v, want %v", called, wantCalls)
	}
}
