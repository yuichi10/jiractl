package usecase

import "github.com/yuichi10/jiractl/entity"

type IJiraAPIAccess interface {
	GetBoardInfo(baseURL, name, basicAuth string) (*entity.JiraBoard, error)
}
