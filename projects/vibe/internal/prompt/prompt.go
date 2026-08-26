package prompt

import (
	"fmt"
	"strings"
)

// serviceDeskProjects lists Jira project keys that are service desks rather
// than engineering backlogs. Tickets in these projects get the support
// bootstrap prompt instead of the implementation one.
var serviceDeskProjects = map[string]bool{
	"SUP": true,
}

// IsServiceDesk reports whether a Jira key belongs to a service desk project.
func IsServiceDesk(jiraKey string) bool {
	project, _, _ := strings.Cut(jiraKey, "-")
	return serviceDeskProjects[project]
}

// Bootstrap returns the OpenCode bootstrap prompt appropriate for a Jira key:
// the support-desk prompt for service desk projects, the implementation
// prompt for everything else.
func Bootstrap(jiraKey string) string {
	if IsServiceDesk(jiraKey) {
		return ServiceDesk(jiraKey)
	}
	return Engineering(jiraKey)
}

// Engineering returns the bootstrap prompt for an engineering ticket that is
// expected to end in a code change.
func Engineering(jiraKey string) string {
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

// ServiceDesk returns the bootstrap prompt for a support-desk request: triage
// first, act by request type, and draft a reply to the requester.
func ServiceDesk(jiraKey string) string {
	return fmt.Sprintf(`You are working a Platform support-desk request. Use the Jira CLI skill to read %[1]s in full, including every comment (jira issue view %[1]s --comments 20), plus any linked issues, PRs, or docs.

Before doing anything else, load relevant skills by exact name:
- jira-cli (reading the ticket and comments, replying to the requester)
- github-cli (PR reviews and repository changes)
- kaas-clusters (anything touching KaaS clusters or tenants)
- confluence-cli (runbooks and reference docs)

Context: requests in this queue are raised by internal engineers, usually via Slack (expect Slack-style links like <url|text> and @mentions), and are handled by the Platform team. The requester is the ticket reporter, not the assignee. Comments on the ticket sync back to the requester.

Triage first and summarize in a few lines:
1. Who is asking, and what do they actually need (the ask behind the ask)?
2. Which system(s) does it touch (Cloudflare, KaaS, AWS/TGW, GCP, Warpstream, Postmark, Humio, CI, ...)?
3. What kind of request is it: PR review, access request, question, change request, or incident?
4. What has already happened in the comments, and what is still outstanding?
5. Is this ours, or should it be redirected (and to whom)?

Then act by request type:
- PR review: use gh to read the diff and CI status, review for correctness and blast radius, and post the review on GitHub. Report the outcome on the ticket.
- Question: find the authoritative answer in the codebase, cluster config, or Confluence. Do not guess; say where the answer came from.
- Change request: locate the relevant repository under ~/code, make the smallest correct change on a branch, open a PR, and link it on the ticket.
- Access request: identify where the access is granted, what approval is required, and the exact steps.
- Incident: gather evidence first (what is failing, since when, blast radius) before proposing a fix.

Communication:
- Draft a short, friendly reply to the requester as a Jira comment. Lead with the answer or current status, then the detail.
- Show me the drafted comment and ask using the questions workflow before posting it, before transitioning the ticket, and before changing any shared infrastructure.
- When the request is fully resolved, move the ticket to Done with a closing comment.

If the request is ambiguous after reading the ticket, comments, and linked context, draft focused clarifying questions for the requester rather than guessing.

Do not run expensive build steps unless explicitly requested.`, jiraKey)
}
