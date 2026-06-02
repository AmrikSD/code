package prompt

import "fmt"

// Bootstrap returns the default OpenCode bootstrap prompt for a given Jira key.
func Bootstrap(jiraKey string) string {
	return fmt.Sprintf(`Use the Jira CLI skill to read issue %s and any immediately relevant linked context.

Before doing anything else, load relevant skills by exact name:
- jira-cli (for ticket analysis)
- github-cli (for repository/PR work)
- confluence-cli (when the ticket or linked material points to Confluence docs)
- kaas-clusters (only if the issue touches KaaS clusters/tenants)

Summarize the task:
1. What is the goal?
2. What are the acceptance criteria?
3. What constraints or technical requirements exist?
4. What is unclear or missing?

If the Jira issue is documentation-based, links to Confluence, or appears to rely on design docs or rollout notes, use the Confluence CLI skill to read the relevant pages before deciding on implementation.

If requirements are ambiguous or incomplete, stop and ask focused clarifying questions using the questions workflow before making any changes.

If the task is clear enough, use the GitHub CLI skill to locate the relevant repository under ~/code, inspect the codebase, implement the smallest correct solution, and verify it.

Do not run expensive build steps unless explicitly requested.
Do not make assumptions about product behavior that are not supported by the ticket or codebase.`, jiraKey)
}
