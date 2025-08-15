// Package clock will create an clock image
package clock

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/murky/internal/config"
	"github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/murky/internal/service/clockhands"
)

// Draw create clock image
func Draw(now time.Time) *ebiten.Image {
	canvas := ebiten.NewImage(2*config.ClockRadius, 2*config.ClockRadius)

	center := clockhands.Point{
		X: config.ClockRadius,
		Y: config.ClockRadius,
	}

	// draw clock surface
	vector.DrawFilledCircle(
		canvas,
		center.X,
		center.Y,
		config.ClockRadius,
		color.White,
		true,
	)

	vector.StrokeCircle(
		canvas,
		center.X,
		center.Y,
		config.ClockRadius,
		4,
		color.Black, true,
	)

	// draw secondHand
	secondHandPoint := clockhands.SecondHand(now)
	vector.StrokeLine(
		canvas,
		center.X,
		center.Y,
		secondHandPoint.X,
		secondHandPoint.Y,
		2,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
		true,
	)

	// draw minute hand
	minuteHandPoint := clockhands.MinuteHand(now)
	vector.StrokeLine(
		canvas,
		center.X,
		center.Y,
		minuteHandPoint.X,
		minuteHandPoint.Y,
		3,
		color.Black,
		true,
	)

	// draw hour hand
	hourHandPoint := clockhands.HourHand(now)
	vector.StrokeLine(
		canvas,
		center.X,
		center.Y,
		hourHandPoint.X,
		hourHandPoint.Y,
		4,
		color.Black,
		true,
	)
	return canvas
}
