package logger

import (
	"encoding/json"
	"fmt"
	"io"
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

func (lm logMessage) String() string {
	body, _ := json.Marshal(lm)
	return string(body)
}

// Transport defines a transport for shipping logs.
type Transport interface {
	Ship(msg logMessage) error
}

// Log formats log messages and sends them on to a transport.
type Log struct {
	t Transport
}

// NewLog returns a pointer to a new Log, with the given transport.
func NewLog(t Transport) *Log {
	return &Log{
		t: t,
	}
}

func (l *Log) log(msg string, lvl string) {
	logMessage := getBaseLogMessage()
	logMessage.Level = lvl
	logMessage.Message = msg
	l.t.Ship(logMessage)
}

// Emergency - System is unusable..
func (l *Log) Emergency(message string) {
	l.log(message, "Emergency")
}

// Alert - Action must be taken immediately.
func (l *Log) Alert(message string) {
	l.log(message, "Alert")
}

// Critical - Critical conditions.
func (l *Log) Critical(message string) {
	l.log(message, "Critical")
}

// Error - Runtime errors that do not require immediate action but should typically be logged and monitored
func (l *Log) Error(message string) {
	l.log(message, "Error")
}

// Warning - Exceptional occurrences that are not errors.
func (l *Log) Warning(message string) {
	l.log(message, "Warning")
}

// Notice - Normal but significant events.
func (l *Log) Notice(message string) {
	l.log(message, "Notice")
}

// Info - Interesting events.
func (l *Log) Info(message string) {
	l.log(message, "Info")
}

// Debug - Detailed debug information.
func (l *Log) Debug(message string) {
	l.log(message, "Debug")
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func getBaseLogMessage() logMessage {
	host, _ := os.Hostname()
	return logMessage{Timestamp: makeTimestamp(), Host: host}
}

// MockTransport is a mock version of Logger
type MockTransport struct {
	Logs map[int]logMessage
}

// NewMockTransport returns a new MockTransport
func NewMockTransport() *MockTransport {
	return &MockTransport{
		Logs: make(map[int]logMessage),
	}
}

// Ship implements Transport.
func (mck *MockTransport) Ship(lm logMessage) error {
	mck.Logs[len(mck.Logs)+1] = lm
	return nil
}

// NewMockLogger - Create new mock instance of logger
func NewMockLogger() *Log {
	return NewLog(NewMockTransport())
}

// ElasticsearchTransport ships logs to ElasticSearch
type ElasticsearchTransport struct {
	Client *elastic.Client
	Index  string
}

// NewElasticSearchTransport returns a pointer to an ElasticsearchTransport.
// TODO: Inject a client here, rather than instantiating one.
func NewElasticSearchTransport(host string, index string) *ElasticsearchTransport {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(host),
	)
	if err != nil {
		panic(err)
	}
	return &ElasticsearchTransport{Client: client, Index: index}
}

// Ship ships a logMessage to ElasticSearch.
func (esl *ElasticsearchTransport) Ship(lm logMessage) error {
	_, err := esl.Client.Index().Index(esl.Index).Type(lm.Level).BodyJson(lm).Refresh(true).Do()
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

// WriterTransport allows an io.Writer to be used as a transport.
type WriterTransport struct {
	W io.Writer
}

// Ship implements Transport.
func (t WriterTransport) Ship(lm logMessage) error {
	t.W.Write(([]byte)(lm.String() + "\n"))
	return nil
}
