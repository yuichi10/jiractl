package presenter

import (
	"fmt"
	"sync"

	"sort"

	"github.com/yuichi10/jiractl/usecase"
)

type SprintPresenter struct {
	View   Viewer
	Format string
}

func NewSprintPresenter(v Viewer, format string) SprintPresenter {
	return SprintPresenter{View: v, Format: format}
}

func (ip SprintPresenter) IssuePresent(out usecase.IssuesOutput, format string) {
	humble := Lines{}
	if format == "markdown" {
		humble = markdownIssuePresenter(out)
	}
	ip.View.Show(humble)
}

func markdownIssuePresenter(out usecase.IssuesOutput) Lines {
	humble := make(Lines, 0, 50)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Status < out[j].Status })
	sort.SliceStable(out, func(i, j int) bool {
		return issueTypeSortNumber(out[i].IssueType) < issueTypeSortNumber(out[j].IssueType)
	})
	var storyOnce sync.Once
	var subTaskOnce sync.Once
	var taskOnce sync.Once
	var issueOnce sync.Once
	for _, issue := range out {
		l := &Line{}
		l.Body = fmt.Sprintf("[%s](%s) %s %s", issue.Summary, issue.URL, issue.Status, issue.Assignee)
		l.Delimiter = "\n"
		if issue.IssueType == "ストーリー" || issue.IssueType == "Story" {
			storyOnce.Do(func() { humble = append(humble, &Line{Body: "Story", Color: "97", Delimiter: "\n"}) })
			l.Color = "36"
		} else if issue.IssueType == "サブタスク" || issue.IssueType == "SubTask" {
			subTaskOnce.Do(func() { humble = append(humble, &Line{Body: "SubTask", Color: "97", Delimiter: "\n"}) })
			l.Color = "34"
		} else if issue.IssueType == "タスク" || issue.IssueType == "Task" {
			taskOnce.Do(func() { humble = append(humble, &Line{Body: "Task", Color: "97", Delimiter: "\n"}) })
			l.Color = "33"
		} else if issue.IssueType == "改善" || issue.IssueType == "Issue" {
			issueOnce.Do(func() { humble = append(humble, &Line{Body: "Issue", Color: "97", Delimiter: "\n"}) })
			l.Color = "95"
		}
		humble = append(humble, l)
	}
	return humble
}

func issueTypeSortNumber(it string) int {
	switch it {
	case "ストーリー", "Story":
		return 1
	case "サブタスク", "SubTask":
		return 2
	case "タスク", "Task":
		return 3
	case "改善", "Issue":
		return 4
	}
	return 99
}
