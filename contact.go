package oasm

// Contact information for the exposed API.
type ContactDoc struct {

	// The identifying name of the contact person/organization.
	Name string `json:"name,omitempty"`

	// The URL pointing to the contact information. MUST be in the format of a URL.
	Url string `json:"url,omitempty"`

	// The email address of the contact person/organization. MUST be in the format of an email address.
	Email string `json:"email,omitempty"`
}
