package baselib

import (
	"fmt"
	"github.com/jcelliott/lumber"
	"os"
	"runtime"
	"strings"
)

// log level
//TRACE = 0
//DEBUG
//INFO
//WARN
//ERROR
//FATAL

type Logger struct {
	logfile *lumber.FileLogger
}

func NewLogger(filename string, level int) (*Logger, error) {
	logger := new(Logger)
	var err error
	logger.logfile, err = lumber.NewFileLogger(filename, level, lumber.ROTATE, 200000, 9, lumber.BUFSIZE)
	if err != nil {
		return nil, err
	}
	logger.logfile.TimeFormat(fmt.Sprintf("[%s]", lumber.TIMEFORMAT))

	return logger, nil
}

func (l *Logger) Debug(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	index := strings.LastIndex(file, string(os.PathSeparator)) + 1
	l.logfile.Debug(fmt.Sprintf("[%s:%d] %s", file[index:], line, format), v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	index := strings.LastIndex(file, string(os.PathSeparator)) + 1
	l.logfile.Info(fmt.Sprintf("[%s:%d] %s", file[index:], line, format), v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	index := strings.LastIndex(file, string(os.PathSeparator)) + 1
	l.logfile.Warn(fmt.Sprintf("[%s:%d] %s", file[index:], line, format), v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	index := strings.LastIndex(file, string(os.PathSeparator)) + 1
	l.logfile.Error(fmt.Sprintf("[%s:%d] %s", file[index:], line, format), v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	index := strings.LastIndex(file, string(os.PathSeparator)) + 1
	l.logfile.Fatal(fmt.Sprintf("[%s:%d] %s", file[index:], line, format), v...)
}
