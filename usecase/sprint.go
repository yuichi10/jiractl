package usecase

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

type ISprintIssuesInput interface {
	GetBoardName() string
	GetSprintName() string
}

// GetSprintIssues get apis
func GetSprintIssues(input ISprintIssuesInput, api IJiraAPIAccess, db ICurrentContextDataAccess, presenter IIssuePresenter) {
	c, err := db.GetCurrentContext()
	if err != nil {
		zap.S().Errorf("failed to get current context %v", err)
		os.Exit(1)
	}
	board, err := api.GetBoardInfo(c.JiraURL, input.GetBoardName(), c.BasicAuth)
	if err != nil {
		zap.S().Errorf("failed to get board info: %v", err)
	}
	s, err := api.GetSprintInfo(c.JiraURL, c.BasicAuth, input.GetSprintName(), board.ID)
	if err != nil {
		zap.S().Errorf("failed to get sprint info: %v", err)
		os.Exit(1)
	}
	issues, err := api.GetSprintIssuesInfo(c.JiraURL, c.BasicAuth, s.SprintID)
	if err != nil {
		zap.S().Errorf("failed to get sprint isses info: %v", err)
		os.Exit(1)
	}
	output := make([]*IssuesOutput, 0, 15)
	for _, issue := range issues {
		o := &IssuesOutput{
			IssueType: issue.IssueType,
			Summary:   issue.Summary,
			Assignee:  issue.AssigneeName,
			Status:    issue.Status,
			URL:       issue.URL,
		}
		output = append(output, o)
	}
	presenter.Present(output)
	fmt.Printf("%+v\n", board)
	fmt.Printf("%+v\n", s)
	for _, i := range issues {
		fmt.Println(i.Summary)
	}
}
