package clock

import "time"

type Clock interface {
	Now() time.Time
	SetNow(now time.Time)
}

type clock struct{}

func New() Clock {
	return new(clock)
}

func (c clock) Now() time.Time {
	return time.Now()
}

func (c clock) SetNow(now time.Time) {}

type Mock struct {
	now time.Time
}

func NewMock(options ...func(*Mock)) Clock {
	c := &Mock{
		now: time.Unix(0, 0),
	}
	for _, o := range options {
		o(c)
	}
	return c
}

func WithNow(now time.Time) func(*Mock) {
	return func(mock *Mock) {
		mock.SetNow(now)
	}
}

func (c Mock) Now() time.Time {
	return c.now
}

func (c *Mock) SetNow(now time.Time) {
	c.now = now
}
