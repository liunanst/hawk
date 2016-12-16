package baselib

import (
	"fmt"
	"github.com/jcelliott/lumber"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	errlog  *lumber.FileLogger
	infolog *lumber.FileLogger
}

func NewLogger(path string, level int) (*Logger, error) {
	logger := new(Logger)
	var err error
	file := fmt.Sprintf("%s-%s", path, "out.log")
	logger.infolog, err = lumber.NewFileLogger(file, level, lumber.ROTATE, 200000, 9, lumber.BUFSIZE)
	if err != nil {
		return nil, err
	}
	logger.infolog.TimeFormat(fmt.Sprintf("[%s]", lumber.TIMEFORMAT))
	file = fmt.Sprintf("%s-%s", path, "err.log")
	logger.errlog, err = lumber.NewFileLogger(file, level, lumber.ROTATE, 100000, 9, lumber.BUFSIZE)
	if err != nil {
		return nil, err
	}
	logger.errlog.TimeFormat(fmt.Sprintf("[%s]", lumber.TIMEFORMAT))
	return logger, nil
}

func (l *Logger) Debug(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	index := strings.LastIndex(file, string(os.PathSeparator)) + 1
	l.infolog.Debug(fmt.Sprintf("[%s:%d] %s", file[index:], line, format), v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	index := strings.LastIndex(file, string(os.PathSeparator)) + 1
	l.infolog.Info(fmt.Sprintf("[%s:%d] %s", file[index:], line, format), v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	index := strings.LastIndex(file, string(os.PathSeparator)) + 1
	l.infolog.Warn(fmt.Sprintf("[%s:%d] %s", file[index:], line, format), v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	index := strings.LastIndex(file, string(os.PathSeparator)) + 1
	l.errlog.Error(fmt.Sprintf("[%s:%d] %s", file[index:], line, format), v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	index := strings.LastIndex(file, string(os.PathSeparator)) + 1
	l.errlog.Fatal(fmt.Sprintf("[%s:%d] %s", file[index:], line, format), v...)
}
