package clockhands

import (
	"time"

	"github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/murky/internal/config"
)

// SecondHand will get the Point of the tail of Second hand at current time
func SecondHand(now time.Time) Point {
	secondHand := hand{
		length:       config.SecondHandLength,
		timePerRound: time.Minute,
	}

	relevantPoint := secondHand.handPoint(now)
	return Point{
		X: config.ClockRadius + relevantPoint.X,
		Y: config.ClockRadius + relevantPoint.Y,
	}
}

// MinuteHand will get the Point of the tail of Minute hand at current time
func MinuteHand(now time.Time) Point {
	minuteHand := hand{
		length:       config.MinuteHandLength,
		timePerRound: time.Hour,
	}

	relevantPoint := minuteHand.handPoint(now)
	return Point{
		X: config.ClockRadius + relevantPoint.X,
		Y: config.ClockRadius + relevantPoint.Y,
	}
}

// HourHand will get the Point of the tail of Hour hand at current time
func HourHand(now time.Time) Point {
	hourHand := hand{
		length:       config.HourHandLength,
		timePerRound: 12 * time.Hour,
	}

	relevantPoint := hourHand.handPoint(now)
	return Point{
		X: config.ClockRadius + relevantPoint.X,
		Y: config.ClockRadius + relevantPoint.Y,
	}
}
