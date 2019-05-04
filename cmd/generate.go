package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate changelog",
	Run: func(cmd *cobra.Command, args []string) {
		c := &Config{}
		err := viper.Unmarshal(c)
		if err != nil {
			er(err)
		}
		ti, err := getIssues(c)
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
{{.Title}}
{{range .Issues}}
{{- if isPR .GetHTMLURL}}
- PR [\#{{.GetNumber}}]({{.GetHTMLURL}}): {{.GetTitle}} (by [{{.GetUser.GetLogin}}]({{.GetUser.GetHTMLURL}}))
{{- else}}
- ISSUE [\#{{.GetNumber}}]({{.GetHTMLURL}}): {{.GetTitle}}
{{- end}}
{{- end}}
{{end}}`
	}

	return executeTemplate(template, td)
}

func getIssues(config *Config) (*TitledIssues, error) {
	ctx := context.Background()
	client := github.NewClient(nil)
	result := &TitledIssues{}
	grouped := make(map[string][]*github.Issue)

	opt := &github.IssueListByRepoOptions{
		Milestone: strconv.Itoa(config.Milestone),
		State:     config.State,
		ListOptions: github.ListOptions{
			PerPage: 2147483647,
		},
	}

	issues, _, err := client.Issues.ListByRepo(ctx, config.Owner, config.Repo, opt)
	if err != nil {
		return nil, err
	}

	for _, issue := range issues {
		if i, ok := containsAny(issue.Labels, config.AllLabels()); ok {
			grouped[config.Groups[i].Title] = append(grouped[config.Groups[i].Title], issue)
		} else {
			grouped["no_label"] = append(grouped["no_label"], issue)
		}
	}

	for _, group := range config.Groups {
		if len(group.Labels) == 0 {
			*result = append(*result, &TitledIssue{Title: group.Title, Issues: grouped["no_label"]})
			continue
		}
		*result = append(*result, &TitledIssue{Title: group.Title, Issues: grouped[group.Title]})
	}

	return result, nil
}

func getMilestone(config *Config) (*github.Milestone, error) {
	client := github.NewClient(nil)
	milestone, _, err := client.Issues.GetMilestone(context.Background(), config.Owner, config.Repo, config.Milestone)
	if err != nil {
		return nil, err
	}
	return milestone, nil
}

func containsAny(s []github.Label, e map[int][]string) (int, bool) {
	for _, a := range s {
		for i, b := range e {
			for _, y := range b {
				if a.GetName() == y {
					return i, true
				}
			}
		}
	}
	return 0, false
}
