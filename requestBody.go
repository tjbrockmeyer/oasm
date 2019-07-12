package oasm

// Describes a single request body.
type RequestBodyDoc struct {

	// A brief description of the request body. This could contain examples of use.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// REQUIRED. The content of the request body. The key is a media type or media type range and the
	// value describes it. For requests that match multiple keys, only the most specific key is applicable.
	// e.g. text/plain overrides text/*
	Content MediaTypesDoc `json:"content"`

	// Determines if the request body is required in the request. Defaults to false.
	Required bool `json:"required,omitempty"`
}
