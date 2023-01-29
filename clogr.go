package clog

import "github.com/go-logr/logr"

func NewClogr() Logger {
	return New(&clogger)
}

type clogger struct {
	level     int
	callDepth int
	prefix    string
	values    []interface{}
}

func (l *clogger) Enabled() bool {
	return true
}

func (l *clogger) Error(err error, msg string, keysAndValues ...interface{}) {

}

func (l clogger) GetSink() logr.LogSink {

}

func (l *clogger) Info(msg string, keysAndValues ...interface{}) {

}

func (l clogger) V(level int) logr.Logger {

}

func (l clogger) WithCallDepth(depth int) logr.Logger {
	return &l
}

func (l clogger) WithName(name string) logr.Logger {
	return &l
}

func (l clogger) WithValues(keysAndValues ...interface{}) logr.Logger {
	return &l
}

func (l clogger) WithCallStackHelper() (func(), logr.Logger) {

}
