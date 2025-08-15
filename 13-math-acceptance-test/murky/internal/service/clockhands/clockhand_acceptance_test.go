package clockhands_test

import (
	"testing"
	"time"

	"github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/murky/internal/config"
	"github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/murky/internal/service/clockhands"
	"github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/murky/internal/testhelpers"
)

func TestSecondHand(t *testing.T) {
	testCases := []struct {
		now   time.Time
		point clockhands.Point
	}{
		{
			testhelpers.SetTime(t, 0, 0, 0), // up
			clockhands.Point{
				X: config.ClockRadius,
				Y: config.ClockRadius - config.SecondHandLength,
			},
		},
		{
			testhelpers.SetTime(t, 0, 0, 15), // right
			clockhands.Point{
				X: config.ClockRadius + config.SecondHandLength,
				Y: config.ClockRadius,
			},
		},
		{
			testhelpers.SetTime(t, 0, 0, 30), // down
			clockhands.Point{
				X: config.ClockRadius,
				Y: config.ClockRadius + config.SecondHandLength,
			},
		},
		{
			testhelpers.SetTime(t, 0, 0, 45), // left
			clockhands.Point{
				X: config.ClockRadius - config.SecondHandLength,
				Y: config.ClockRadius,
			},
		},
	}

	for _, testCase := range testCases {
		got := clockhands.SecondHand(testCase.now)

		testhelpers.RoughlyEqualFloat32(t, got.X, testCase.point.X)
		testhelpers.RoughlyEqualFloat32(t, got.Y, testCase.point.Y)

	}
}

func TestMinuteHand(t *testing.T) {
	testCases := []struct {
		now   time.Time
		point clockhands.Point
	}{
		{
			testhelpers.SetTime(t, 0, 0, 0), // up
			clockhands.Point{
				X: config.ClockRadius,
				Y: config.ClockRadius - config.MinuteHandLength,
			},
		},
		{
			testhelpers.SetTime(t, 0, 15, 0), // right
			clockhands.Point{
				X: config.ClockRadius + config.MinuteHandLength,
				Y: config.ClockRadius,
			},
		},
		{
			testhelpers.SetTime(t, 0, 30, 0), // down
			clockhands.Point{
				X: config.ClockRadius,
				Y: config.ClockRadius + config.MinuteHandLength,
			},
		},
		{
			testhelpers.SetTime(t, 0, 45, 0), // left
			clockhands.Point{
				X: config.ClockRadius - config.MinuteHandLength,
				Y: config.ClockRadius,
			},
		},
	}

	for _, testCase := range testCases {
		got := clockhands.MinuteHand(testCase.now)

		testhelpers.RoughlyEqualFloat32(t, got.X, testCase.point.X)
		testhelpers.RoughlyEqualFloat32(t, got.Y, testCase.point.Y)

	}
}

func TestHourHand(t *testing.T) {
	testCases := []struct {
		now   time.Time
		point clockhands.Point
	}{
		{
			testhelpers.SetTime(t, 0, 0, 0), // up
			clockhands.Point{
				X: config.ClockRadius,
				Y: config.ClockRadius - config.HourHandLength,
			},
		},
		{
			testhelpers.SetTime(t, 3, 0, 0), // right
			clockhands.Point{
				X: config.ClockRadius + config.HourHandLength,
				Y: config.ClockRadius,
			},
		},
		{
			testhelpers.SetTime(t, 6, 0, 0), // down
			clockhands.Point{
				X: config.ClockRadius,
				Y: config.ClockRadius + config.HourHandLength,
			},
		},
		{
			testhelpers.SetTime(t, 9, 0, 0), // left
			clockhands.Point{
				X: config.ClockRadius - config.HourHandLength,
				Y: config.ClockRadius,
			},
		},
	}

	for _, testCase := range testCases {
		got := clockhands.HourHand(testCase.now)

		testhelpers.RoughlyEqualFloat32(t, got.X, testCase.point.X)
		testhelpers.RoughlyEqualFloat32(t, got.Y, testCase.point.Y)

	}
}
