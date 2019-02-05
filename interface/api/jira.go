package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/yuichi10/jiractl/entity"
)

type JiraAPIClient struct {
	api IAPI
}

// NewJiraAPI return jira api interface
func NewJiraAPI(api IAPI) *JiraAPIClient {
	return &JiraAPIClient{api: api}
}

func (c JiraAPIClient) GetCurrentSprintIssues() {
	// TODO: Get Board info
	// /rest/agile/1.0/board?name=DPS-CD1 7661
	// TODO: Get all sprint info
	// /rest/agile/1.0/board/7661/sprint
	// TODO: get active sprint id (state: active) 16245
	// TODO: get sprint issues
	// /rest/agile/1.0/sprint/16245/issue
}

func (c JiraAPIClient) GetBoardInfo(baseURL, name, basicAuth string) (*entity.JiraBoard, error) {
	uri, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to base baseURL %q: %v", uri, err)
	}
	uri.Path = path.Join(uri.Path, "/rest/agile/1.0/board")
	param := url.Values{}
	param.Add("name", name)
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Basic %s", basicAuth))
	header.Add("Content-Type", "application/json")
	res, err := c.api.Get(uri.String(), "", param, header)
	if err != nil {
		return nil, err
	}
	board := &JiraBoards{}
	err = json.Unmarshal(res, board)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal get board response: %v", err)
	}
	for _, b := range board.Values {
		if b.Name == name {
			return &entity.JiraBoard{Name: b.Name, ID: b.ID}, nil
		}
	}
	return nil, fmt.Errorf("failed to find board")
}
