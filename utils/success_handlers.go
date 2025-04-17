package utils

import (
	"github.com/gin-gonic/gin"
)

// HandleSuccess sends a JSON response with the provided status code, message, and data.
// It uses the Gin framework's IndentedJSON method to format the response.
//
// Parameters:
//   - c: The Gin context used to write the response.
//   - statusCode: The HTTP status code to be sent in the response.
//   - message: A string containing a message to include in the response.
//   - data: An interface{} containing the data to include in the response body.
func HandleSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	c.IndentedJSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}
