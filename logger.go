package logger

import (
	"encoding/json"

	elastic "gopkg.in/olivere/elastic.v3"
)

// Logger interface
type Logger interface {
	Emergency(message string)
	Alert(message string)
	Critical(message string)
	Error(message string)
	Warning(message string)
	Notice(message string)
	Info(message string)
	Debug(message string)
	log(lm logMessage)
}

type logMessage struct {
	Level   string
	Message string
}

// MockLogger - MockLogger is a mock version of Logger
type MockLogger struct {
	Logs map[int]logMessage
}

// ElasticSearchLogger - Elasticsearch type of Logger
type ElasticSearchLogger struct {
	Client elastic.Client
	Index  string
}

// NewMockLogger - Create new mock instance of logger
func NewMockLogger() *MockLogger {
	return &MockLogger{make(map[int]logMessage)}
}

// Emergency - System is unusable..
func (mck *MockLogger) Emergency(message string) {
	logMessage := logMessage{Level: "Emergency", Message: message}
	mck.log(logMessage)
}

// Alert - Action must be taken immediately.
func (mck *MockLogger) Alert(message string) {
	logMessage := logMessage{Level: "Alert", Message: message}
	mck.log(logMessage)
}

// Critical - Critical conditions.
func (mck *MockLogger) Critical(message string) {
	logMessage := logMessage{Level: "Critical", Message: message}
	mck.log(logMessage)
}

// Error - Runtime errors that do not require immediate action but should typically be logged and monitored
func (mck *MockLogger) Error(message string) {
	logMessage := logMessage{Level: "Error", Message: message}
	mck.log(logMessage)
}

// Warning - Exceptional occurrences that are not errors.
func (mck *MockLogger) Warning(message string) {
	logMessage := logMessage{Level: "Warning", Message: message}
	mck.log(logMessage)
}

// Notice - Normal but significant events.
func (mck *MockLogger) Notice(message string) {
	logMessage := logMessage{Level: "Notice", Message: message}
	mck.log(logMessage)
}

// Info - Interesting events.
func (mck *MockLogger) Info(message string) {
	logMessage := logMessage{Level: "Info", Message: message}
	mck.log(logMessage)
}

// Debug - Detailed debug information.
func (mck *MockLogger) Debug(message string) {
	logMessage := logMessage{Level: "Debug", Message: message}
	mck.log(logMessage)
}

func (mck *MockLogger) log(lm logMessage) {
	mck.Logs[len(mck.Logs)+1] = lm
}

// NewElasticSearchLogger - Create a new instance of an Elasticsearch logger
func NewElasticSearchLogger(host string, index string) *ElasticSearchLogger {
	client, err := elastic.NewClient(elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	return &ElasticSearchLogger{Client: *client, Index: index}
}

// Emergency - System is unusable..
func (esl *ElasticSearchLogger) Emergency(message string) {
	logMessage := logMessage{Level: "Emergency", Message: message}
	esl.log(logMessage)
}

// Alert - Action must be taken immediately.
func (esl *ElasticSearchLogger) Alert(message string) {
	logMessage := logMessage{Level: "Alert", Message: message}
	esl.log(logMessage)
}

// Critical - Critical conditions.
func (esl *ElasticSearchLogger) Critical(message string) {
	logMessage := logMessage{Level: "Critical", Message: message}
	esl.log(logMessage)
}

// Error - Runtime errors that do not require immediate action but should typically be logged and monitored
func (esl *ElasticSearchLogger) Error(message string) {
	logMessage := logMessage{Level: "Error", Message: message}
	esl.log(logMessage)
}

// Warning - Exceptional occurrences that are not errors.
func (esl *ElasticSearchLogger) Warning(message string) {
	logMessage := logMessage{Level: "Warning", Message: message}
	esl.log(logMessage)
}

// Notice - Normal but significant events.
func (esl *ElasticSearchLogger) Notice(message string) {
	logMessage := logMessage{Level: "Notice", Message: message}
	esl.log(logMessage)
}

// Info - Interesting events.
func (esl *ElasticSearchLogger) Info(message string) {
	logMessage := logMessage{Level: "Info", Message: message}
	esl.log(logMessage)
}

// Debug - Detailed debug information.
func (esl *ElasticSearchLogger) Debug(message string) {
	logMessage := logMessage{Level: "Debug", Message: message}
	esl.log(logMessage)
}

// entry point for Elasticsearch
func (esl *ElasticSearchLogger) log(lm logMessage) {
	logJSON, err := json.Marshal(lm)
	if err != nil {
		panic(err)
	}
	esl.Client.Index().Index(esl.Index).BodyJson(logJSON)
}
