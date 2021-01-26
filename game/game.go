package game

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rytrose/soup-the-moon/game/input"
	"github.com/rytrose/soup-the-moon/game/screens"
	"github.com/rytrose/soup-the-moon/game/util"
)

// Game implements ebiten.Game and maintains state about the game.
type Game struct {
	w      int              // Screen size width.
	h      int              // Screen size height.
	c      int              // Frame counter
	screen screens.ScreenID // An enumeration of the current screen being displayed.
}

// Run starts the game.
func Run() {
	// Seed RNG
	rand.Seed(time.Now().UnixNano())

	// Set up the game window
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Soup The Moon")

	if util.IsRasPi() {
		// TODO: Setup RPIO button input
	}

	// Run game
	if err := ebiten.RunGame(newGame(640, 480)); err != nil {
		log.Fatal(err)
	}
}

// newGame is a Game factory.
func newGame(width, height int) *Game {
	return &Game{
		w: width,
		h: height,
	}
}

// Update updates the game state.
func (g *Game) Update() error {
	// Increment frame counter
	g.c++

	// Update button states
	if util.IsRasPi() {
		input.RPIOButtonUpdate()
	}

	// Screen state machine
	var nextScreen screens.ScreenID
	switch g.screen {
	case screens.ScreenMenu:
		nextScreen = screens.UpdateMenu()
	case screens.ScreenInitials:
		nextScreen = screens.UpdateInitials()
	}

	// Set the next screen
	g.screen = nextScreen

	return nil
}

// Draw draws a frame.
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))

	// Draw appropriate screen
	switch g.screen {
	case screens.ScreenMenu:
		screens.DrawMenu(g.c, g.w, g.h, screen)
	case screens.ScreenInitials:
		screens.DrawInitials(g.c, g.w, g.h, screen)
	}
}

// Layout determines the game's layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.w, g.h
}
