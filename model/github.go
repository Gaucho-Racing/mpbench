package model

import "time"

type GithubCheckSuiteEvent struct {
	Action     string `json:"action"`
	CheckSuite struct {
		ID         int         `json:"id"`
		NodeID     string      `json:"node_id"`
		HeadBranch string      `json:"head_branch"`
		HeadSha    string      `json:"head_sha"`
		Status     string      `json:"status"`
		Conclusion interface{} `json:"conclusion"`
		URL        string      `json:"url"`
		Before     string      `json:"before"`
		After      string      `json:"after"`
		App        struct {
			ID          int       `json:"id"`
			ClientID    string    `json:"client_id"`
			Slug        string    `json:"slug"`
			NodeID      string    `json:"node_id"`
			Name        string    `json:"name"`
			Description string    `json:"description"`
			ExternalURL string    `json:"external_url"`
			HTMLURL     string    `json:"html_url"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
		} `json:"app"`
		CreatedAt            time.Time `json:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"`
		Rerequestable        bool      `json:"rerequestable"`
		RunsRerequestable    bool      `json:"runs_rerequestable"`
		LatestCheckRunsCount int       `json:"latest_check_runs_count"`
		CheckRunsURL         string    `json:"check_runs_url"`
		HeadCommit           struct {
			ID        string    `json:"id"`
			TreeID    string    `json:"tree_id"`
			Message   string    `json:"message"`
			Timestamp time.Time `json:"timestamp"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
			Committer struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"committer"`
		} `json:"head_commit"`
	} `json:"check_suite"`
	Sender struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		UserViewType      string `json:"user_view_type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"sender"`
	Installation struct {
		ID     int    `json:"id"`
		NodeID string `json:"node_id"`
	} `json:"installation"`
}

type GithubCheckRunEvent struct {
	Action   string `json:"action"`
	CheckRun struct {
		ID          int         `json:"id"`
		Name        string      `json:"name"`
		NodeID      string      `json:"node_id"`
		HeadSha     string      `json:"head_sha"`
		ExternalID  string      `json:"external_id"`
		URL         string      `json:"url"`
		HTMLURL     string      `json:"html_url"`
		DetailsURL  string      `json:"details_url"`
		Status      string      `json:"status"`
		Conclusion  interface{} `json:"conclusion"`
		StartedAt   time.Time   `json:"started_at"`
		CompletedAt interface{} `json:"completed_at"`
		Output      struct {
			Title            interface{} `json:"title"`
			Summary          interface{} `json:"summary"`
			Text             interface{} `json:"text"`
			AnnotationsCount int         `json:"annotations_count"`
			AnnotationsURL   string      `json:"annotations_url"`
		} `json:"output"`
		CheckSuite struct {
			ID           int         `json:"id"`
			NodeID       string      `json:"node_id"`
			HeadBranch   string      `json:"head_branch"`
			HeadSha      string      `json:"head_sha"`
			Status       string      `json:"status"`
			Conclusion   interface{} `json:"conclusion"`
			URL          string      `json:"url"`
			Before       string      `json:"before"`
			After        string      `json:"after"`
			PullRequests []struct {
				URL    string `json:"url"`
				ID     int    `json:"id"`
				Number int    `json:"number"`
				Head   struct {
					Ref  string `json:"ref"`
					Sha  string `json:"sha"`
					Repo struct {
						ID   int    `json:"id"`
						URL  string `json:"url"`
						Name string `json:"name"`
					} `json:"repo"`
				} `json:"head"`
				Base struct {
					Ref  string `json:"ref"`
					Sha  string `json:"sha"`
					Repo struct {
						ID   int    `json:"id"`
						URL  string `json:"url"`
						Name string `json:"name"`
					} `json:"repo"`
				} `json:"base"`
			} `json:"pull_requests"`
			App struct {
				ID          int       `json:"id"`
				ClientID    string    `json:"client_id"`
				Slug        string    `json:"slug"`
				NodeID      string    `json:"node_id"`
				Name        string    `json:"name"`
				Description string    `json:"description"`
				ExternalURL string    `json:"external_url"`
				HTMLURL     string    `json:"html_url"`
				CreatedAt   time.Time `json:"created_at"`
				UpdatedAt   time.Time `json:"updated_at"`
			} `json:"app"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"check_suite"`
		App struct {
			ID          int       `json:"id"`
			ClientID    string    `json:"client_id"`
			Slug        string    `json:"slug"`
			NodeID      string    `json:"node_id"`
			Name        string    `json:"name"`
			Description string    `json:"description"`
			ExternalURL string    `json:"external_url"`
			HTMLURL     string    `json:"html_url"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
		} `json:"app"`
		PullRequests []struct {
			URL    string `json:"url"`
			ID     int    `json:"id"`
			Number int    `json:"number"`
			Head   struct {
				Ref  string `json:"ref"`
				Sha  string `json:"sha"`
				Repo struct {
					ID   int    `json:"id"`
					URL  string `json:"url"`
					Name string `json:"name"`
				} `json:"repo"`
			} `json:"head"`
			Base struct {
				Ref  string `json:"ref"`
				Sha  string `json:"sha"`
				Repo struct {
					ID   int    `json:"id"`
					URL  string `json:"url"`
					Name string `json:"name"`
				} `json:"repo"`
			} `json:"base"`
		} `json:"pull_requests"`
	} `json:"check_run"`
	Sender struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		UserViewType      string `json:"user_view_type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"sender"`
	Installation struct {
		ID     int    `json:"id"`
		NodeID string `json:"node_id"`
	} `json:"installation"`
}

type CheckRunPayload struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	HeadSHA     string `json:"head_sha,omitempty"`
	Status      string `json:"status,omitempty"`
	Conclusion  string `json:"conclusion,omitempty"`
	ExternalID  string `json:"external_id,omitempty"`
	DetailsURL  string `json:"details_url,omitempty"`
	StartedAt   string `json:"started_at,omitempty"`
	CompletedAt string `json:"completed_at,omitempty"`
	Output      struct {
		Title   string `json:"title,omitempty"`
		Summary string `json:"summary,omitempty"`
		Text    string `json:"text,omitempty"`
	} `json:"output,omitempty"`
}
