package model

type (
	// The Social struct represents the social media links.
	Social struct {
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
)
