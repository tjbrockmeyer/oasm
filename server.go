package oasm

import (
	"fmt"
)

// An object representing a Server.
type ServerDoc struct {

	// REQUIRED. A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the
	// host location is relative to the location where the OpenAPI document is being served. Variable substitutions will
	// be made when a variable is named in {brackets}.
	Url string `json:"url"`

	// An optional string describing the host designated by the URL.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// A map between a variable name and its value. The value is used for substitution in the server's URL template.
	Variables map[string]*ServerVariableDoc `json:"variables,omitempty"`
}

func (s *ServerDoc) Validate() error {
	if s.Url == "" {
		return fmt.Errorf("for servers, 'url' is required")
	}
	return nil
}
