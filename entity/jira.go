package entity

// Credential have credential information of user
type Credential struct {
	Name      string
	UserID    string
	BasicAuth string
	Password  string
}

type JiraURL struct {
	Name string
	URL  string
}

type Context struct {
	Name      string
	UserID    string
	BasicAuth string
	URL       string
}

type Current struct {
	ContextName string
	UserID      string
	URL         string
	BasicAuth   string
}

// Me has information of me user
type Me struct {
	URL          string
	Name         string
	EmailAddress string
	AccountID    string
}
