package clog

import (
	"sync"
	"sync/atomic"

	"github.com/cowkcowk/clog/internal/severity"
	"github.com/go-logr/logr"
)

type severityValue struct {
	severity.Severity
}

func (s *severityValue) get() severity.Severity {
	return severity.Severity(atomic.LoadInt32((*int32)(&s.Severity)))
}

func (s *severityValue) set(val severity.Severity) {
	atomic.StoreInt32((*int32)(&s.Severity), int32(val))
}

type OutputStats struct {
	lines int64
	bytes int64
}

// Lines returns the number of lines written.
func (s *OutputStats) Lines() int64 {
	return atomic.LoadInt64(&s.lines)
}

// Bytes returns the number of bytes written.
func (s *OutputStats) Bytes() int64 {
	return atomic.LoadInt64(&s.bytes)
}

var Stats struct {
	Info, Warning, Error OutputStats
}

// Level specifies a level of verbosity for V logs. *Level implements
// flag.Value; the -v flag is of type Level and should be modified
// only through the flag.Value interface.
type Level int32

type loggingT struct {
	settings

	flushD *
}

// flushDaemon periodically flushes the log file buffers.
type flushDaemon struct {
	mu       sync.Mutex
	clock    clock.WithTicker
	flush    func()
	stopC    chan struct{}
	stopDone chan struct{}
}

func newFlushDaemon(flush func(), tickClock)

// Verbose is a boolean type that implements Infof (like Printf) etc.
// See the documentation of V for more information.
type Verbose struct {
	enabled bool
	logr    *logr.Logger
}

func newVerbose(level Level, b bool) Verbose {
	if logging
}
