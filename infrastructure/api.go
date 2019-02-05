package infrastructure

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type jiraAPIClient struct {
	client *http.Client
}

// NewJiraAPIClient return jira api accesser
func NewJiraAPIClient() *jiraAPIClient {
	return &jiraAPIClient{client: &http.Client{}}
}

func (c jiraAPIClient) Get(url, body string, params url.Values, header http.Header) ([]byte, error) {
	req, err := http.NewRequest("GET", url, bytes.NewBufferString(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create reqest of %q: %v", url, err)
	}
	req.Header = header
	req.URL.RawQuery = params.Encode()
	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get from %q api: %v", url, err)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	return b, nil
}
