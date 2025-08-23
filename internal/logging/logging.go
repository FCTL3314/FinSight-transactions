package logging

const (
	LogsDir              = "logs"
	GeneralLogsFile      = "general.json"
	TransactionsLogsFile = "transactions.json"
	DetailingLogsFile    = "detailing.json"
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

func WithError(error error) Field {
	return NewField("error", error)
}

type LoggersGroup struct {
	General     Logger
	Transaction Logger
	Detailing   Logger
}

func NewLoggersGroup(generalLogger Logger, transactionLogger Logger, detailingLogger Logger) *LoggersGroup {
	return &LoggersGroup{
		General:     generalLogger,
		Transaction: transactionLogger,
		Detailing:   detailingLogger,
	}
}
