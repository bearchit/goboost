package clock_test

import (
	"testing"
	"time"

	"github.com/bearchit/goboost/clock"
	"github.com/stretchr/testify/assert"
)

func TestClock_Now(t *testing.T) {
	a := time.Now().Round(time.Second)
	b := clock.New().Now().Round(time.Second)
	assert.Equal(t, a, b)
}

func TestMock_Now(t *testing.T) {
	c := clock.NewMock()
	assert.Equal(t, time.Unix(0, 0), c.Now())
}

func TestWithNow(t *testing.T) {
	now := time.Now()
	c := clock.NewMock(
		clock.WithNow(now),
	)
	assert.Equal(t, now, c.Now())
}
