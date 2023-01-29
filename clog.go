package clog

import (
	"sync/atomic"

	"github.com/cowkcowk/clog/internal/severity"
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