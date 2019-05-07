package internal

import (
	"github.com/google/go-github/github"
)

type GroupedIssues struct {
	Title  string
	Issues []*github.Issue
}

func GroupIssues(groups []*Group, issues []*github.Issue) []*GroupedIssues {
	if issues == nil {
		return nil
	}
	var result []*GroupedIssues
	grouped := make(map[string][]*github.Issue)

	for _, issue := range issues {
		if i, ok := containsAny(issue.Labels, AllLabels(groups)); ok {
			grouped[groups[i].Title] = append(grouped[groups[i].Title], issue)
		} else {
			grouped["no_label"] = append(grouped["no_label"], issue)
		}
	}

	for _, group := range groups {
		if len(group.Labels) == 0 {
			result = append(result, &GroupedIssues{Title: group.Title, Issues: grouped["no_label"]})
			continue
		}
		result = append(result, &GroupedIssues{Title: group.Title, Issues: grouped[group.Title]})
	}

	return result
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
