package logger

import (
	"fmt"
	"os"
	"time"

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
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Host      string `json:"host"`
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func getBaseLogMessage() logMessage {
	host, _ := os.Hostname()
	return logMessage{Timestamp: makeTimestamp(), Host: host}
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
	logMessage := getBaseLogMessage()
	logMessage.Level = "Emergency"
	logMessage.Message = message
	mck.log(logMessage)
}

// Alert - Action must be taken immediately.
func (mck *MockLogger) Alert(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Alert"
	logMessage.Message = message
	mck.log(logMessage)
}

// Critical - Critical conditions.
func (mck *MockLogger) Critical(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Critical"
	logMessage.Message = message
	mck.log(logMessage)
}

// Error - Runtime errors that do not require immediate action but should typically be logged and monitored
func (mck *MockLogger) Error(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Error"
	logMessage.Message = message
	mck.log(logMessage)
}

// Warning - Exceptional occurrences that are not errors.
func (mck *MockLogger) Warning(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Warning"
	logMessage.Message = message
	mck.log(logMessage)
}

// Notice - Normal but significant events.
func (mck *MockLogger) Notice(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Notice"
	logMessage.Message = message
	mck.log(logMessage)
}

// Info - Interesting events.
func (mck *MockLogger) Info(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Info"
	logMessage.Message = message
	mck.log(logMessage)
}

// Debug - Detailed debug information.
func (mck *MockLogger) Debug(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Debug"
	logMessage.Message = message
	mck.log(logMessage)
}

func (mck *MockLogger) log(lm logMessage) {
	mck.Logs[len(mck.Logs)+1] = lm
}

// NewElasticSearchLogger - Create a new instance of an Elasticsearch logger
func NewElasticSearchLogger(host string, index string) *ElasticSearchLogger {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(host),
	)
	if err != nil {
		panic(err)
	}
	return &ElasticSearchLogger{Client: *client, Index: index}
}

// Emergency - System is unusable..
func (esl *ElasticSearchLogger) Emergency(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Emergency"
	logMessage.Message = message
	esl.log(logMessage)
}

// Alert - Action must be taken immediately.
func (esl *ElasticSearchLogger) Alert(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Alert"
	logMessage.Message = message
	esl.log(logMessage)
}

// Critical - Critical conditions.
func (esl *ElasticSearchLogger) Critical(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Critical"
	logMessage.Message = message
	esl.log(logMessage)
}

// Error - Runtime errors that do not require immediate action but should typically be logged and monitored
func (esl *ElasticSearchLogger) Error(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Error"
	logMessage.Message = message
	esl.log(logMessage)
}

// Warning - Exceptional occurrences that are not errors.
func (esl *ElasticSearchLogger) Warning(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Warning"
	logMessage.Message = message
	esl.log(logMessage)
}

// Notice - Normal but significant events.
func (esl *ElasticSearchLogger) Notice(message string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = "Notice"
	logMessage.Message = message
	esl.log(logMessage)
}

// Info - Interesting events.
func (esl *ElasticSearchLogger) Info(message string) {
	fmt.Println(message)
	logMessage := getBaseLogMessage()
	logMessage.Level = "Info"
	logMessage.Message = message
	esl.log(logMessage)
}

// Debug - Detailed debug information.
func (esl *ElasticSearchLogger) Debug(message string) {
	fmt.Println(message)
	logMessage := getBaseLogMessage()
	logMessage.Level = "Debug"
	logMessage.Message = message
	esl.log(logMessage)
}

// entry point for Elasticsearch
func (esl *ElasticSearchLogger) log(lm logMessage) {
	_, err := esl.Client.Index().Index(esl.Index).Type(lm.Level).BodyJson(lm).Refresh(true).Do()
	if err != nil {
		fmt.Println(err.Error())
	}
}
