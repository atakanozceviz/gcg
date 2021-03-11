# Templating

gcg uses Go’s `text/template` library as the basis for the templating.

For an in-depth look into Go Templates, check the official [Go docs](http://golang.org/pkg/text/template/).

## Template Fields and Functions

### IssuesByMilestone & IssuesByTag

`IssuesByMilestone` and `IssuesByTag` are arrays of issues and pull requests that are grouped by related labels.

`IssuesByMilestone` only available if **milestone** setting is set and `IssuesByTag` only available if **since-tag** and/or **until-tag** setting is set.

#### Example

```
{{range .IssuesByMilestone}}
{{.Title}}
    {{range .Issues}}
    {{- end -}}
{{end}}
```

### SinceTag & UntilTag

`SinceTag` and `UntilTag` are tag titles which can be set by **since-tag** and **until-tag** settings and only available if they are set.

#### Example

```
{{.SinceTag}}-{{.UntilTag}}
```

### [Repository](https://pkg.go.dev/github.com/google/go-github/github#Repository)

Repository represents a GitHub repository.

#### Functions

- [GetAllowMergeCommit](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetAllowMergeCommit)
- [GetAllowRebaseMerge](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetAllowRebaseMerge)
- [GetAllowSquashMerge](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetAllowSquashMerge)
- [GetArchiveURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetArchiveURL)
- [GetArchived](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetArchived)
- [GetAssigneesURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetAssigneesURL)
- [GetAutoInit](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetAutoInit)
- [GetBlobsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetBlobsURL)
- [GetBranchesURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetBranchesURL)
- [GetCloneURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetCloneURL)
- [GetCodeOfConduct](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetCodeOfConduct)
- [GetCollaboratorsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetCollaboratorsURL)
- [GetCommentsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetCommentsURL)
- [GetCommitsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetCommitsURL)
- [GetCompareURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetCompareURL)
- [GetContentsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetContentsURL)
- [GetContributorsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetContributorsURL)
- [GetCreatedAt](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetCreatedAt)
- [GetDefaultBranch](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetDefaultBranch)
- [GetDeploymentsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetDeploymentsURL)
- [GetDescription](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetDescription)
- [GetDisabled](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetDisabled)
- [GetDownloadsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetDownloadsURL)
- [GetEventsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetEventsURL)
- [GetFork](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetFork)
- [GetForksCount](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetForksCount)
- [GetForksURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetForksURL)
- [GetFullName](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetFullName)
- [GetGitCommitsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetGitCommitsURL)
- [GetGitRefsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetGitRefsURL)
- [GetGitTagsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetGitTagsURL)
- [GetGitURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetGitURL)
- [GetGitignoreTemplate](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetGitignoreTemplate)
- [GetHTMLURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetHTMLURL)
- [GetHasDownloads](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetHasDownloads)
- [GetHasIssues](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetHasIssues)
- [GetHasPages](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetHasPages)
- [GetHasProjects](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetHasProjects)
- [GetHasWiki](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetHasWiki)
- [GetHomepage](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetHomepage)
- [GetHooksURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetHooksURL)
- [GetID](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetID)
- [GetIssueCommentURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetIssueCommentURL)
- [GetIssueEventsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetIssueEventsURL)
- [GetIssuesURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetIssuesURL)
- [GetKeysURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetKeysURL)
- [GetLabelsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetLabelsURL)
- [GetLanguage](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetLanguage)
- [GetLanguagesURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetLanguagesURL)
- [GetLicense](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetLicense)
- [GetLicenseTemplate](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetLicenseTemplate)
- [GetMasterBranch](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetMasterBranch)
- [GetMergesURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetMergesURL)
- [GetMilestonesURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetMilestonesURL)
- [GetMirrorURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetMirrorURL)
- [GetName](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetName)
- [GetNetworkCount](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetNetworkCount)
- [GetNodeID](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetNodeID)
- [GetNotificationsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetNotificationsURL)
- [GetOpenIssuesCount](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetOpenIssuesCount)
- [GetOrganization](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetOrganization)
- [GetOwner](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetOwner)
- [GetParent](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetParent)
- [GetPermissions](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetPermissions)
- [GetPrivate](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetPrivate)
- [GetPullsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetPullsURL)
- [GetPushedAt](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetPushedAt)
- [GetReleasesURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetReleasesURL)
- [GetSSHURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetSSHURL)
- [GetSVNURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetSVNURL)
- [GetSize](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetSize)
- [GetSource](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetSource)
- [GetStargazersCount](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetStargazersCount)
- [GetStargazersURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetStargazersURL)
- [GetStatusesURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetStatusesURL)
- [GetSubscribersCount](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetSubscribersCount)
- [GetSubscribersURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetSubscribersURL)
- [GetSubscriptionURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetSubscriptionURL)
- [GetTagsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetTagsURL)
- [GetTeamID](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetTeamID)
- [GetTeamsURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetTeamsURL)
- [GetTreesURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetTreesURL)
- [GetURL](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetURL)
- [GetUpdatedAt](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetUpdatedAt)
- [GetWatchersCount](https://pkg.go.dev/github.com/google/go-github/github#Repository.GetWatchersCount)

#### Example

```
[{{.Repository.GetName}}]({{.Repository.GetHTMLURL}})
```

### [Milestone](https://pkg.go.dev/github.com/google/go-github/github#Milestone)

Milestone represents a GitHub repository milestone.
Only available if **milestone** setting is set.

#### Functions

- [GetClosedAt](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetClosedAt)
- [GetClosedIssues](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetClosedIssues)
- [GetCreatedAt](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetCreatedAt)
- [GetCreator](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetCreator)
- [GetDescription](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetDescription)
- [GetDueOn](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetDueOn)
- [GetHTMLURL](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetHTMLURL)
- [GetID](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetID)
- [GetLabelsURL](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetLabelsURL)
- [GetNodeID](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetNodeID)
- [GetNumber](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetNumber)
- [GetOpenIssues](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetOpenIssues)
- [GetState](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetState)
- [GetTitle](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetTitle)
- [GetURL](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetURL)
- [GetUpdatedAt](https://pkg.go.dev/github.com/google/go-github/github#Milestone.GetUpdatedAt)

#### Example

```
{{.Milestone.GetTitle}} ({{.Milestone.GetClosedAt.Format "2006-01-02"}})
```

### [Commit](https://pkg.go.dev/github.com/google/go-github/github#Commit) (SinceTagCommit & UntilTagCommit )

Commit represents a GitHub commit.
Only available if **since-tag** and/or **until-tag** setting is set.

#### Functions

- [GetAuthor](https://pkg.go.dev/github.com/google/go-github/github#Commit.GetAuthor)
- [GetCommentCount](https://pkg.go.dev/github.com/google/go-github/github#Commit.GetCommentCount)
- [GetCommitter](https://pkg.go.dev/github.com/google/go-github/github#Commit.GetCommitter)
- [GetHTMLURL](https://pkg.go.dev/github.com/google/go-github/github#Commit.GetHTMLURL)
- [GetMessage](https://pkg.go.dev/github.com/google/go-github/github#Commit.GetMessage)
- [GetNodeID](https://pkg.go.dev/github.com/google/go-github/github#Commit.GetNodeID)
- [GetSHA](https://pkg.go.dev/github.com/google/go-github/github#Commit.GetSHA)
- [GetStats](https://pkg.go.dev/github.com/google/go-github/github#Commit.GetStats)
- [GetTree](https://pkg.go.dev/github.com/google/go-github/github#Commit.GetTree)
- [GetURL](https://pkg.go.dev/github.com/google/go-github/github#Commit.GetURL)
- [GetVerification](https://pkg.go.dev/github.com/google/go-github/github#Commit.GetVerification)

#### Example

```
[{{.SinceTag}}]({{.SinceTagCommit.GetHTMLURL}}) - [{{.UntilTag}}]({{.UntilTagCommit.GetHTMLURL}})
```

### [Issue](https://pkg.go.dev/github.com/google/go-github/github#Issue)

Issue represents a GitHub issue on a repository.

Note: As far as the GitHub API is concerned, every pull request is an issue, but not every issue is a pull request. Some endpoints, events, and webhooks may also return pull requests via this struct. If PullRequestLinks is nil, this is an issue, and if PullRequestLinks is not nil, this is a pull request. The IsPullRequest helper method can be used to check that.

#### Functions

- [GetActiveLockReason](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetActiveLockReason)
- [GetAssignee](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetAssignee)
- [GetBody](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetBody)
- [GetClosedAt](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetClosedAt)
- [GetClosedBy](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetClosedBy)
- [GetComments](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetComments)
- [GetCommentsURL](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetCommentsURL)
- [GetCreatedAt](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetCreatedAt)
- [GetEventsURL](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetEventsURL)
- [GetHTMLURL](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetHTMLURL)
- [GetID](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetID)
- [GetLabelsURL](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetLabelsURL)
- [GetLocked](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetLocked)
- [GetMilestone](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetMilestone)
- [GetNodeID](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetNodeID)
- [GetNumber](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetNumber)
- [GetPullRequestLinks](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetPullRequestLinks)
- [GetReactions](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetReactions)
- [GetRepository](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetRepository)
- [GetRepositoryURL](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetRepositoryURL)
- [GetState](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetState)
- [GetTitle](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetTitle)
- [GetURL](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetURL)
- [GetUpdatedAt](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetUpdatedAt)
- [GetUser](https://pkg.go.dev/github.com/google/go-github/github#Issue.GetUser)
- [IsPullRequest](https://pkg.go.dev/github.com/google/go-github/github#Issue.IsPullRequest)

#### Example

```
{{range .IssuesByMilestone}}
{{.Title}}
{{range .Issues}}
{{if .IsPullRequest -}}
- PR [\#{{.GetNumber}}]({{.GetHTMLURL}}): {{.GetTitle}} (by [{{.GetUser.GetLogin}}]({{.GetUser.GetHTMLURL}}))
{{- else -}}
- ISSUE [\#{{.GetNumber}}]({{.GetHTMLURL}}): {{.GetTitle}}
{{- end -}}
{{end}}
```
