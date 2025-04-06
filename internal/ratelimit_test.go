package internal

import (
	"testing"
	"time"
)

func TestCalculateRateLimitIntervals(t *testing.T) {
	t.Run("calculate the a minute in the future", func(t *testing.T) {
		t.Parallel()
		minuteAhead := time.Now().Add(time.Minute)
		rl := RateLimit{
			Count: 2,
			RemainingCalls: 4,
			Reset: minuteAhead,
		}

		 duration := rl.CalculateRateLimitIntervals()

		 if duration != (30 * time.Second) {
			t.Errorf("the duration is %+v", duration)
		 }
	})

	t.Run("no remaining calls left", func(t *testing.T) {
		t.Parallel()
		minuteAhead := time.Now().Add(60 * time.Second)
		rl := RateLimit{
			Count: 2,
			RemainingCalls: 0,
			Reset: minuteAhead,
		}

		 duration := rl.CalculateRateLimitIntervals()

		 if duration != (60 * time.Second) {
			t.Errorf("the duration is %+v", duration)
		 }
	})
}