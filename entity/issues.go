package entity

type Issue struct {
	URL string
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
