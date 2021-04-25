package screens

import (
	"image/color"

	// Import png decoding
	_ "image/png"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/rytrose/soup-the-moon/game/animation"
	"github.com/rytrose/soup-the-moon/game/audio"
	"github.com/rytrose/soup-the-moon/game/fonts"
	"github.com/rytrose/soup-the-moon/game/input"
	"github.com/rytrose/soup-the-moon/game/util"
)

// MenuOption is a string name of an option.
type MenuOption string

// Option strings.
var (
	NewGame     MenuOption = "New Game"
	Leaderboard MenuOption = "Leaderboard"
)

// options are the various menu options.
var options = []MenuOption{
	NewGame,
	Leaderboard,
}

// menuState maintains all state needed for the menu screen.
type menuState struct {
	playingTheme bool
	selected     int
	nextStar     int
	numStars     int
	stars        map[int]*animation.Star
}

// theMenuState is the state of the menu screen.
var theMenuState = &menuState{
	stars: map[int]*animation.Star{},
}

// Star constants
const (
	menuStarMaxSpawnLength = 30
	menuStarMinDuration    = 2
	menuStarMaxDuration    = 10
)

// UpdateMenu updates menu screen state before every frame.
func UpdateMenu() ScreenID {
	// Play theme music
	if !theMenuState.playingTheme {
		audio.ThemePlayer.Rewind()
		audio.ThemePlayer.Play()
		theMenuState.playingTheme = true
	}

	// Scroll down menu
	if input.Down() {
		theMenuState.selected = util.Mod(theMenuState.selected+1, len(options))
	}

	// Scroll up menu
	if input.Up() {
		theMenuState.selected = util.Mod(theMenuState.selected-1, len(options))
	}

	// Select option
	if input.Enter() {
		// Stop theme music before leaving page
		audio.ThemePlayer.Pause()
		theMenuState.playingTheme = false

		selectedOption := options[theMenuState.selected]
		switch selectedOption {
		case NewGame:
			return ScreenInitials
		}
	}

	return ScreenMenu
}

// DrawMenu draws one frame of the menu screen.
func DrawMenu(count uint64, w, h int, screen *ebiten.Image) {
	drawMenuTitle(w, screen)
	drawMenuOptions(w, screen)
	drawMenuStars(w, h, screen)
}

// drawMenuTitle draws the title to the top of the screen.
func drawMenuTitle(w int, screen *ebiten.Image) {
	title1 := "Shoot"
	title2 := "the Moon"

	// Draw the title centered
	title1X := (w - len(title1)*64) / 2
	title2X := (w - len(title2)*64) / 2
	text.Draw(screen, title1, fonts.ArcadeFont64, title1X, 6*16, color.White)
	text.Draw(screen, title2, fonts.ArcadeFont64, title2X, 12*16, color.White)
}

// drawMenuOptions draws the menu options.
func drawMenuOptions(w int, screen *ebiten.Image) {
	startingY := 20 * 16
	startingX := w / 6
	tab := 64

	// Draw all options
	for i, option := range options {
		text.Draw(screen, string(option), fonts.ArcadeFont32, startingX+tab, startingY+i*64, color.White)
	}

	// Draw cursor
	text.Draw(screen, ">", fonts.ArcadeFont32, startingX, startingY+theMenuState.selected*64, color.White)
}

// drawMenuStars draws random starbursts in the background.
func drawMenuStars(w, h int, screen *ebiten.Image) {
	// Spawn a new star if we've waited long enough
	if theMenuState.nextStar == 0 {
		// Set a duration to wait before drawing the next new star
		theMenuState.nextStar = rand.Intn(menuStarMaxSpawnLength) + 1

		// Create the new star
		newStar := animation.NewStar(
			theMenuState.numStars, // Monotonically increasing ID number
			rand.Intn(w),          // Randomly place the star on the x-axis
			rand.Intn(h),          // Randomly place the star on the y-axis
			rand.Intn(menuStarMaxDuration)+menuStarMinDuration, // Vary the speed
			rand.Intn(animation.StarNumFrames),                 // Vary the starting frame
			rand.Intn(2) == 1,                                  // Vary the direction of the animation
			1+rand.Float64()*3,                                 // Set the scaling factor of the star
		)
		theMenuState.numStars++

		// Set the new star to state
		theMenuState.stars[newStar.ID] = newStar
	}

	// Draw stars
	for _, star := range theMenuState.stars {
		done := star.Draw(w, h, screen)
		if done {
			// Remove star when done animating
			delete(theMenuState.stars, star.ID)
		}
	}

	// Decrement the next star counter
	theMenuState.nextStar--
}
