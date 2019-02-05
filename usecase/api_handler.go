package usecase

import "github.com/yuichi10/jiractl/entity"

type IJiraAPIAccess interface {
	GetBoardInfo(baseURL, name, basicAuth string) (*entity.JiraBoard, error)
	GetSprintInfo(baseURL, basicAuth, sprintName string, boardID int) (*entity.JiraSprint, error)
}
