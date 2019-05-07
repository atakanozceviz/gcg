package cmd

import (
	"fmt"

	"github.com/atakanozceviz/gcg/internal"
	"github.com/atakanozceviz/gcg/pkg"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate changelog",
	Long: `Generates changelog based on a milestone.
Edit the config.yaml file to customize the generated output.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := &internal.Config{}
		err := viper.Unmarshal(c)
		if err != nil {
			er(err)
		}

		repo, err := pkg.NewRepo(c.Repo, c.Token)
		if err != nil {
			er(err)
		}

		issuesByMilestone, err := repo.IssuesByMilestone(c.Milestone, c.State)
		if err != nil {
			er(err)
		}
		milestone, err := repo.Milestone(c.Milestone)
		if err != nil {
			er(err)
		}
		sinceTagCommit, err := repo.TagCommit(c.SinceTag)
		if err != nil {
			er(err)
		}
		untilTagCommit, err := repo.TagCommit(c.UntilTag)
		if err != nil {
			er(err)
		}

		var issuesByTag []*github.Issue
		switch {
		case sinceTagCommit != nil:
			issuesByTag, err = repo.IssuesSince(sinceTagCommit.Committer.GetDate())
			if err != nil {
				er(err)
			}
			if c.UntilTag != "" {
				issuesByTag = pkg.FilterUntil(issuesByTag, untilTagCommit.Committer.GetDate())
			}
		case untilTagCommit != nil:
			issuesByTag, err = repo.AllIssues(c.State)
			er(err)
			issuesByTag = pkg.FilterUntil(issuesByTag, untilTagCommit.Committer.GetDate())
		}

		groupedIssuesByMilestone := internal.GroupIssues(c.Groups, issuesByMilestone)
		if err != nil {
			er(err)
		}

		groupedIssuesByTag := internal.GroupIssues(c.Groups, issuesByTag)
		if err != nil {
			er(err)
		}

		changelog, err := createChangelog(&TplData{
			IssuesByMilestone: groupedIssuesByMilestone,
			IssuesByTag:       groupedIssuesByTag,
			Milestone:         milestone,
			SinceTag:          c.SinceTag,
			SinceTagCommit:    sinceTagCommit,
			UntilTag:          c.UntilTag,
			UntilTagCommit:    untilTagCommit,
		})
		if err != nil {
			er(err)
		}

		fmt.Printf("%s", changelog)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("milestone", "m", "", "Github milestone title to get issues and pull requests for")
	generateCmd.Flags().StringP("repo", "r", "", "Repository name to generate the Changelog for, in the form user/repo")
	generateCmd.Flags().String("since-tag", "", "Github issues and pull requests since tag")
	generateCmd.Flags().StringP("state", "s", "", "State of the issues to get (open,closed or all)")
	generateCmd.Flags().StringP("token", "t", "", "Github access token")
	generateCmd.Flags().String("until-tag", "", "Github issues and pull requests until tag")

	err := viper.BindPFlags(generateCmd.Flags())
	if err != nil {
		er(err)
	}
}

type TplData struct {
	IssuesByMilestone []*internal.GroupedIssues
	IssuesByTag       []*internal.GroupedIssues
	Milestone         *github.Milestone
	SinceTag          string
	SinceTagCommit    *github.Commit
	UntilTag          string
	UntilTagCommit    *github.Commit
}

func createChangelog(td *TplData) ([]byte, error) {
	template := viper.GetString("template")
	if template == "" {
		template = `{{if .Milestone}}## {{.Milestone.GetTitle}} ({{.Milestone.GetClosedAt.Format "2006-01-02"}}){{end -}}
{{if .IssuesByMilestone}}
{{range .IssuesByMilestone}}
### {{.Title}}
{{range .Issues}}
{{if .IsPullRequest -}}
- PR [\#{{.GetNumber}}]({{.GetHTMLURL}}): {{.GetTitle}} (by [{{.GetUser.GetLogin}}]({{.GetUser.GetHTMLURL}}))
{{- else -}}
- ISSUE [\#{{.GetNumber}}]({{.GetHTMLURL}}): {{.GetTitle}}
{{- end -}}
{{end}}
{{end}}
{{end -}}
{{if .SinceTagCommit}}## {{.SinceTag}}{{if .UntilTagCommit}} - {{.UntilTag}}{{end}}{{end -}}
{{if .IssuesByTag}}
{{range .IssuesByTag}}
### {{.Title}}
{{range .Issues}}
{{if .IsPullRequest -}}
- PR [\#{{.GetNumber}}]({{.GetHTMLURL}}): {{.GetTitle}} (by [{{.GetUser.GetLogin}}]({{.GetUser.GetHTMLURL}}))
{{- else -}}
- ISSUE [\#{{.GetNumber}}]({{.GetHTMLURL}}): {{.GetTitle}}
{{- end -}}
{{end}}
{{end}}
{{- end -}}
`
	}

	return executeTemplate(template, td)
}
