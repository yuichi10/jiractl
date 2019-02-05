package entity

type JiraIssue struct {
	ID           string
	IssueType    string
	Summary      string
	Description  string
	AssigneeName string
	Status       string
	URL          string
}

type JiraSprint struct {
	SprintID   int
	SprintName string
	State      string
}

type JiraBoard struct {
	ID   int
	Name string
}
