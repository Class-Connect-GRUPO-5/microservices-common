package logger

import (
	"io"

	"github.com/Class-Connect-GRUPO-5/microservices-common/rabbitmq"
	"github.com/sirupsen/logrus"
)

func InitLogger(name string, logLevel LogLevel, output io.Writer, remote bool) error {
	logrus_instance := logrus.New()

	logrus_instance.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus_instance.SetReportCaller(false)

	l := &logger{
		name:   name,
		logrus: logrus_instance,
	}
	err := l.SetLogLevel(logLevel)
	if err != nil {
		return err
	}
	err = l.setOutput(output)
	if err != nil {
		return err
	}
	if remote {
		err = l.connectRabbitMQ(rabbitmq.Config{
			Host: "rabbitmq",
			Port: 5672,
		})
		if err != nil {
			return err
		}
	}
	Logger = l
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
//
// Returns:
//   - error: An error if there is an issue opening the file, otherwise nil.
//
// Example:
//
//	err := setOutput([]string{"app.log"})
//	if err != nil {
//	    log.Fatalf("Failed to set logger output: %v", err)
//	}
func (l *logger) setOutput(output io.Writer) error {
	// if len(filePath) > 0 {
	// 	file, err := os.OpenFile(filePath[0], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to open log file: %w", err)
	// 	}
	// 	l.logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	// } else {
	l.logrus.SetOutput(output)
	// }
	return nil
}

// SetLogLevel sets the logging level for the provided logrus.Logger instance
// based on the given logLevel. The logLevel is an enum:
//   - Info: Sets the log level to Info.
//   - Debug: Sets the log level to Debug.
//   - Warn: Sets the log level to Warn.
//   - Error: Sets the log level to Error.
//   - Fatal: Sets the log level to Fatal.
//   - Panic: Sets the log level to Panic.
//
// Parameters:
//   - logLevel: An enum representing the desired log level.
//   - logger: A pointer to a logrus.Logger instance whose log level will be set.
//
// Returns:
//   - error: An error if the logLevel is invalid, otherwise nil.
func (l *logger) SetLogLevel(logLevel LogLevel) error {
	switch logLevel {
	case Info:
		l.logrus.SetLevel(logrus.InfoLevel)
	case Debug:
		l.logrus.SetLevel(logrus.DebugLevel)
	case Warn:
		l.logrus.SetLevel(logrus.WarnLevel)
	case Error:
		l.logrus.SetLevel(logrus.ErrorLevel)
	case Fatal:
		l.logrus.SetLevel(logrus.FatalLevel)
	case Panic:
		l.logrus.SetLevel(logrus.PanicLevel)
	}
	l.level = logLevel
	return nil
}
