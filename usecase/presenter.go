package usecase

import "os"

type IssuesOutput struct {
	IssueType string
	Summary   string
	Assignee  string
	Status    string
	URL       string
}

type IIssuePresenter interface {
	Present(out []*IssuesOutput)
	Format() string
	File() *os.File
}
