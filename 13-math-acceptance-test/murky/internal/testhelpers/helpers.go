// Package testhelpers has some helper functions
package testhelpers

import (
	"math"
	"reflect"
	"testing"
	"time"

	"github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/murky/internal/config"
)

// SetTime return time.Time that is base on 1970/1/1 with hour minute second
func SetTime(t testing.TB, hour, minute, second int) time.Time {
	t.Helper()

	timezone, _ := time.LoadLocation(config.TimezoneUTC8)
	return time.Date(1970, time.January, 1, hour, minute, second, 0, timezone)
}

// RoughlyEqualFloat32 will pass if the difference between
// got and want is less than 1e-7
func RoughlyEqualFloat32(t testing.TB, got, want float32) {
	t.Helper()

	const threshold = 1e-7

	if math.Abs(float64(got-want)) > threshold {
		t.Errorf("expect \"%v\", got \"%v\"", want, got)
	}
}

// RoughlyEqualFloat64 will pass if the difference between
// got and want is less than 1e-7
func RoughlyEqualFloat64(t testing.TB, got, want float64) {
	t.Helper()

	const threshold = 1e-7

	if math.Abs(got-want) > threshold {
		t.Errorf("expect \"%v\", got \"%v\"", want, got)
	}
}

// AssertEqual uses reflect.DeepEqual to determine equality
func AssertEqual[T any](t testing.TB, got, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expect \"%v\", got \"%v\"", want, got)
	}
}
