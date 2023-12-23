package log

import (
	"fmt"
	"io"
	"log"
)

var _ Logger = (*stdLogger)(nil)

type stdLogger struct {
	fields []LogField
	log    *log.Logger
}

func (s *stdLogger) Debug(v ...any) {
	prefix := s.prefix()
	s.log.Printf(prefix)
	s.log.Println(v...)
}

func (s *stdLogger) Debugf(format string, v ...any) {
	prefix := s.prefix()
	format = prefix + format
	s.log.Printf(format, v...)
}

func (s *stdLogger) Info(v ...any) {
	prefix := s.prefix()
	s.log.Printf(prefix)
	s.log.Println(v...)
}

func (s *stdLogger) Infof(format string, v ...any) {
	prefix := s.prefix()
	format = prefix + format
	s.log.Printf(format, v...)
}

func (s *stdLogger) Warn(v ...any) {
	prefix := s.prefix()
	s.log.Printf(prefix)
	s.log.Println(v...)
}

func (s *stdLogger) Warnf(format string, v ...any) {
	prefix := s.prefix()
	format = prefix + format
	s.log.Printf(format, v...)
}

func (s *stdLogger) Error(v ...any) {
	prefix := s.prefix()
	s.log.Printf(prefix)
	s.log.Println(v...)
}

func (s *stdLogger) Errorf(format string, v ...any) {
	prefix := s.prefix()
	format = prefix + format
	s.log.Printf(format, v...)
}

func (s *stdLogger) prefix() string {
	if s.fields == nil {
		return ""
	}
	var prefix = "["
	for _, field := range s.fields {
		prefix += fmt.Sprintf("%s=%v", field.Key, field.Value)
	}
	prefix += "] "
	return prefix
}

func (s *stdLogger) With(fields ...LogField) Logger {
	var l = NewStdLogger(s.log.Writer(), fields...)
	return l
}

// NewStdLogger new a logger with writer.
func NewStdLogger(w io.Writer, fields ...LogField) Logger {
	return &stdLogger{
		log:    log.New(w, "", 0),
		fields: fields,
	}
}
