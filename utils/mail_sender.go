package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger"
)

type Sender interface {
	SendVerificationEmail(email string, pin string, name string) error
}

// MailSender struct that implements the Sender interface
type MailSender struct {
	notification_service_url string
}

func NewMailSender(notification_service_url string) MailSender {
	return MailSender{
		notification_service_url: notification_service_url,
	}
}

// Implementation of SendVerificationEmail for MailSender
func (ms *MailSender) SendVerificationEmail(email string, pin string, name string) error {
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
		return fmt.Errorf("error preparing email request: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", ms.notification_service_url, bytes.NewBuffer(jsonBody))
	if err != nil {
		logger.Logger.Errorf("Error creating HTTP request: %v", err)
		return fmt.Errorf("error creating email request: %w", err)
	}

	// Get the mail key from environment variables
	mailKey := os.Getenv("MAIL_KEY")
	if mailKey == "" {
		logger.Logger.Error("MAIL_KEY environment variable is not set")
		return fmt.Errorf("MAIL_KEY environment variable is not set")
	}
	req.Header.Set("Key", mailKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Logger.Errorf("Error sending email request: %v", err)
		return fmt.Errorf("error sending email request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status is 201 Created
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		logger.Logger.Errorf("Email service returned non-OK status: %d, body: %s", resp.StatusCode, string(body))
		return fmt.Errorf("email service returned non-OK status: %d", resp.StatusCode)
	}

	logger.Logger.Infof("Verification email sent successfully to %s", email)
	return nil
}
