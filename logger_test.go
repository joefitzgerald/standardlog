package standardlog

import (
	"bytes"
	"log"
	"testing"
)

func TestLogLoggerIsStandardLogLogger(t *testing.T) {
	var l Logger
	var buf bytes.Buffer
	l = log.New(&buf, "", 0)
	if l == nil {
		t.Error("Expected Logger, got ", l)
	}

	_, ok := l.(Logger)
	if !ok {
		t.Error("Expected log.Logger is standardlog.Logger to be true, got ", ok)
	}
}
