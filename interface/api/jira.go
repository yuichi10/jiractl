package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/yuichi10/jiractl/entity"
	"go.uber.org/zap"
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
	zap.S().Infof("get board info url: %s", uri.String())
	param := url.Values{}
	param.Add("name", name)
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Basic %s", basicAuth))
	header.Add("Content-Type", "application/json")
	res, statusCode, err := c.api.Get(uri.String(), "", param, header)
	if err != nil {
		return nil, fmt.Errorf("failed to get board info: %v", err)
	}
	board := &JiraBoards{}
	err = json.Unmarshal(res, board)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal board get. status code %v, err: %v", statusCode, err)
	}
	for _, b := range board.Values {
		if b.Name == name {
			return &entity.JiraBoard{Name: b.Name, ID: b.ID}, nil
		}
	}
	return nil, fmt.Errorf("failed to find board")
}

// GetSprintInfo get sprint info from jira sprint api
func (c JiraAPIClient) GetSprintInfo(baseURL, basicAuth, sprintName string, boardID int) (*entity.JiraSprint, error) {
	uri, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to base baseURL %q: %v", uri, err)
	}
	uri.Path = path.Join(uri.Path, "/rest/agile/1.0/board/", strconv.Itoa(boardID), "sprint")
	zap.S().Infof("get sprint info url: %s", uri.String())
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Basic %s", basicAuth))
	header.Add("Content-Type", "application/json")
	res, statusCode, err := c.api.Get(uri.String(), "", url.Values{}, header)
	if err != nil {
		return nil, fmt.Errorf("failed to get sprint info: %v", err)
	}
	sprints := &JiraSprints{}
	err = json.Unmarshal(res, sprints)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal. status code %v, err: %v", statusCode, err)
	}
	for _, s := range sprints.Values {
		if sprintName == "" && s.State == "active" {
			js := &entity.JiraSprint{SprintID: s.ID, SprintName: s.Name, State: s.State}
			return js, nil
		}
		if sprintName == s.Name {
			js := &entity.JiraSprint{SprintID: s.ID, SprintName: s.Name, State: s.State}
			return js, nil
		}
	}
	return nil, fmt.Errorf("there is no sprint")
}
