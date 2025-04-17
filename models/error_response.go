package models

import "net/http"

type ProblemDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

// NewProblemDetails creates a new instance of ProblemDetails with the provided parameters.
// It initializes the Type field to "about:blank" and sets the other fields based on the input.
//
// Parameters:
//   - status: The HTTP status code associated with the problem.
//   - title: A short, human-readable summary of the problem.
//   - detail: A detailed explanation of the problem.
//   - instance: A URI reference that identifies the specific occurrence of the problem.
//
// Returns:
//
//	A ProblemDetails struct populated with the provided values.
func NewProblemDetails(status int, title, detail, instance string) ProblemDetails {
	return ProblemDetails{
		Type:     "about:blank",
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
	}
}

func BadRequest(detail, instance string) ProblemDetails {
	return NewProblemDetails(http.StatusBadRequest, "Invalid request", detail, instance)
}

func InternalServerError(detail, instance string) ProblemDetails {
	return NewProblemDetails(http.StatusInternalServerError, "Internal server error", detail, instance)
}

func NotFound(detail, instance string) ProblemDetails {
	return NewProblemDetails(http.StatusNotFound, "Not Found", detail, instance)
}

func Forbidden(detail, instance string) ProblemDetails {
	return NewProblemDetails(http.StatusForbidden, "Forbidden", detail, instance)
}

func Unauthorized(detail, instance string) ProblemDetails {
	return NewProblemDetails(http.StatusUnauthorized, "Unauthorized", detail, instance)
}
