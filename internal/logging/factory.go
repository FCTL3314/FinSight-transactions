package logging

import "log"

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
