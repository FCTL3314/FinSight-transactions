package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
)

func toZapFields(fields ...Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		zapFields[i] = zap.Any(field.Key, field.Value)
	}
	return zapFields
}

type ZapLogger struct {
	logger *zap.Logger
	mu     sync.RWMutex
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

func (f *LoggerFactory) NewZapLogger(module, fileName string) (Logger, error) {
	logPath := filepath.Join(f.baseDir, module, fileName)

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

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(file),
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
	)

	zapLogger := zap.New(
		fileCore,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	).With(zap.String("module", module))

	return &ZapLogger{logger: zapLogger}, nil
}
