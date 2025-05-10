package logger

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger/events"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
	Panic
)

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	case Panic:
		return "panic"
	default:
		return "unreachable"
	}
}

func LogLevelFromString(s string) (LogLevel, error) {
	switch s {
	case "debug":
		return Debug, nil
	case "info":
		return Info, nil
	case "warn":
		return Warn, nil
	case "error":
		return Error, nil
	case "fatal":
		return Fatal, nil
	case "panic":
		return Panic, nil
	default:
		return -1, fmt.Errorf("invalid level")
	}
}

type LoggerI interface {
	SetLogLevel(logLevel LogLevel) error
	Log(level LogLevel, msg string)
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Panic(msg string)
	Logf(level LogLevel, format string, fields ...interface{})
	Debugf(format string, fields ...interface{})
	Infof(format string, fields ...interface{})
	Warnf(format string, fields ...interface{})
	Errorf(format string, fields ...interface{})
	Fatalf(format string, fields ...interface{})
	Panicf(format string, fields ...interface{})
	Emit(event events.Event)
}

var Logger LoggerI

type logger struct {
	name     string
	level    LogLevel
	logrus   *logrus.Logger
	rabbitmq RabbitMQ
}

type RabbitMQConfig struct {
	Host string
	Port uint16
}

func Url(c RabbitMQConfig) string {
	return fmt.Sprintf("amqp://guest:guest@%s:%d/", c.Host, c.Port)
}

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

const logExchangeName = "logs"
const statsExchangeName = "stats"

func (l *logger) connectRabbitMQ(config RabbitMQConfig) error {
	var conn *amqp.Connection
	for {
		var err error
		conn, err = amqp.Dial(Url(config))
		if err != nil {
			l.Infof("Failed to connect to RabbitMQ: %v", err)
		} else {
			l.Infof("Connected to RabbitMQ")
			break
		}
		time.Sleep(time.Second * 5)
		l.Infof("Retrying connection...")
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	err = ch.ExchangeDeclare(
		logExchangeName,
		"fanout",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %v", err)
	}
	l.rabbitmq = RabbitMQ{conn, ch}
	return nil
}

func InitLogger(name string, logLevel LogLevel, output io.Writer, remote bool) error {
	logrus_instance := logrus.New()

	logrus_instance.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus_instance.SetReportCaller(true)

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
		err = l.connectRabbitMQ(RabbitMQConfig{
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

func (l *logger) Log(level LogLevel, msg string) {
	l.logrusLog(level, msg)

	err := l.rabbitmq.Log(level, msg)
	if err != nil {
		l.logrusLog(Error, fmt.Sprintf("failed to emit event to rabbitMQ: %v", err))
	}
}

func (l *logger) logrusLog(level LogLevel, msg string) {
	if level >= l.level {
		switch level {
		case Info:
			l.logrus.Info(msg)
		case Debug:
			l.logrus.Debug(msg)
		case Warn:
			l.logrus.Warn(msg)
		case Error:
			l.logrus.Error(msg)
		case Fatal:
			l.logrus.Fatal(msg)
		case Panic:
			l.logrus.Panic(msg)
		}
	}
}

func (r *RabbitMQ) Log(level LogLevel, msg string) error {
	return r.Send(logExchangeName, amqp.Table{"level": level.String()}, []byte(msg))
}

func (r *RabbitMQ) Send(exchange string, headers amqp.Table, body []byte) error {
	if r.ch == nil {
		return nil
	}
	return r.ch.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			Headers:     headers,
			ContentType: "text/plain",
			Body:        body,
		},
	)
}

func (l *logger) Debug(msg string) {
	l.Log(Debug, msg)
}

func (l *logger) Info(msg string) {
	l.Log(Info, msg)
}

func (l *logger) Warn(msg string) {
	l.Log(Warn, msg)
}

func (l *logger) Error(msg string) {
	l.Log(Error, msg)
}

func (l *logger) Fatal(msg string) {
	l.Log(Fatal, msg)
}

func (l *logger) Panic(msg string) {
	l.Log(Panic, msg)
}

func (l *logger) Logf(level LogLevel, format string, fields ...interface{}) {
	l.Log(level, fmt.Sprintf(format, fields...))
}

func (l *logger) Debugf(format string, fields ...interface{}) {
	l.Logf(Debug, format, fields...)
}

func (l *logger) Infof(format string, fields ...interface{}) {
	l.Logf(Info, format, fields...)
}

func (l *logger) Warnf(format string, fields ...interface{}) {
	l.Logf(Warn, format, fields...)
}

func (l *logger) Errorf(format string, fields ...interface{}) {
	l.Logf(Error, format, fields...)
}

func (l *logger) Fatalf(format string, fields ...interface{}) {
	l.Logf(Fatal, format, fields...)
}

func (l *logger) Panicf(format string, fields ...interface{}) {
	l.Logf(Panic, format, fields...)
}

func (l *logger) Emit(event events.Event) {
	l.rabbitmq.sendEvent(event)
}

func (r RabbitMQ) sendEvent(event events.Event) {
	b, err := event.Encode()
	if err != nil {
		panic(err)
	}
	r.Send(statsExchangeName, amqp.Table{"type": event.Type()}, b)
}

func EncodeString(s string) []byte {
	b := make([]byte, 0)
	b = binary.BigEndian.AppendUint16(b, uint16(len(s)))
	b = append(b, []byte(s)...)
	return b
}

func DecodeString(r io.Reader) string {
	uintbuf := make([]byte, 2)
	n, err := r.Read(uintbuf)
	if err != nil {
		panic(err)
	}
	if n != 2 {
		panic("Not enought bytes")
	}
	len := binary.BigEndian.Uint16(uintbuf)
	strbuf := make([]byte, len)
	n, err = r.Read(strbuf)
	if err != nil {
		panic(err)
	}
	if n != int(len) {
		panic("Not enought bytes")
	}
	return string(strbuf)
}
