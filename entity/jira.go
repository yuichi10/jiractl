package entity

// Credential have credential information of user
type Credential struct {
	UserID    string
	BasicAuth string
	Password  string
}

// Me has information of me user
type Me struct {
	URL          string
	Name         string
	EmailAddress string
	AccountID    string
}
