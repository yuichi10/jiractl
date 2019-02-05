package usecase

import "fmt"

type ISprintIssuesInput interface {
	GetBoardName() string
}

// GetSprintIssues get apis
func GetSprintIssues(input ISprintIssuesInput, api IJiraAPIAccess, db ICurrentContextDataAccess) {
	c, err := db.GetCurrentContext()
	if err != nil {
		panic(err)
	}
	board, err := api.GetBoardInfo(c.JiraURL, input.GetBoardName(), c.BasicAuth)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", board)
}
