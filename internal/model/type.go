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
		Job         string
		Company     string
		Url         string
		Start       string
		End         string
		Description string
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
		Suffix      string
		Host        string
		Competition string
		Translation string
		Place       int
	}

	// The Certification struct represents the certification details.
	Certification struct {
		Title         string
		CredentialId  string
		CredentialUrl string
	}

	// The Blog struct represents the response from the dev.to API for articles.
	Blog struct {
		CreatedAt              time.Time `json:"created_at"`
		LastCommentAt          time.Time `json:"last_comment_at"`
		PublishedAt            time.Time `json:"published_at"`
		CrosspostedAt          time.Time `json:"crossposted_at"`
		EditedAt               time.Time `json:"edited_at"`
		User                   BlogUser  `json:"user"`
		Path                   string    `json:"path"`
		CanonicalUrl           string    `json:"canonical_url"`
		Tags                   string    `json:"tags"`
		Title                  string    `json:"title"`
		Description            string    `json:"description"`
		PublishedTimestamp     string    `json:"published_timestamp"`
		ReadablePublishDate    string    `json:"readable_publish_date"`
		CoverImage             string    `json:"cover_image"`
		SocialImage            string    `json:"social_image"`
		Url                    string    `json:"url"`
		TypeOf                 string    `json:"type_of"`
		Slug                   string    `json:"slug"`
		TagList                []string  `json:"tag_list"`
		PositiveReactionsCount int       `json:"positive_reactions_count"`
		CollectionID           int       `json:"collection_id"`
		PublicReactionsCount   int       `json:"public_reactions_count"`
		ReadingTimeMinutes     int       `json:"reading_time_minutes"`
		CommentsCount          int       `json:"comments_count"`
		ID                     int       `json:"id"`
	}

	// The BlogUser struct represents the user details of the author of the article.
	BlogUser struct {
		Name            string `json:"name"`
		Username        string `json:"username"`
		TwitterUsername string `json:"twitter_username"`
		GithubUsername  string `json:"github_username"`
		WebsiteUrl      string `json:"website_url"`
		ProfileImage    string `json:"profile_image"`
		ProfileImage90  string `json:"profile_image_90"`
		UserID          int    `json:"user_id"`
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
