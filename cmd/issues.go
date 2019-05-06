package cmd

import (
	"context"
	"strconv"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type TitledIssue struct {
	Title  string
	Issues []*github.Issue
}

type TitledIssues []*TitledIssue

func issuesByMilestone(config *Config) (*TitledIssues, error) {
	opt := &github.IssueListByRepoOptions{
		Milestone: strconv.Itoa(config.Milestone),
		State:     config.State,
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}
	ctx := context.Background()
	issues, err := allIssues(config, ctx, opt)
	if err != nil {
		return nil, err
	}

	return groupIssues(config, issues)
}

func allIssues(config *Config, ctx context.Context, opt *github.IssueListByRepoOptions) ([]*github.Issue, error) {
	client := github.NewClient(nil)

	if config.Token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: config.Token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}

	var allIssues []*github.Issue
	for {
		issues, resp, err := client.Issues.ListByRepo(ctx, config.Owner, config.Repo, opt)
		if err != nil {
			return nil, err
		}
		allIssues = append(allIssues, issues...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allIssues, nil
}

func groupIssues(config *Config, issues []*github.Issue) (*TitledIssues, error) {
	result := &TitledIssues{}
	grouped := make(map[string][]*github.Issue)

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
