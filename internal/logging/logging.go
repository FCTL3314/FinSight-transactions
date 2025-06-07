package logging

import (
	"log"
)

const (
	LogsDir               = "logs"
	TransactionsLogsDir   = "transactions"
	ControllerLogFileName = "controller.json"
)

type Logger interface {
	With(fields ...Field) Logger
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

type Field struct {
	Key   string
	Value interface{}
}

func NewField(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

func WithField(key string, value interface{}) Field {
	return NewField(key, value)
}

type LoggerGroup struct {
	Transaction Logger
}

func NewLoggerGroup(transactionsLogger Logger) *LoggerGroup {
	return &LoggerGroup{
		Transaction: transactionsLogger,
	}
}

type LoggerFactory struct {
	baseDir string
}

func NewLoggerFactory(baseDir string) *LoggerFactory {
	return &LoggerFactory{baseDir: baseDir}
}

func InitTransactionLogger() Logger {
	factory := NewLoggerFactory(LogsDir)
	logger, err := factory.NewZapLogger(TransactionsLogsDir, ControllerLogFileName)
	if err != nil {
		log.Fatalf("Failed to initialize transaction logger: %v", err)
	}
	return logger
}
