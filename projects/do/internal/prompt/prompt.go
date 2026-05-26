package prompt

import "fmt"

// Bootstrap returns the default OpenCode bootstrap prompt for a given Jira key.
func Bootstrap(jiraKey string) string {
	return fmt.Sprintf(`Use the Jira CLI skill to read issue %s and any immediately relevant linked context.

Summarize the task:
1. What is the goal?
2. What are the acceptance criteria?
3. What constraints or technical requirements exist?
4. What is unclear or missing?

If requirements are ambiguous or incomplete, stop and ask focused clarifying questions before making any changes.

If the task is clear enough, use the GitHub CLI skill to locate the relevant repository under ~/code, inspect the codebase, implement the smallest correct solution, and verify it.

Do not run expensive build steps unless explicitly requested.
Do not make assumptions about product behavior that are not supported by the ticket or codebase.`, jiraKey)
}
