package log

import (
	"log"
	"testing"
)

func TestStdLogger(t *testing.T) {
	var l = NewStdLogger(log.Writer())
	l.With(Field("a", "b")).Debug("test")
	l.With(Field("a", "b")).Debugf("test")
}
