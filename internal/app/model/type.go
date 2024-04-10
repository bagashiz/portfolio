package model

import "time"

type (
	// The Social struct represents the social media links.
	Social struct {
		Name    string
		Href    string
		Rel     string
		FaClass string
	}

	// The Experience struct represents the work experiences details.
	Experience struct {
		Job          string
		Company      string
		Url          string
		Start        string
		End          string
		Descriptions []string
	}

	// The Education struct represents the education details.
	Education struct {
		School      string
		Url         string
		Major       string
		Start       string
		End         string
		Description string
	}

	// The Award struct represents the award details.
	Award struct {
		Place       int
		Suffix      string
		Host        string
		Competition string
		Translation string
	}

	// The Certification struct represents the certification details.
	Certification struct {
		Title         string
		CredentialId  string
		CredentialUrl string
	}

	// The Blog struct represents the response from the dev.to API for articles.
	Blog struct {
		TypeOf                 string    `json:"type_of"`
		ID                     int       `json:"id"`
		Title                  string    `json:"title"`
		Description            string    `json:"description"`
		ReadablePublishDate    string    `json:"readable_publish_date"`
		Slug                   string    `json:"slug"`
		Path                   string    `json:"path"`
		Url                    string    `json:"url"`
		CommentsCount          int       `json:"comments_count"`
		PublicReactionsCount   int       `json:"public_reactions_count"`
		CollectionID           int       `json:"collection_id"`
		PublishedTimestamp     string    `json:"published_timestamp"`
		PositiveReactionsCount int       `json:"positive_reactions_count"`
		CoverImage             string    `json:"cover_image"`
		SocialImage            string    `json:"social_image"`
		CanonicalUrl           string    `json:"canonical_url"`
		CreatedAt              time.Time `json:"created_at"`
		EditedAt               time.Time `json:"edited_at"`
		CrosspostedAt          time.Time `json:"crossposted_at"`
		PublishedAt            time.Time `json:"published_at"`
		LastCommentAt          time.Time `json:"last_comment_at"`
		ReadingTimeMinutes     int       `json:"reading_time_minutes"`
		TagList                []string  `json:"tag_list"`
		Tags                   string    `json:"tags"`
		User                   BlogUser  `json:"user"`
	}

	// The BlogUser struct represents the user details of the author of the article.
	BlogUser struct {
		Name            string `json:"name"`
		Username        string `json:"username"`
		TwitterUsername string `json:"twitter_username"`
		GithubUsername  string `json:"github_username"`
		UserID          int    `json:"user_id"`
		WebsiteUrl      string `json:"website_url"`
		ProfileImage    string `json:"profile_image"`
		ProfileImage90  string `json:"profile_image_90"`
	}

	// The Project struct represents the response from the GitHub GraphQL API for pinned projects.
	Project struct {
		Data Data `json:"data"`
	}

	// The Data struct represents the user details and pinned projects.
	Data struct {
		User User `json:"user"`
	}

	// The User struct represents the pinned projects.
	User struct {
		PinnedItems PinnedItems `json:"pinnedItems"`
	}

	// The PinnedItems struct represents the pinned projects.
	PinnedItems struct {
		Nodes []Node `json:"nodes"`
	}

	// The Node struct represents the details of the pinned project.
	Node struct {
		Name            string          `json:"name"`
		Description     string          `json:"description"`
		Url             string          `json:"url"`
		PrimaryLanguage PrimaryLanguage `json:"primaryLanguage"`
		Stargazers      Forks           `json:"stargazers"`
		Forks           Forks           `json:"forks"`
	}

	// The Forks struct represents the count of stars and forks of the project.
	Forks struct {
		TotalCount int `json:"totalCount"`
	}

	// The PrimaryLanguage struct represents the primary language of the project.
	PrimaryLanguage struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}
)
