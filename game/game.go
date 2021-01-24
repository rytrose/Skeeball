package game

import (
	"log"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rytrose/soup-the-moon/game/input"
)

// Screen defines the screen to be displayed.
type Screen int

// Enumeration of screens.
const (
	ScreenMenu Screen = iota
)

// Game implements ebiten.Game and maintains state about the game.
type Game struct {
	w       int    // Screen size width.
	h       int    // Screen size height.
	screen  Screen // An enumeration of the current screen being displayed.
	isRasPi bool   // A flag determining whether the game is being run on a raspberry pi.
}

// Run starts the game.
func Run() {
	// Set up the game window
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Soup The Moon")

	// Check if on raspberry pi
	log.Printf("GOOS: %s\n", runtime.GOOS)
	isRasPi := false

	if runtime.GOOS == "linux" {
		isRasPi = true

		// TODO: Setup RPIO button input
	}

	// Run game
	if err := ebiten.RunGame(newGame(640, 480, isRasPi)); err != nil {
		log.Fatal(err)
	}
}

// newGame is a Game factory.
func newGame(width, height int, isRasPi bool) *Game {
	return &Game{
		w:       width,
		h:       height,
		isRasPi: isRasPi,
	}
}

// Update updates the game state.
func (g *Game) Update() error {
	input.RPIOButtonUpdate()
	return nil
}

// Draw draws a frame.
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout determines the game's layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.w, g.h
}
