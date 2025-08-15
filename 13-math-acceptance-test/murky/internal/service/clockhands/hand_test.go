package clockhands

import (
	"math"
	"testing"
	"time"

	"github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/murky/internal/testhelpers"
)

func TestHandInRadius(t *testing.T) {
	t.Run("each run should have 60 second", func(t *testing.T) {

		testCases := []struct {
			now    time.Time
			radius float64
		}{
			{
				now:    testhelpers.SetTime(t, 1, 0, 0),
				radius: 0,
			},

			{
				now:    testhelpers.SetTime(t, 6, 4, 15),
				radius: math.Pi / 2,
			},
			{
				now:    testhelpers.SetTime(t, 10, 56, 30),
				radius: math.Pi,
			},

			{
				now:    testhelpers.SetTime(t, 23, 59, 45),
				radius: math.Pi * 1.5,
			},
		}

		for _, testCase := range testCases {
			secondHand := hand{
				length:       1,
				timePerRound: 60 * time.Second,
			}

			got := secondHand.handInRadius(testCase.now)
			want := testCase.radius

			testhelpers.RoughlyEqualFloat64(t, got, want)
		}

	})

	t.Run("each run should have 60 minute", func(t *testing.T) {

		testCases := []struct {
			now    time.Time
			radius float64
		}{
			{
				now:    testhelpers.SetTime(t, 1, 0, 0),
				radius: 0,
			},

			{
				now:    testhelpers.SetTime(t, 6, 15, 0),
				radius: math.Pi / 2,
			},
			{
				now:    testhelpers.SetTime(t, 10, 30, 0),
				radius: math.Pi,
			},

			{
				now:    testhelpers.SetTime(t, 23, 45, 0),
				radius: math.Pi * 1.5,
			},
		}

		for _, testCase := range testCases {
			secondHand := hand{
				length:       1,
				timePerRound: 60 * time.Minute,
			}

			got := secondHand.handInRadius(testCase.now)
			want := testCase.radius

			testhelpers.RoughlyEqualFloat64(t, got, want)
		}

	})

	t.Run("each run should have 12 hours", func(t *testing.T) {

		testCases := []struct {
			now    time.Time
			radius float64
		}{
			{
				now:    testhelpers.SetTime(t, 0, 0, 0),
				radius: 0,
			},

			{
				now:    testhelpers.SetTime(t, 3, 0, 0),
				radius: math.Pi / 2,
			},
			{
				now:    testhelpers.SetTime(t, 6, 0, 0),
				radius: math.Pi,
			},

			{
				now:    testhelpers.SetTime(t, 9, 0, 0),
				radius: math.Pi * 1.5,
			},
		}

		for _, testCase := range testCases {
			secondHand := hand{
				length:       1,
				timePerRound: 12 * time.Hour,
			}

			got := secondHand.handInRadius(testCase.now)
			want := testCase.radius

			testhelpers.RoughlyEqualFloat64(t, got, want)
		}

	})
}

func TestDurationSinceStartOfDate(t *testing.T) {
	testCases := []struct {
		now      time.Time
		duration time.Duration
	}{
		{
			now:      testhelpers.SetTime(t, 1, 0, 0),
			duration: 3600 * time.Second,
		},
		{
			now:      testhelpers.SetTime(t, 23, 59, 59),
			duration: 86399 * time.Second,
		},
	}

	for _, testCase := range testCases {
		h := &hand{}
		got := h.durationSinceStartOfDay(testCase.now)
		want := testCase.duration

		testhelpers.AssertEqual(t, got, want)
	}
}

func TestHandPoint(t *testing.T) {
	t.Run("60 second per round", func(t *testing.T) {

		testCases := []struct {
			now   time.Time
			point Point
		}{
			{
				now: testhelpers.SetTime(t, 1, 0, 0),
				point: Point{
					X: 0,
					Y: -10,
				},
			},

			{
				now: testhelpers.SetTime(t, 6, 4, 15),
				point: Point{
					X: 10,
					Y: 0,
				},
			},
			{
				now: testhelpers.SetTime(t, 10, 56, 30),
				point: Point{
					X: 0,
					Y: 10,
				},
			},

			{
				now: testhelpers.SetTime(t, 23, 59, 45),
				point: Point{
					X: -10,
					Y: 0,
				},
			},
		}

		for _, testCase := range testCases {
			secondHand := hand{
				length:       10,
				timePerRound: 60 * time.Second,
			}

			got := secondHand.handPoint(testCase.now)
			want := testCase.point

			testhelpers.RoughlyEqualFloat32(t, got.X, want.X)
			testhelpers.RoughlyEqualFloat32(t, got.Y, want.Y)
		}

	})
}
