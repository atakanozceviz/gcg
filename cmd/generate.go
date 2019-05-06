package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate changelog",
	Long: `Generates changelog based on a milestone.
Edit the config.yaml file to customize the generated output.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := &Config{}
		err := viper.Unmarshal(c)
		if err != nil {
			er(err)
		}
		ti, err := issuesByMilestone(c)
		if err != nil {
			er(err)
		}

		mil, err := getMilestone(c)
		if err != nil {
			er(err)
		}

		result, err := createChangelog(&TplData{ti, mil})
		if err != nil {
			er(err)
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("milestone", "m", 0, "milestone number")
	generateCmd.Flags().StringP("owner", "o", "", "owner of the repo")
	generateCmd.Flags().StringP("repo", "r", "", "name of the repo")
	generateCmd.Flags().StringP("state", "s", "", "state of the issues to fetch")
	generateCmd.Flags().StringP("token", "t", "", "personal access token")

	err := viper.BindPFlags(generateCmd.Flags())
	if err != nil {
		er(err)
	}
}

type TplData struct {
	TitledIssues *TitledIssues
	Milestone    *github.Milestone
}

func createChangelog(td *TplData) (string, error) {
	template := viper.GetString("template")
	if template == "" {
		template = `## [{{.Milestone.GetTitle}}]({{.Milestone.GetHTMLURL}})
{{range .TitledIssues}}
### {{.Title}}
{{range .Issues}}
{{- if .IsPullRequest }}
- PR [\#{{.GetNumber}}]({{.GetHTMLURL}}): {{.GetTitle}} (by [{{.GetUser.GetLogin}}]({{.GetUser.GetHTMLURL}}))
{{- else}}
- ISSUE [\#{{.GetNumber}}]({{.GetHTMLURL}}): {{.GetTitle}}
{{- end}}
{{- end}}
{{end}}`
	}

	return executeTemplate(template, td)
}

func getMilestone(config *Config) (*github.Milestone, error) {
	client := github.NewClient(nil)

	if config.Token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: config.Token},
		)
		tc := oauth2.NewClient(context.Background(), ts)
		client = github.NewClient(tc)
	}

	milestone, _, err := client.Issues.GetMilestone(context.Background(), config.Owner, config.Repo, config.Milestone)
	if err != nil {
		return nil, err
	}
	return milestone, nil
}
