// Package utils provides utility functions and components for the users microservice.
package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger"
	model "github.com/Class-Connect-GRUPO-5/microservices-common/models"

	"github.com/gin-gonic/gin"
)

// Sender defines the interface for email sending functionality.
// This interface allows for different implementations of email sending
// services and makes testing easier through mocking.
type Sender interface {
	// SendVerificationEmail sends a verification email with a PIN to the specified email address.
	// It returns an APIResponse indicating success or failure.
	//
	// Parameters:
	// - email: the recipient's email address
	// - pin: the verification PIN to include in the email
	// - name: the recipient's name to personalize the email
	// - ctx: the Gin context for handling the response
	SendVerificationEmail(email, pin, name string, ctx *gin.Context) model.APIResponse
}

// MailSender implements the Sender interface by communicating with an external
// notification service to deliver emails.
type MailSender struct {
	notification_service_url string // URL endpoint of the notification service API
}

// NewMailSender creates and returns a new MailSender instance configured with
// the provided notification service URL.
//
// Parameters:
// - notification_service_url: the URL endpoint for the notification service
func NewMailSender(notification_service_url string) MailSender {
	return MailSender{
		notification_service_url: notification_service_url,
	}
}

// SendVerificationEmail implements the Sender interface by sending a verification
// email with PIN through an external notification service.
//
// The method sends a POST request to the notification service with the email,
// PIN, and name in the request body. It requires a MAIL_KEY environment variable
// to be set for authentication with the notification service.
//
// Parameters:
// - email: the recipient's email address
// - pin: the verification PIN to include in the email
// - name: the recipient's name to personalize the email
// - ctx: the Gin context for handling the response
//
// Returns:
//   - model.APIResponse: a response object indicating success (200 OK) or containing
//     error details if the operation failed
func (ms *MailSender) SendVerificationEmail(email, pin, name string, ctx *gin.Context) model.APIResponse {
	logger.Logger.Infof("Sending verification email to %s with PIN: %s", email, pin)

	// Prepare the request body
	requestBody := map[string]string{
		"pin":   pin,
		"email": email,
		"name":  name,
	}

	// Marshal the request body into JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		logger.Logger.Errorf("Error marshalling request body: %v", err)
		return model.NewProblemDetails(http.StatusInternalServerError, "Internal Server Error", "Failed to marshal request body", ctx.Request.URL.Path)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", ms.notification_service_url, bytes.NewBuffer(jsonBody))
	if err != nil {
		logger.Logger.Errorf("Error creating HTTP request: %v", err)
		return model.NewProblemDetails(http.StatusInternalServerError, "Internal Server Error", "Failed to create HTTP request", ctx.Request.URL.Path)
	}

	// Get the mail key from environment variables
	mailKey := os.Getenv("MAIL_KEY")
	if mailKey == "" {
		logger.Logger.Error("MAIL_KEY environment variable is not set")
		return model.NewProblemDetails(http.StatusInternalServerError, "Internal Server Error", "Internal Server Error", ctx.Request.URL.Path)
	}
	req.Header.Set("Key", mailKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Logger.Errorf("Error sending email request: %v", err)
		return model.NewProblemDetails(http.StatusInternalServerError, "Internal Server Error", "Failed to send email request", ctx.Request.URL.Path)
	}
	defer resp.Body.Close()

	// Check if the response status is 201 Created
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		logger.Logger.Errorf("Email service returned non-OK status: %d, body: %s", resp.StatusCode, string(body))
		return model.NewProblemDetails(resp.StatusCode, "Email Sender Error", "Failed to send email", ctx.Request.URL.Path)
	}

	logger.Logger.Infof("Verification email sent successfully to %s", email)
	return model.NewSuccessDetails(http.StatusOK, "Email sent successfully", "Verification email sent successfully", ctx.Request.URL.Path, "")
}
