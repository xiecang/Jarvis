package log

import (
	"fmt"
	"log"
	"time"
)

// DefaultLogger is default logger.
var DefaultLogger = NewStdLogger(log.Writer())

type LogField struct {
	Key   string
	Value any
}

// Field returns a LogField for the given key and value.
// value must be string, int ... basic value
func Field(key string, value any) LogField {
	switch val := value.(type) {
	case error:
		return LogField{Key: key, Value: val.Error()}
	case []error:
		var errs []string
		for _, err := range val {
			errs = append(errs, err.Error())
		}
		return LogField{Key: key, Value: errs}
	case time.Duration:
		return LogField{Key: key, Value: fmt.Sprint(val)}
	case []time.Duration:
		var durs []string
		for _, dur := range val {
			durs = append(durs, fmt.Sprint(dur))
		}
		return LogField{Key: key, Value: durs}
	case []time.Time:
		var times []string
		for _, t := range val {
			times = append(times, fmt.Sprint(t))
		}
		return LogField{Key: key, Value: times}
	case fmt.Stringer:
		return LogField{Key: key, Value: val.String()}
	case []fmt.Stringer:
		var strs []string
		for _, str := range val {
			strs = append(strs, str.String())
		}
		return LogField{Key: key, Value: strs}
	default:
		return LogField{Key: key, Value: val}
	}
}

// Logger logger interface
type Logger interface {
	Debug(v ...any)
	Debugf(format string, v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Warn(v ...any)
	Warnf(format string, v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
	With(fields ...LogField) Logger
}
