package logging

import "log"

type LoggerFactory struct {
	baseDir string
}

func NewLoggerFactory(baseDir string) *LoggerFactory {
	return &LoggerFactory{baseDir: baseDir}
}

func InitGeneralLogger() Logger {
	factory := NewLoggerFactory(LogsDir)
	logger, err := factory.NewZapLogger(GeneralLogsFile)
	if err != nil {
		log.Fatalf("Failed to initialize general logger: %v", err)
	}
	return logger
}

func InitTransactionLogger() Logger {
	factory := NewLoggerFactory(LogsDir)
	logger, err := factory.NewZapLogger(TransactionsLogsFile)
	if err != nil {
		log.Fatalf("Failed to initialize transaction logger: %v", err)
	}
	return logger
}
