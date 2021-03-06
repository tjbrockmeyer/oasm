package oasm

// Defines a security scheme that can be used by the operations. Supported schemes are HTTP authentication, an API key
// (either as a header or as a query parameter), OAuth2's common flows (implicit, password, application and access code)
// as defined in RFC6749, and OpenID Connect Discovery.
type SecurityScheme struct {

	// REQUIRED. The type of the security scheme. Valid values are "apiKey", "http", "oauth2", "openIdConnect".
	Type string `json:"type,omitempty"`

	// A short description for security scheme. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// 	REQUIRED for apiKey. The name of the header, query or cookie parameter to be used.
	Name string `json:"name,omitempty"`

	// REQUIRED for apiKey. The location of the API key. Valid values are "query", "header" or "cookie".
	In string `json:"in,omitempty"`

	// REQUIRED for http.
	// The name of the HTTP Authorization scheme to be used in the Authorization header as defined in RFC7235.
	Scheme string `json:"scheme,omitempty"`

	// For http (bearer). A hint to the client to identify how the bearer token is formatted. Bearer tokens are usually
	// generated by an authorization server, so this information is primarily for documentation purposes.
	BearerFormat string `json:"bearerFormat,omitempty"`

	// REQUIRED for oauth2. An object containing configuration information for the flow types supported.
	Flows OAuthFlowsMap `json:"flows,omitempty"`

	// REQUIRED for openIdConnect.
	// OpenId Connect URL to discover OAuth2 configuration values. This MUST be in the form of a URL.
	OpenIdConnectUrl string `json:"openIdConnectUrl,omitempty"`
}
