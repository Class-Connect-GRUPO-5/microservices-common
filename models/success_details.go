package models

// SuccessDetails represents a standardized HTTP success response.
//
// It mirrors the structure of ProblemDetails (RFC 7807), but for successful responses,
// providing consistency and extensibility across API endpoints.
//
// Fields:
//   - Type: A URI reference that categorizes the success type. Defaults to "about:blank".
//   - Title: A short, human-readable summary of the success.
//   - Status: The HTTP status code associated with the response.
//   - Message: A detailed explanation of the successful operation.
//   - Instance: A URI reference that identifies the specific occurrence of the success.
//   - Data: A stringified JSON representing the payload.
type SuccessDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Instance string `json:"instance"`
	Data     string `json:"data"`
}

// NewSuccessDetails creates a new SuccessDetails instance with the given parameters.
//
// Parameters:
//   - status: The HTTP status code (e.g., 200, 201).
//   - title: A short summary of the success.
//   - message: A detailed message describing the successful outcome.
//   - instance: A URI reference that identifies the specific occurrence.
//   - data: The JSON string containing the payload.
//
// Returns:
//
//	A populated SuccessDetails struct.
func NewSuccessDetails(status int, title, message, instance, data string) SuccessDetails {
	return SuccessDetails{
		Type:     "about:blank",
		Title:    title,
		Status:   status,
		Message:  message,
		Instance: instance,
		Data:     data,
	}
}

// GetStatus returns the HTTP status code.
func (s SuccessDetails) GetStatus() int { return s.Status }

// GetType returns the type URI.
func (s SuccessDetails) GetType() string { return s.Type }

// GetTitle returns the title.
func (s SuccessDetails) GetTitle() string { return s.Title }

// GetInstance returns the instance URI.
func (s SuccessDetails) GetInstance() string { return s.Instance }

// GetData returns the data payload.
func (s SuccessDetails) GetData() string { return s.Data }
