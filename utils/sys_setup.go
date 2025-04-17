package utils

import (
	"os"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger"

	"github.com/gin-gonic/gin"
)

// readEnvironmentVariables retrieves necessary environment variables for the application.
// It returns the host, port, environment, log level, and secret key.
// If any variable is missing, it uses default values for host, port, and environment.
func ReadEnvironmentVariables() (string, string, string, string, string, string) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	environment := os.Getenv("ENVIRONMENT")
	logLevel := os.Getenv("LOG_LEVEL")
	secret := os.Getenv("SECRET")
	mailSenderUrl := os.Getenv("NOTIFICATIONS_URL")

	// Set default values if environment variables are not set
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "8080"
	}
	if environment == "" {
		environment = "development"
	}

	return host, port, environment, logLevel, secret, mailSenderUrl
}

// getRouter initializes the Gin router with recovery middleware and debug logging.
// It returns a pointer to the router.
func GetRouter() *gin.Engine {
	logger.Logger.Debug("Initializing router")
	r := gin.New()
	r.Use(gin.Recovery())
	logger.Logger.Debug("Router initialized successfully")
	return r
}
