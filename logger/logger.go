package logger

import (
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger/events"
	"github.com/Class-Connect-GRUPO-5/microservices-common/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

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
	rabbitmq rabbitmq.Client
}

const LogExchangeName = "logs"
const StatsExchangeName = "stats"

func (l *logger) connectRabbitMQ(config rabbitmq.Config) error {
	exchanges := []string{LogExchangeName, StatsExchangeName}
	c, err := rabbitmq.NewClient(l.name, config, exchanges)
	if err != nil {
		return fmt.Errorf("error connecting to rabbitmq")
	}
	l.rabbitmq = c
	return nil
}
func (l *logger) Log(level LogLevel, msg string) {
	l.logrusLog(level, msg)

	err := l.rabbitmq.Send(LogExchangeName, amqp.Table{"level": level.String()}, []byte(msg))
	if err != nil {
		l.logrusLog(Error, fmt.Sprintf("failed to emit event to rabbitMQ: %v", err))
	}
}

func (l *logger) Emit(event events.Event) {
	b, err := event.Encode()
	if err != nil {
		panic(err)
	}
	l.rabbitmq.Send(StatsExchangeName, amqp.Table{"type": event.Type()}, b)
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
