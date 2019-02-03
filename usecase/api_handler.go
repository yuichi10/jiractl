package usecase

import "github.com/yuichi10/jiractl/entity"

type IJiraAPIAccess interface {
	GetActiveSprintIssues() []*entity.Issue
}
