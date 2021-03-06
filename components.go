package oasm

// Holds a set of reusable objects for different aspects of the OAS.
// All objects defined within the components object will have no effect on the API unless they are explicitly
// referenced from properties outside the components object.
type Components struct {

	// An object to hold reusable Schema Objects.
	Schemas map[string]interface{} `json:"schemas,omitempty"`

	// An object to hold reusable Responses Objects.
	Responses map[string]Response `json:"responses,omitempty"`

	// An object to hold reusable Parameters Objects.
	Parameters map[string]Parameter `json:"parameters,omitempty"`

	// An object to hold reusable Examples Objects.
	Examples map[string]Example `json:"examples,omitempty"`

	// An object to hold reusable Request Body Objects.
	RequestBodies map[string]RequestBody `json:"requestBodies,omitempty"`

	// An object to hold reusable Header Objects.
	Headers map[string]Header `json:"headers,omitempty"`

	// An object to hold reusable Security Scheme Objects.
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"`

	// An object to hold reusable Link Objects.
	Links map[string]Link `json:"links,omitempty"`

	// An object to hold reusable Callback Objects.
	Callbacks map[string]CallbacksMap `json:"callbacks,omitempty"`
}
