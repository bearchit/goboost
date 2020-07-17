package nullable_test

import (
	"testing"
	"time"

	"github.com/bearchit/goboost/nullable"
	"github.com/stretchr/testify/assert"
)

func TestPtrToTime(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		t1 := nullable.PtrToTime(nil)
		assert.True(t, t1.IsZero())
	})

	t.Run("object", func(t *testing.T) {
		now := time.Date(2020, time.July, 17, 0, 0, 0, 0, time.Local)
		t1 := nullable.PtrToTime(&now)
		assert.Equal(t, now, t1)
	})
}

func TestTimeToPtr(t *testing.T) {
	now := time.Time{}
	t1 := nullable.TimeToPtr(now)
	assert.Equal(t, &now, t1)
}
