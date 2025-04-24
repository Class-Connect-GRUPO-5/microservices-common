package models

// APIResponse represents a generic response returned by an API endpoint.
// It abstracts over both successful and error responses.
type APIResponse interface {
	GetStatus() int
	GetType() string
	GetTitle() string
	GetInstance() string
}
