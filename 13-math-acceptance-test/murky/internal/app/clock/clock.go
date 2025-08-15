package clockapp

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/murky/internal/config"

	"github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/murky/internal/service/clock"
)

// ClockApp implements ebiten.Game interface.
type ClockApp struct {
	now time.Time
}

var _ ebiten.Game = (*ClockApp)(nil)

// NewClock create a new clock app
func NewClock() *ClockApp {
	return &ClockApp{
		now: time.Now(),
	}
}

// Update proceeds the Clock state.
// Update is called every tick (1/60 [s] by default).
func (c *ClockApp) Update() error {
	// Write your game's logical update.
	c.now = time.Now()
	return nil
}

// Draw draws the Clock screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (c *ClockApp) Draw(screen *ebiten.Image) {
	clockImg := clock.Draw(c.now)
	screen.DrawImage(clockImg, nil)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (c *ClockApp) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
