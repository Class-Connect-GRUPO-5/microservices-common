package logger

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger(logLevel string, filePath ...string) error {
	Logger = logrus.New()

	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Logger.SetReportCaller(true)

	err := setLogLevel(logLevel, Logger)
	if err != nil {
		return err
	}
	err = setOutput(filePath, Logger)
	if err != nil {
		return err
	}

	return nil
}

// setOutput configures the output destination for the provided logger.
// If a file path is provided in the filePath slice, the function attempts
// to open the file in append mode (creating it if it does not exist) and
// sets the logger to write to both the file and standard output.
// If no file path is provided, the logger will write to standard output only.
// If an error occurs while opening the file, the function returns an error.
// Parameters:
//   - filePath: A slice of strings representing the file path(s) for log output.
//   - logger: A pointer to the logrus.Logger instance to configure.
//
// Returns:
//   - error: An error if there is an issue opening the file, otherwise nil.
//
// Example:
//
//	err := setOutput([]string{"app.log"}, logger)
//	if err != nil {
//	    log.Fatalf("Failed to set logger output: %v", err)
//	}
func setOutput(filePath []string, logger *logrus.Logger) error {
	if len(filePath) > 0 {
		file, err := os.OpenFile(filePath[0], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}
		logger.SetOutput(io.MultiWriter(file, os.Stdout))
	} else {
		logger.SetOutput(os.Stdout)
	}
	return nil
}

// setLogLevel sets the logging level for the provided logrus.Logger instance
// based on the given logLevel string. The logLevel string is case-insensitive
// and can be one of the following values:
//   - "" or "info": Sets the log level to Info.
//   - "debug": Sets the log level to Debug.
//   - "warn": Sets the log level to Warn.
//   - "error": Sets the log level to Error.
//   - "fatal": Sets the log level to Fatal.
//   - "panic": Sets the log level to Panic.
//
// If an invalid logLevel is provided, the function returns an error indicating
// the invalid value.
//
// Parameters:
//   - logLevel: A string representing the desired log level.
//   - logger: A pointer to a logrus.Logger instance whose log level will be set.
//
// Returns:
//   - error: An error if the logLevel is invalid, otherwise nil.
func setLogLevel(logLevel string, logger *logrus.Logger) error {
	switch strings.ToLower(logLevel) {
	case "", "info":
		logger.SetLevel(logrus.InfoLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	default:
		return fmt.Errorf("invalid log level: %s", logLevel)
	}
	return nil
}
