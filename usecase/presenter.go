package usecase

type Line struct {
	Body      string
	Color     string
	Delimiter string
}
type Lines []*Line

type Presenter interface {
	Present(Lines)
}

type IssueOutput struct {
	IssueType string
	Summary   string
	Assignee  string
	Status    string
	URL       string
}

type IssuesOutput []*IssueOutput

type IIssuePresenter interface {
	IssuePresent(out IssuesOutput, format string, detail bool)
}
