package log

import (
	"log"
	"os"
)

var _logger *logger

type logger struct {
	*log.Logger
}

func init() {
	_logger = &logger{Logger: log.New(os.Stderr, "[ip2fsd] ", 0)}
}

func (l *logger) Fatal(v ...interface{}) {
	l.Logger.Fatal(v...)
}

func (l *logger) Panic(v ...interface{}) {
	l.Logger.Panic(v...)
}

func (l *logger) Debug(v ...interface{}) {
	l.Logger.Print(v...)
}

func (l *logger) Debugln(v ...interface{}) {
	l.Logger.Println(v...)
}

func (l *logger) Debugf(fmt string, v ...interface{}) {
	l.Logger.Printf(fmt, v...)
}

func Fatal(v ...interface{}) {
	_logger.Fatal(v...)
}

func Panic(v ...interface{}) {
	_logger.Panic(v...)
}

func Debug(v ...interface{}) {
	_logger.Debug(v...)
}

func Debugln(v ...interface{}) {
	_logger.Debugln(v...)
}

func Debugf(fmt string, v ...interface{}) {
	_logger.Debugf(fmt, v...)
}
