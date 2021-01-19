package v2

import (
	"fmt"
	"io"
	"sync"

	"github.com/sirupsen/logrus"
)

// init - Intializes "v2" package
func init() {
	loggers = &sync.Map{}
}

// LoggerKeyAlreadyPresentError - Defines error for already present logging key in the loggers map
type LoggerKeyAlreadyPresentError struct {
	msg string
}

// NewLoggerKeyAlreadyPresentError - Returns a new instance of LoggerKeyAlreadyPresentError
func NewLoggerKeyAlreadyPresentError(msg string) *LoggerKeyAlreadyPresentError {
	return &LoggerKeyAlreadyPresentError{msg}
}

// Error - Implements error interface
func (e *LoggerKeyAlreadyPresentError) Error() string {
	return e.msg
}

// LoggerNotFoundError - Defines error for already present logging key in the loggers map
type LoggerNotFoundError struct {
	msg string
}

// NewLoggerNotFoundError - Returns a new instance of LoggerNotFoundError
func NewLoggerNotFoundError(msg string) *LoggerNotFoundError {
	return &LoggerNotFoundError{msg}
}

// Error - Implements error interface
func (e *LoggerNotFoundError) Error() string {
	return e.msg
}

// DefaultLoggerName - default logger key
const DefaultLoggerName = "default"

var (
	loggers *sync.Map
)

// GetLogger - Returns the default logger
func GetLogger() *logrus.Logger {
	log, err := GetLoggerWithKey(DefaultLoggerName)
	if err != nil {
		panic(err)
	}
	return log
}

// GetLoggerWithKey - Returns the specific log by key, error if not found
func GetLoggerWithKey(key string) (*logrus.Logger, *LoggerNotFoundError) {

	log, exists := loggers.Load(key)
	if !exists {
		return nil, NewLoggerNotFoundError(fmt.Sprintf("'%s' logger not found", key))
	}

	return log.(*logrus.Logger), nil
}

// InitLogger - Initializes default logger, with specifics io.Writer (stderr, stdoout, file, etc), a formatter(see logrus.Formatter) and a level
func InitLogger(out io.Writer, formatter logrus.Formatter, level logrus.Level) error {
	return InitLoggerWithKey(DefaultLoggerName, out, formatter, level)
}

// InitLoggerWithKey - Initializes a custom key logger, with specifics io.Writer (stderr, stdoout, file, etc), a formatter(see logrus.Formatter) and a level
func InitLoggerWithKey(key string, out io.Writer, formatter logrus.Formatter, level logrus.Level) error {

	l := logrus.New()

	l.Out = out
	l.Formatter = formatter
	l.Level = level

	if _, exists := loggers.LoadOrStore(key, l); exists {
		NewLoggerKeyAlreadyPresentError(fmt.Sprintf("'%s' already present in logging map", key))
	}

	return nil
}
