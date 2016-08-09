package logger

import "testing"

// Tests the mock logger.
func TestCreateAndLogToMockLogger(t *testing.T) {
	mockLogger := NewMockLogger()
	mockLogger.Info("Stuff has just happened here")
	mockLogger.Info("MORE Stuff has just happened here!!")
	logCount := len(mockLogger.Logs)
	logsExpected := 2
	if logCount != logsExpected {
		t.Errorf("expected number of log entries created = %v, got: %v", logsExpected, logCount)
	}
}
