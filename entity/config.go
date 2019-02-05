package entity

// Credential have credential information of user
type Credential struct {
	Name      string
	UserID    string
	BasicAuth string
	Password  string
}

// JiraURL is jira url info
type JiraURL struct {
	Name string
	URL  string
}

// Context is jira context info
type Context struct {
	Name      string
	UserID    string
	BasicAuth string
	URL       string
}

// Current is current context data
type Current struct {
	ContextName string
	UserID      string
	URL         string
	BasicAuth   string
	JiraURL     string
}

// Me has information of me user
type Me struct {
	URL          string
	Name         string
	EmailAddress string
	AccountID    string
}
