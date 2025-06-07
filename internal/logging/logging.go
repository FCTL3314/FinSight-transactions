package logging

const (
	LogsDir              = "logs"
	GeneralLogsFile      = "general.json"
	TransactionsLogsFile = "transactions.json"
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

type LoggerGroup struct {
	General     Logger
	Transaction Logger
}

func NewLoggerGroup(generalLogger Logger, transactionLogger Logger) *LoggerGroup {
	return &LoggerGroup{
		General:     generalLogger,
		Transaction: transactionLogger,
	}
}
