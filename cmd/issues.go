package cmd

import (
	"github.com/google/go-github/github"
)

type TitledIssue struct {
	Title  string
	Issues []*github.Issue
}

type TitledIssues []*TitledIssue

func (t *TitledIssues) AllIssues() []*github.Issue {
	var allIssues []*github.Issue
	for _, ti := range *t {
		allIssues = append(allIssues, ti.Issues...)
	}
	return allIssues
}
