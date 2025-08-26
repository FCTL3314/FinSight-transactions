package logging

import "log"

type LoggerFactory struct {
	baseDir string
}

func NewLoggerFactory(baseDir string) *LoggerFactory {
	return &LoggerFactory{baseDir: baseDir}
}

func InitGeneralLogger(baseDir string) Logger {
	factory := NewLoggerFactory(baseDir)
	logger, err := factory.NewZapLogger(GeneralLogsFile)
	if err != nil {
		log.Fatalf("Failed to initialize general logger: %v", err)
	}
	return logger
}

func InitTransactionLogger(baseDir string) Logger {
	factory := NewLoggerFactory(baseDir)
	logger, err := factory.NewZapLogger(TransactionsLogsFile)
	if err != nil {
		log.Fatalf("Failed to initialize transaction logger: %v", err)
	}
	return logger
}

func InitDetailingLogger(baseDir string) Logger {
	factory := NewLoggerFactory(baseDir)
	logger, err := factory.NewZapLogger(DetailingLogsFile)
	if err != nil {
		log.Fatalf("Failed to initialize detailing logger: %v", err)
	}
	return logger
}
