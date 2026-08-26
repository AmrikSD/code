// Package jira is a thin wrapper around the ankitpokhrel/jira-cli binary
// ("jira"). It shells out rather than talking to the REST API directly so that
// authentication, server config, and custom fields are all handled by the CLI
// the user already has configured.
package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Issue is the subset of a Jira issue that the queue view needs.
type Issue struct {
	Key      string
	Summary  string
	Status   string
	Assignee string
	Reporter string
	Priority string
	Comments int
	Created  time.Time
}

// rawIssue mirrors the shape of `jira issue list --raw`.
type rawIssue struct {
	Key    string `json:"key"`
	Fields struct {
		Summary  string `json:"summary"`
		Created  string `json:"created"`
		Status   named  `json:"status"`
		Assignee *user  `json:"assignee"`
		Reporter *user  `json:"reporter"`
		Priority *named `json:"priority"`
		Comment  *struct {
			Total int `json:"total"`
		} `json:"comment"`
	} `json:"fields"`
}

type named struct {
	Name string `json:"name"`
}

type user struct {
	DisplayName string `json:"displayName"`
}

// Search runs a JQL query and returns matching issues, oldest first.
func Search(jql string, limit int) ([]Issue, error) {
	out, err := run(
		"issue", "list",
		"-q", jql,
		"--order-by", "created",
		"--reverse",
		"--raw",
		"--paginate", fmt.Sprintf("0:%d", limit),
	)
	if err != nil {
		return nil, err
	}

	out = bytes.TrimSpace(out)
	if len(out) == 0 {
		return nil, nil
	}

	var raw []rawIssue
	if err := json.Unmarshal(out, &raw); err != nil {
		return nil, fmt.Errorf("parsing jira output: %w", err)
	}

	issues := make([]Issue, 0, len(raw))
	for _, r := range raw {
		is := Issue{
			Key:     r.Key,
			Summary: r.Fields.Summary,
			Status:  r.Fields.Status.Name,
		}
		if r.Fields.Assignee != nil {
			is.Assignee = r.Fields.Assignee.DisplayName
		}
		if r.Fields.Reporter != nil {
			is.Reporter = r.Fields.Reporter.DisplayName
		}
		if r.Fields.Priority != nil {
			is.Priority = r.Fields.Priority.Name
		}
		if r.Fields.Comment != nil {
			is.Comments = r.Fields.Comment.Total
		}
		if t, err := time.Parse("2006-01-02T15:04:05.000-0700", r.Fields.Created); err == nil {
			is.Created = t
		}
		issues = append(issues, is)
	}
	return issues, nil
}

// Me returns the login (email) of the configured Jira user.
func Me() (string, error) {
	out, err := run("me")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// Assign assigns an issue to the given user (email or display name).
func Assign(key, assignee string) error {
	_, err := run("issue", "assign", key, assignee)
	return err
}

// Move transitions an issue to the given status name.
func Move(key, status string) error {
	_, err := run("issue", "move", key, status)
	return err
}

// run executes the jira CLI and returns stdout. Stderr is captured and folded
// into the returned error so callers get the CLI's own message on failure.
func run(args ...string) ([]byte, error) {
	cmd := exec.Command("jira", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = err.Error()
		}
		return nil, fmt.Errorf("jira %s: %s", strings.Join(args[:min(2, len(args))], " "), msg)
	}
	return stdout.Bytes(), nil
}
