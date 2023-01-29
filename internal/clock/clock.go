package clock

import "time"

// PassiveClock allows for injecting fake or real clocks into code
// that needs to read the current time but does not support scheduling
// activity in the future.
type PassiveClock interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

type Clock interface {
	PassiveClock

	After(d time.Duration) <-chan time.Time
	
	NewTimer(d time.Duration) Timer

	Sleep(d time.Duration)

	Tick(d time.Duration) <-chan time.Time
}

// RealClock really calls time.Now()
type RealClock struct{}

// Now returns the current time.
func (RealClock) Now() time.Time{
	return time.Now()
}

func (RealClock) Since(ts time.Time) time.Duration {
	return time.Since(ts)
}

func (RealClock) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

func (RealClock) NewTimer(d time.Duration) Timer {

}

type Timer interface {
	C() <-chan time.Time
	Stop() bool
	Reset(d time.Duration) bool
}

type realTimer struct {
	timer *time.Timer
}

func (r *realTimer) C()<-chan time.Time {
	return r.timer.C
}

func (r *realTimer) Stop() bool {
	return r.timer.Reset()
}