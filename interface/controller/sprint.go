package controller

import (
	"github.com/yuichi10/jiractl/interface/api"
	"github.com/yuichi10/jiractl/interface/database"
	"github.com/yuichi10/jiractl/interface/presenter"
	"github.com/yuichi10/jiractl/usecase"
)

type SprintIssueInput struct {
	Board  string
	Sprint string
}

func (input SprintIssueInput) GetBoardName() string {
	return input.Board
}

func (input SprintIssueInput) GetSprintName() string {
	return input.Sprint
}

func GetSprintIssue(board, sprint, format string, iapi api.IAPI, ds database.IDataStore, viewer presenter.Viewer) {
	input := SprintIssueInput{Board: board, Sprint: sprint}
	a := api.NewJiraAPI(iapi)
	db := database.NewConfig(ds)
	p := presenter.NewSprintPresenter(viewer, format)
	// call usecase
	usecase.GetSprintIssues(input, a, db, p)
}
