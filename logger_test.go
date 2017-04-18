package logger

import (
	"bytes"
	"encoding/json"
	"testing"
)

// Tests the mock logger.
func TestCreateAndLogToMockLogger(t *testing.T) {
	mockTransport := NewMockTransport()
	mockLog := NewLog(mockTransport)
	mockLog.Info("Stuff has just happened here")
	mockLog.Info("MORE Stuff has just happened here!!")
	logCount := len(mockTransport.Logs)
	logsExpected := 2
	if logCount != logsExpected {
		t.Fatalf("expected number of log entries created = %v, got: %v", logsExpected, logCount)
	}
}

func TestStdOut(t *testing.T) {
	var buf bytes.Buffer
	out := WriterTransport{W: &buf}
	mockLog := NewLog(out)
	mockLog.Info("WriterMessage")

	res := make(map[string]interface{})
	json.Unmarshal(buf.Bytes(), &res)
	if res["message"] != "WriterMessage" {
		t.Fatalf("logging to Writer. Expected message of 'WriterMessage', got %s.", res["message"])
	}
}

func TestSetContext(t *testing.T) {
	mockTransport := NewMockTransport()
	mockLog := NewLog(mockTransport)
	mockLog.SetContext("TEST-CONTEXT")
	mockLog.Info("Stuff has just happened here")
	mockLog.Info("MORE Stuff has just happened here!!")
}
