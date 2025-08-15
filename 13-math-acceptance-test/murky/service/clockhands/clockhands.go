package clockhands

import (
	"time"

	"github.com/TinyMurky/go-clock/internal/config"
)

// SecondHand will get the Point of the tail of Second hand at current time
func SecondHand(now time.Time) Point {
	secondHand := hand{
		length:       config.SecondHandLength,
		timePerRound: time.Minute,
	}

	releventPoint := secondHand.handPoint(now)
	return Point{
		X: config.ClockRadius + releventPoint.X,
		Y: config.ClockRadius + releventPoint.Y,
	}
}

// MinuteHand will get the Point of the tail of Minute hand at current time
func MinuteHand(now time.Time) Point {
	minuteHand := hand{
		length:       config.MinuteHandLength,
		timePerRound: time.Hour,
	}

	releventPoint := minuteHand.handPoint(now)
	return Point{
		X: config.ClockRadius + releventPoint.X,
		Y: config.ClockRadius + releventPoint.Y,
	}
}

// HourHand will get the Point of the tail of Minute hand at current time
func HourHand(now time.Time) Point {
	hourHand := hand{
		length:       config.HourHandLength,
		timePerRound: 12 * time.Hour,
	}

	releventPoint := hourHand.handPoint(now)
	return Point{
		X: config.ClockRadius + releventPoint.X,
		Y: config.ClockRadius + releventPoint.Y,
	}
}
