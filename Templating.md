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

### [Repository](https://godoc.org/github.com/google/go-github/github#Repository)

Repository represents a GitHub repository.

#### Functions

- [GetAllowMergeCommit](https://godoc.org/github.com/google/go-github/github#Repository.GetAllowMergeCommit)
- [GetAllowRebaseMerge](https://godoc.org/github.com/google/go-github/github#Repository.GetAllowRebaseMerge)
- [GetAllowSquashMerge](https://godoc.org/github.com/google/go-github/github#Repository.GetAllowSquashMerge)
- [GetArchiveURL](https://godoc.org/github.com/google/go-github/github#Repository.GetArchiveURL)
- [GetArchived](https://godoc.org/github.com/google/go-github/github#Repository.GetArchived)
- [GetAssigneesURL](https://godoc.org/github.com/google/go-github/github#Repository.GetAssigneesURL)
- [GetAutoInit](https://godoc.org/github.com/google/go-github/github#Repository.GetAutoInit)
- [GetBlobsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetBlobsURL)
- [GetBranchesURL](https://godoc.org/github.com/google/go-github/github#Repository.GetBranchesURL)
- [GetCloneURL](https://godoc.org/github.com/google/go-github/github#Repository.GetCloneURL)
- [GetCodeOfConduct](https://godoc.org/github.com/google/go-github/github#Repository.GetCodeOfConduct)
- [GetCollaboratorsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetCollaboratorsURL)
- [GetCommentsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetCommentsURL)
- [GetCommitsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetCommitsURL)
- [GetCompareURL](https://godoc.org/github.com/google/go-github/github#Repository.GetCompareURL)
- [GetContentsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetContentsURL)
- [GetContributorsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetContributorsURL)
- [GetCreatedAt](https://godoc.org/github.com/google/go-github/github#Repository.GetCreatedAt)
- [GetDefaultBranch](https://godoc.org/github.com/google/go-github/github#Repository.GetDefaultBranch)
- [GetDeploymentsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetDeploymentsURL)
- [GetDescription](https://godoc.org/github.com/google/go-github/github#Repository.GetDescription)
- [GetDisabled](https://godoc.org/github.com/google/go-github/github#Repository.GetDisabled)
- [GetDownloadsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetDownloadsURL)
- [GetEventsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetEventsURL)
- [GetFork](https://godoc.org/github.com/google/go-github/github#Repository.GetFork)
- [GetForksCount](https://godoc.org/github.com/google/go-github/github#Repository.GetForksCount)
- [GetForksURL](https://godoc.org/github.com/google/go-github/github#Repository.GetForksURL)
- [GetFullName](https://godoc.org/github.com/google/go-github/github#Repository.GetFullName)
- [GetGitCommitsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetGitCommitsURL)
- [GetGitRefsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetGitRefsURL)
- [GetGitTagsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetGitTagsURL)
- [GetGitURL](https://godoc.org/github.com/google/go-github/github#Repository.GetGitURL)
- [GetGitignoreTemplate](https://godoc.org/github.com/google/go-github/github#Repository.GetGitignoreTemplate)
- [GetHTMLURL](https://godoc.org/github.com/google/go-github/github#Repository.GetHTMLURL)
- [GetHasDownloads](https://godoc.org/github.com/google/go-github/github#Repository.GetHasDownloads)
- [GetHasIssues](https://godoc.org/github.com/google/go-github/github#Repository.GetHasIssues)
- [GetHasPages](https://godoc.org/github.com/google/go-github/github#Repository.GetHasPages)
- [GetHasProjects](https://godoc.org/github.com/google/go-github/github#Repository.GetHasProjects)
- [GetHasWiki](https://godoc.org/github.com/google/go-github/github#Repository.GetHasWiki)
- [GetHomepage](https://godoc.org/github.com/google/go-github/github#Repository.GetHomepage)
- [GetHooksURL](https://godoc.org/github.com/google/go-github/github#Repository.GetHooksURL)
- [GetID](https://godoc.org/github.com/google/go-github/github#Repository.GetID)
- [GetIssueCommentURL](https://godoc.org/github.com/google/go-github/github#Repository.GetIssueCommentURL)
- [GetIssueEventsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetIssueEventsURL)
- [GetIssuesURL](https://godoc.org/github.com/google/go-github/github#Repository.GetIssuesURL)
- [GetKeysURL](https://godoc.org/github.com/google/go-github/github#Repository.GetKeysURL)
- [GetLabelsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetLabelsURL)
- [GetLanguage](https://godoc.org/github.com/google/go-github/github#Repository.GetLanguage)
- [GetLanguagesURL](https://godoc.org/github.com/google/go-github/github#Repository.GetLanguagesURL)
- [GetLicense](https://godoc.org/github.com/google/go-github/github#Repository.GetLicense)
- [GetLicenseTemplate](https://godoc.org/github.com/google/go-github/github#Repository.GetLicenseTemplate)
- [GetMasterBranch](https://godoc.org/github.com/google/go-github/github#Repository.GetMasterBranch)
- [GetMergesURL](https://godoc.org/github.com/google/go-github/github#Repository.GetMergesURL)
- [GetMilestonesURL](https://godoc.org/github.com/google/go-github/github#Repository.GetMilestonesURL)
- [GetMirrorURL](https://godoc.org/github.com/google/go-github/github#Repository.GetMirrorURL)
- [GetName](https://godoc.org/github.com/google/go-github/github#Repository.GetName)
- [GetNetworkCount](https://godoc.org/github.com/google/go-github/github#Repository.GetNetworkCount)
- [GetNodeID](https://godoc.org/github.com/google/go-github/github#Repository.GetNodeID)
- [GetNotificationsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetNotificationsURL)
- [GetOpenIssuesCount](https://godoc.org/github.com/google/go-github/github#Repository.GetOpenIssuesCount)
- [GetOrganization](https://godoc.org/github.com/google/go-github/github#Repository.GetOrganization)
- [GetOwner](https://godoc.org/github.com/google/go-github/github#Repository.GetOwner)
- [GetParent](https://godoc.org/github.com/google/go-github/github#Repository.GetParent)
- [GetPermissions](https://godoc.org/github.com/google/go-github/github#Repository.GetPermissions)
- [GetPrivate](https://godoc.org/github.com/google/go-github/github#Repository.GetPrivate)
- [GetPullsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetPullsURL)
- [GetPushedAt](https://godoc.org/github.com/google/go-github/github#Repository.GetPushedAt)
- [GetReleasesURL](https://godoc.org/github.com/google/go-github/github#Repository.GetReleasesURL)
- [GetSSHURL](https://godoc.org/github.com/google/go-github/github#Repository.GetSSHURL)
- [GetSVNURL](https://godoc.org/github.com/google/go-github/github#Repository.GetSVNURL)
- [GetSize](https://godoc.org/github.com/google/go-github/github#Repository.GetSize)
- [GetSource](https://godoc.org/github.com/google/go-github/github#Repository.GetSource)
- [GetStargazersCount](https://godoc.org/github.com/google/go-github/github#Repository.GetStargazersCount)
- [GetStargazersURL](https://godoc.org/github.com/google/go-github/github#Repository.GetStargazersURL)
- [GetStatusesURL](https://godoc.org/github.com/google/go-github/github#Repository.GetStatusesURL)
- [GetSubscribersCount](https://godoc.org/github.com/google/go-github/github#Repository.GetSubscribersCount)
- [GetSubscribersURL](https://godoc.org/github.com/google/go-github/github#Repository.GetSubscribersURL)
- [GetSubscriptionURL](https://godoc.org/github.com/google/go-github/github#Repository.GetSubscriptionURL)
- [GetTagsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetTagsURL)
- [GetTeamID](https://godoc.org/github.com/google/go-github/github#Repository.GetTeamID)
- [GetTeamsURL](https://godoc.org/github.com/google/go-github/github#Repository.GetTeamsURL)
- [GetTreesURL](https://godoc.org/github.com/google/go-github/github#Repository.GetTreesURL)
- [GetURL](https://godoc.org/github.com/google/go-github/github#Repository.GetURL)
- [GetUpdatedAt](https://godoc.org/github.com/google/go-github/github#Repository.GetUpdatedAt)
- [GetWatchersCount](https://godoc.org/github.com/google/go-github/github#Repository.GetWatchersCount)

#### Example

```
[{{.Repository.GetName}}]({{.Repository.GetHTMLURL}})
```

### [Milestone](https://godoc.org/github.com/google/go-github/github#Milestone)

Milestone represents a GitHub repository milestone.
Only available if **milestone** setting is set.

#### Functions

- [GetClosedAt](https://godoc.org/github.com/google/go-github/github#Milestone.GetClosedAt)
- [GetClosedIssues](https://godoc.org/github.com/google/go-github/github#Milestone.GetClosedIssues)
- [GetCreatedAt](https://godoc.org/github.com/google/go-github/github#Milestone.GetCreatedAt)
- [GetCreator](https://godoc.org/github.com/google/go-github/github#Milestone.GetCreator)
- [GetDescription](https://godoc.org/github.com/google/go-github/github#Milestone.GetDescription)
- [GetDueOn](https://godoc.org/github.com/google/go-github/github#Milestone.GetDueOn)
- [GetHTMLURL](https://godoc.org/github.com/google/go-github/github#Milestone.GetHTMLURL)
- [GetID](https://godoc.org/github.com/google/go-github/github#Milestone.GetID)
- [GetLabelsURL](https://godoc.org/github.com/google/go-github/github#Milestone.GetLabelsURL)
- [GetNodeID](https://godoc.org/github.com/google/go-github/github#Milestone.GetNodeID)
- [GetNumber](https://godoc.org/github.com/google/go-github/github#Milestone.GetNumber)
- [GetOpenIssues](https://godoc.org/github.com/google/go-github/github#Milestone.GetOpenIssues)
- [GetState](https://godoc.org/github.com/google/go-github/github#Milestone.GetState)
- [GetTitle](https://godoc.org/github.com/google/go-github/github#Milestone.GetTitle)
- [GetURL](https://godoc.org/github.com/google/go-github/github#Milestone.GetURL)
- [GetUpdatedAt](https://godoc.org/github.com/google/go-github/github#Milestone.GetUpdatedAt)

#### Example

```
{{.Milestone.GetTitle}} ({{.Milestone.GetClosedAt.Format "2006-01-02"}})
```

### [Commit](https://godoc.org/github.com/google/go-github/github#Commit) (SinceTagCommit & UntilTagCommit )

Commit represents a GitHub commit.
Only available if **since-tag** and/or **until-tag** setting is set.

#### Functions

- [GetAuthor](https://godoc.org/github.com/google/go-github/github#Commit.GetAuthor)
- [GetCommentCount](https://godoc.org/github.com/google/go-github/github#Commit.GetCommentCount)
- [GetCommitter](https://godoc.org/github.com/google/go-github/github#Commit.GetCommitter)
- [GetHTMLURL](https://godoc.org/github.com/google/go-github/github#Commit.GetHTMLURL)
- [GetMessage](https://godoc.org/github.com/google/go-github/github#Commit.GetMessage)
- [GetNodeID](https://godoc.org/github.com/google/go-github/github#Commit.GetNodeID)
- [GetSHA](https://godoc.org/github.com/google/go-github/github#Commit.GetSHA)
- [GetStats](https://godoc.org/github.com/google/go-github/github#Commit.GetStats)
- [GetTree](https://godoc.org/github.com/google/go-github/github#Commit.GetTree)
- [GetURL](https://godoc.org/github.com/google/go-github/github#Commit.GetURL)
- [GetVerification](https://godoc.org/github.com/google/go-github/github#Commit.GetVerification)

#### Example

```
[{{.SinceTag}}]({{.SinceTagCommit.GetHTMLURL}}) - [{{.UntilTag}}]({{.UntilTagCommit.GetHTMLURL}})
```

### [Issue](https://godoc.org/github.com/google/go-github/github#Issue)

Issue represents a GitHub issue on a repository.

Note: As far as the GitHub API is concerned, every pull request is an issue, but not every issue is a pull request. Some endpoints, events, and webhooks may also return pull requests via this struct. If PullRequestLinks is nil, this is an issue, and if PullRequestLinks is not nil, this is a pull request. The IsPullRequest helper method can be used to check that.

#### Functions

- [GetActiveLockReason](https://godoc.org/github.com/google/go-github/github#Issue.GetActiveLockReason)
- [GetAssignee](https://godoc.org/github.com/google/go-github/github#Issue.GetAssignee)
- [GetBody](https://godoc.org/github.com/google/go-github/github#Issue.GetBody)
- [GetClosedAt](https://godoc.org/github.com/google/go-github/github#Issue.GetClosedAt)
- [GetClosedBy](https://godoc.org/github.com/google/go-github/github#Issue.GetClosedBy)
- [GetComments](https://godoc.org/github.com/google/go-github/github#Issue.GetComments)
- [GetCommentsURL](https://godoc.org/github.com/google/go-github/github#Issue.GetCommentsURL)
- [GetCreatedAt](https://godoc.org/github.com/google/go-github/github#Issue.GetCreatedAt)
- [GetEventsURL](https://godoc.org/github.com/google/go-github/github#Issue.GetEventsURL)
- [GetHTMLURL](https://godoc.org/github.com/google/go-github/github#Issue.GetHTMLURL)
- [GetID](https://godoc.org/github.com/google/go-github/github#Issue.GetID)
- [GetLabelsURL](https://godoc.org/github.com/google/go-github/github#Issue.GetLabelsURL)
- [GetLocked](https://godoc.org/github.com/google/go-github/github#Issue.GetLocked)
- [GetMilestone](https://godoc.org/github.com/google/go-github/github#Issue.GetMilestone)
- [GetNodeID](https://godoc.org/github.com/google/go-github/github#Issue.GetNodeID)
- [GetNumber](https://godoc.org/github.com/google/go-github/github#Issue.GetNumber)
- [GetPullRequestLinks](https://godoc.org/github.com/google/go-github/github#Issue.GetPullRequestLinks)
- [GetReactions](https://godoc.org/github.com/google/go-github/github#Issue.GetReactions)
- [GetRepository](https://godoc.org/github.com/google/go-github/github#Issue.GetRepository)
- [GetRepositoryURL](https://godoc.org/github.com/google/go-github/github#Issue.GetRepositoryURL)
- [GetState](https://godoc.org/github.com/google/go-github/github#Issue.GetState)
- [GetTitle](https://godoc.org/github.com/google/go-github/github#Issue.GetTitle)
- [GetURL](https://godoc.org/github.com/google/go-github/github#Issue.GetURL)
- [GetUpdatedAt](https://godoc.org/github.com/google/go-github/github#Issue.GetUpdatedAt)
- [GetUser](https://godoc.org/github.com/google/go-github/github#Issue.GetUser)
- [IsPullRequest](https://godoc.org/github.com/google/go-github/github#Issue.IsPullRequest)

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
