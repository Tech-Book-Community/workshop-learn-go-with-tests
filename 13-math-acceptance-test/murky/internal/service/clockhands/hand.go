package clockhands

import (
	"math"
	"time"
)

// hand use second as base
type hand struct {
	length float32
	// how much baseTime will the hand travels one round of clock
	timePerRound time.Duration
}

// handInRadius giving time and will return the radius of hand
func (h *hand) handInRadius(now time.Time) float64 {
	duration := h.durationSinceStartOfDay(now)
	seconds := math.Mod(duration.Seconds(), h.timePerRound.Seconds())
	radius := (2 * math.Pi * time.Second.Seconds() * seconds) / h.timePerRound.Seconds()

	return radius
}

// durationSinceStartOfDay will return how much time has passed since today
func (h *hand) durationSinceStartOfDay(now time.Time) time.Duration {
	beginOfDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return now.Sub(beginOfDate)
}

// HandPoint will return the point relevant to the center of clock
func (h *hand) handPoint(now time.Time) Point {
	radius := h.handInRadius(now)

	point := Point{
		X: float32(math.Sin(radius)),
		Y: float32(math.Cos(radius)),
	}

	// the direction of y in screen is from top to down
	// so we need to reverse y
	point = Point{
		X: point.X,
		Y: -1 * point.Y,
	}

	// then multiply by length
	point = Point{
		X: point.X * h.length,
		Y: point.Y * h.length,
	}
	return point
}
