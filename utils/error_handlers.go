package utils

import (
	"net/http"

	"github.com/Class-Connect-GRUPO-5/microservices-common/models"
	"github.com/gin-gonic/gin"
)

// HandleInternalServerError is a utility function that handles internal server errors
// by sending a JSON response with a 500 status code and a generic error message.
// It uses the Gin framework's context to format and send the response.
//
// Parameters:
//   - c: *gin.Context - The Gin context used to send the JSON response.
//
// Behavior:
//   - It sends a JSON response with a 500 status code and a generic error message.
//   - The response includes the request path where the error occurred.
//   - It uses the models.InternalServerError function to format the error response.
//   - The function does not return any value.
func HandleInternalServerError(c *gin.Context) {
	HandleError(c, http.StatusInternalServerError, "Internal Server error", c.Request.URL.Path)
}

// HandleError is a utility function that handles HTTP errors by sending a JSON response
// with the appropriate status code and error message. It uses the Gin framework's context
// to format and send the response.
//
// Parameters:
//   - c: *gin.Context - The Gin context used to send the JSON response.
//   - statusCode: int - The HTTP status code to be sent in the response.
//   - errorMessage: string - A descriptive error message to include in the response.
//   - path: string - The request path where the error occurred.
//
// Behavior:
//   - Depending on the provided statusCode, the function sends a JSON response
//     with a corresponding error model (e.g., BadRequest, Unauthorized, etc.).
//   - If the statusCode does not match any predefined cases, it defaults to sending
//     an InternalServerError response.
func HandleError(c *gin.Context, statusCode int, errorMessage string, path string) {
	switch statusCode {
	case http.StatusBadRequest:
		c.IndentedJSON(statusCode, models.BadRequest(
			errorMessage,
			path,
		))
	case http.StatusUnauthorized:
		c.IndentedJSON(statusCode, models.Unauthorized(
			errorMessage,
			path,
		))
	case http.StatusForbidden:
		c.IndentedJSON(statusCode, models.Forbidden(
			errorMessage,
			path,
		))
	case http.StatusNotFound:
		c.IndentedJSON(statusCode, models.NotFound(
			errorMessage,
			path,
		))
	default:
		c.IndentedJSON(statusCode, models.InternalServerError(
			errorMessage,
			path,
		))
	}
}
