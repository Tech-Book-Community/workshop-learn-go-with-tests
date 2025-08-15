package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	clockapp "github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/murky/internal/app/clock"
)

func main() {
	clockApp := clockapp.NewClock()
	// ebiten.SetWindowSize(500, 500)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Clock")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(clockApp); err != nil {
		log.Fatal(err)
	}
}
