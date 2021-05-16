package screens

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/rytrose/soup-the-moon/game/audio"
	"github.com/rytrose/soup-the-moon/game/fonts"
	"github.com/rytrose/soup-the-moon/game/input"
	"github.com/rytrose/soup-the-moon/game/state"
	"github.com/rytrose/soup-the-moon/game/util"
)

// tokens contains all possible initials tokens.
var tokens = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "!", "?", "$", "*"}

// initialsState maintains all state needed for the initials input screen.
type initialsState struct {
	selected     int
	playingTheme bool
}

// theInitialsState is the state of the initials input screen.
var theInitialsState = &initialsState{}

// UpdateInitials updates initials input screen state before every frame.
func UpdateInitials() ScreenID {
	// Play theme music
	if !theInitialsState.playingTheme {
		audio.InitialsThemePlayer.Rewind()
		audio.InitialsThemePlayer.Play()
		theInitialsState.playingTheme = true
	}

	if input.Back() {
		if theInitialsState.selected == 0 {
			// Stop theme music before leaving page
			audio.InitialsThemePlayer.Pause()
			theInitialsState.playingTheme = false

			return ScreenMenu
		}

		// Move cursor back
		theInitialsState.selected--

		return ScreenInitials
	}

	if input.Enter() {
		if theInitialsState.selected == len(state.Global.CurrentInitials)-1 {
			theInitialsState.selected = 0

			// Stop theme music before leaving page
			audio.InitialsThemePlayer.Pause()
			theInitialsState.playingTheme = false

			return ScreenScoring
		}

		// Move cursor forward
		theInitialsState.selected++

		return ScreenInitials
	}

	if input.Up() {
		// Change token
		state.Global.CurrentInitials[theInitialsState.selected] = util.Mod(state.Global.CurrentInitials[theInitialsState.selected]-1, len(tokens))
		state.Save()
	}

	if input.Down() {
		// Change token
		state.Global.CurrentInitials[theInitialsState.selected] = util.Mod(state.Global.CurrentInitials[theInitialsState.selected]+1, len(tokens))
		state.Save()
	}

	return ScreenInitials
}

// DrawInitials draws one frame of the initials input screen.
func DrawInitials(count uint64, w, h int, screen *ebiten.Image) {
	drawInitialsPrompt(w, screen)
	drawInitialsInitials(w, screen)
}

// drawInitialsPrompt draws the input prompt to the top of the screen.
func drawInitialsPrompt(w int, screen *ebiten.Image) {
	prompt1 := "Enter your"
	prompt2 := "initials..."
	prompt1Y := 3 * 32
	prompt2Y := 5 * 32

	// Draw the title centered
	prompt1X := (w - len(prompt1)*32) / 2
	prompt2X := (w - len(prompt2)*32) / 2
	text.Draw(screen, prompt1, fonts.ArcadeFont32, prompt1X, prompt1Y, color.White)
	text.Draw(screen, prompt2, fonts.ArcadeFont32, prompt2X, prompt2Y, color.White)
}

// drawInitialsInitials draws the initials selectors.
func drawInitialsInitials(w int, screen *ebiten.Image) {
	topCursorY := (9 * 32) - 16
	bottomCursorY := (11 * 32) + 16
	initialsY := 10 * 32
	initialsX := []int{
		(3 * (w - 32)) / 8,
		(w - 32) / 2,
		(5 * (w - 32)) / 8,
	}

	// Draw the initials
	for i, initial := range state.Global.CurrentInitials {
		// Draw initial
		text.Draw(screen, tokens[initial], fonts.ArcadeFont32, initialsX[i], initialsY, color.White)

		if i == theInitialsState.selected {
			// Draw the cursors
			text.Draw(screen, "/\\", fonts.ArcadeFont32, initialsX[i]-16, topCursorY, color.White)
			text.Draw(screen, "\\/", fonts.ArcadeFont32, initialsX[i]-16, bottomCursorY, color.White)
		}
	}
}

// getInitials retrieves the string initials from a list of token indices.
func getInitials(indices []int) []string {
	return []string{
		tokens[indices[0]],
		tokens[indices[1]],
		tokens[indices[2]],
	}
}
