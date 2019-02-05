package api

import "time"

// JiraBoards is agile board list
type JiraBoards struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID    int    `json:"id"`
		Self  string `json:"self"`
		Name  string `json:"name"`
		State string `json:"state"`
		Type  string `json:"type"`
	} `json:"values"`
}

// JiraSprints is sprint list of jira
type JiraSprints struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID            int       `json:"id"`
		Self          string    `json:"self"`
		State         string    `json:"state"`
		Name          string    `json:"name"`
		StartDate     time.Time `json:"startDate,omitempty"`
		EndDate       time.Time `json:"endDate,omitempty"`
		CompleteDate  time.Time `json:"completeDate,omitempty"`
		OriginBoardID int       `json:"originBoardId"`
	} `json:"values"`
}

// Issues is information of jira  issue
type Issues struct {
	Expand     string `json:"expand"`
	StartAt    int    `json:"startAt"`
	MaxResults int    `json:"maxResults"`
	Total      int    `json:"total"`
	Issues     []struct {
		Expand string `json:"expand"`
		ID     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
		Fields struct {
			Issuetype struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Description string `json:"description"`
				IconURL     string `json:"iconUrl"`
				Name        string `json:"name"`
				Subtask     bool   `json:"subtask"`
				AvatarID    int    `json:"avatarId"`
			} `json:"issuetype"`
			Timespent interface{} `json:"timespent"`
			Sprint    struct {
				ID            int       `json:"id"`
				Self          string    `json:"self"`
				State         string    `json:"state"`
				Name          string    `json:"name"`
				StartDate     time.Time `json:"startDate"`
				EndDate       time.Time `json:"endDate"`
				OriginBoardID int       `json:"originBoardId"`
			} `json:"sprint"`
			Project struct {
				Self       string `json:"self"`
				ID         string `json:"id"`
				Key        string `json:"key"`
				Name       string `json:"name"`
				AvatarUrls struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
			} `json:"project"`
			FixVersions        []interface{} `json:"fixVersions"`
			Aggregatetimespent interface{}   `json:"aggregatetimespent"`
			Resolution         struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"resolution"`
			Resolutiondate string `json:"resolutiondate"`
			Workratio      int    `json:"workratio"`
			LastViewed     string `json:"lastViewed"`
			Watches        struct {
				Self       string `json:"self"`
				WatchCount int    `json:"watchCount"`
				IsWatching bool   `json:"isWatching"`
			} `json:"watches"`
			Created string `json:"created"`
			Epic    struct {
				ID      int    `json:"id"`
				Key     string `json:"key"`
				Self    string `json:"self"`
				Name    string `json:"name"`
				Summary string `json:"summary"`
				Color   struct {
					Key string `json:"key"`
				} `json:"color"`
				Done bool `json:"done"`
			} `json:"epic"`
			Priority struct {
				Self    string `json:"self"`
				IconURL string `json:"iconUrl"`
				Name    string `json:"name"`
				ID      string `json:"id"`
			} `json:"priority"`
			Labels                        []string      `json:"labels"`
			Timeestimate                  interface{}   `json:"timeestimate"`
			Aggregatetimeoriginalestimate interface{}   `json:"aggregatetimeoriginalestimate"`
			Versions                      []interface{} `json:"versions"`
			Issuelinks                    []interface{} `json:"issuelinks"`
			Assignee                      struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"assignee"`
			Updated string `json:"updated"`
			Status  struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
					Name      string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
			Components []struct {
				Self string `json:"self"`
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"components"`
			Timeoriginalestimate interface{} `json:"timeoriginalestimate"`
			Description          string      `json:"description"`
			Timetracking         struct {
			} `json:"timetracking"`
			Attachment            []interface{} `json:"attachment"`
			Aggregatetimeestimate interface{}   `json:"aggregatetimeestimate"`
			Flagged               bool          `json:"flagged"`
			Summary               string        `json:"summary"`
			Creator               struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"creator"`
			Subtasks []struct {
				ID     string `json:"id"`
				Key    string `json:"key"`
				Self   string `json:"self"`
				Fields struct {
					Summary string `json:"summary"`
					Status  struct {
						Self           string `json:"self"`
						Description    string `json:"description"`
						IconURL        string `json:"iconUrl"`
						Name           string `json:"name"`
						ID             string `json:"id"`
						StatusCategory struct {
							Self      string `json:"self"`
							ID        int    `json:"id"`
							Key       string `json:"key"`
							ColorName string `json:"colorName"`
							Name      string `json:"name"`
						} `json:"statusCategory"`
					} `json:"status"`
					Priority struct {
						Self    string `json:"self"`
						IconURL string `json:"iconUrl"`
						Name    string `json:"name"`
						ID      string `json:"id"`
					} `json:"priority"`
					Issuetype struct {
						Self        string `json:"self"`
						ID          string `json:"id"`
						Description string `json:"description"`
						IconURL     string `json:"iconUrl"`
						Name        string `json:"name"`
						Subtask     bool   `json:"subtask"`
						AvatarID    int    `json:"avatarId"`
					} `json:"issuetype"`
				} `json:"fields"`
			} `json:"subtasks"`
			Reporter struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"reporter"`
			Aggregateprogress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"aggregateprogress"`
			Environment   interface{} `json:"environment"`
			Duedate       interface{} `json:"duedate"`
			ClosedSprints []struct {
				ID            int       `json:"id"`
				Self          string    `json:"self"`
				State         string    `json:"state"`
				Name          string    `json:"name"`
				StartDate     time.Time `json:"startDate"`
				EndDate       time.Time `json:"endDate"`
				CompleteDate  time.Time `json:"completeDate"`
				OriginBoardID int       `json:"originBoardId"`
			} `json:"closedSprints"`
			Progress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"progress"`
			Comment struct {
				Comments []struct {
					Self   string `json:"self"`
					ID     string `json:"id"`
					Author struct {
						Self         string `json:"self"`
						Name         string `json:"name"`
						Key          string `json:"key"`
						EmailAddress string `json:"emailAddress"`
						AvatarUrls   struct {
							Four8X48  string `json:"48x48"`
							Two4X24   string `json:"24x24"`
							One6X16   string `json:"16x16"`
							Three2X32 string `json:"32x32"`
						} `json:"avatarUrls"`
						DisplayName string `json:"displayName"`
						Active      bool   `json:"active"`
						TimeZone    string `json:"timeZone"`
					} `json:"author"`
					Body         string `json:"body"`
					UpdateAuthor struct {
						Self         string `json:"self"`
						Name         string `json:"name"`
						Key          string `json:"key"`
						EmailAddress string `json:"emailAddress"`
						AvatarUrls   struct {
							Four8X48  string `json:"48x48"`
							Two4X24   string `json:"24x24"`
							One6X16   string `json:"16x16"`
							Three2X32 string `json:"32x32"`
						} `json:"avatarUrls"`
						DisplayName string `json:"displayName"`
						Active      bool   `json:"active"`
						TimeZone    string `json:"timeZone"`
					} `json:"updateAuthor"`
					Created string `json:"created"`
					Updated string `json:"updated"`
				} `json:"comments"`
				MaxResults int `json:"maxResults"`
				Total      int `json:"total"`
				StartAt    int `json:"startAt"`
			} `json:"comment"`
			Votes struct {
				Self     string `json:"self"`
				Votes    int    `json:"votes"`
				HasVoted bool   `json:"hasVoted"`
			} `json:"votes"`
			Worklog struct {
				StartAt    int           `json:"startAt"`
				MaxResults int           `json:"maxResults"`
				Total      int           `json:"total"`
				Worklogs   []interface{} `json:"worklogs"`
			} `json:"worklog"`
		} `json:"fields"`
	} `json:"issues"`
}
