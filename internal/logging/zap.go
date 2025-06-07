package logging

import (
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
	mu     sync.RWMutex
}

func toZapFields(fields ...Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		zapFields[i] = zap.Any(field.Key, field.Value)
	}
	return zapFields
}

func (l *ZapLogger) With(fields ...Field) Logger {
	l.mu.RLock()
	defer l.mu.RUnlock()

	newLogger := l.logger.With(toZapFields(fields...)...)
	return &ZapLogger{logger: newLogger}
}

func (l *ZapLogger) Debug(msg string, fields ...Field) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	l.logger.Debug(msg, toZapFields(fields...)...)
}

func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	l.logger.Info(msg, toZapFields(fields...)...)
}

func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	l.logger.Warn(msg, toZapFields(fields...)...)
}

func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	l.logger.Error(msg, toZapFields(fields...)...)
}

func (l *ZapLogger) Fatal(msg string, fields ...Field) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	l.logger.Fatal(msg, toZapFields(fields...)...)
}

func (f *LoggerFactory) NewZapLogger(fileName string) (Logger, error) {
	logPath := filepath.Join(f.baseDir, fileName)

	logDir := filepath.Dir(logPath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	level := zap.NewAtomicLevelAt(zapcore.DebugLevel)

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	fileWriter := zapcore.AddSync(file)
	consoleWriter := zapcore.AddSync(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWriter, level),
		zapcore.NewCore(consoleEncoder, consoleWriter, level),
	)

	zapLogger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	).With(zap.String("log_path", logPath))

	return &ZapLogger{logger: zapLogger}, nil
}
