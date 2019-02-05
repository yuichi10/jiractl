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
func GetSprintIssues(input ISprintIssuesInput, api IJiraAPIAccess, db ICurrentContextDataAccess) {
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
	fmt.Printf("%+v\n", board)
	fmt.Printf("%+v\n", s)
}
