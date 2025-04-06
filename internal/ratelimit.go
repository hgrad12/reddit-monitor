package internal

import "time"

// max calls per mintue (in seconds) / total calls needed per call = seconds between e3ach call

type RateLimit struct {
	Count int
	RemainingCalls int
	Reset time.Time
}

func (rl *RateLimit) CalculateRateLimitIntervals() time.Duration {
	if rl.RemainingCalls < rl.Count {
		return time.Duration(time.Until(rl.Reset))
	}
	now := time.Now()

	duration := rl.Reset.Sub(now).Seconds()

	interations := rl.RemainingCalls / rl.Count

	sum := duration / float64(interations)

	return time.Duration(sum * float64(time.Second))
}
