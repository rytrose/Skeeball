package screens

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/rytrose/soup-the-moon/game/fonts"
	"github.com/rytrose/soup-the-moon/game/input"
	"github.com/rytrose/soup-the-moon/game/util"
)

// tokens contains all possible initials tokens.
var tokens = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "!", "?", "$", "*"}

// initialsState maintains all state needed for the initials input screen.
type initialsState struct {
	selected int
	initials []int
}

// theInitialsState is the state of the initials input screen.
var theInitialsState = &initialsState{
	initials: make([]int, 3),
}

// UpdateInitials updates initials input screen state before every frame.
func UpdateInitials() ScreenID {
	if input.Back() {
		if theInitialsState.selected == 0 {
			return ScreenMenu
		}

		// Move cursor back
		theInitialsState.selected--

		return ScreenInitials
	}

	if input.Enter() {
		if theInitialsState.selected == len(theInitialsState.initials)-1 {
			theInitialsState.selected = 0
			return ScreenScoring
		}

		// Move cursor forward
		theInitialsState.selected++

		return ScreenInitials
	}

	if input.Up() {
		// Change token
		theInitialsState.initials[theInitialsState.selected] = util.Mod(theInitialsState.initials[theInitialsState.selected]-1, len(tokens))
	}

	if input.Down() {
		// Change token
		theInitialsState.initials[theInitialsState.selected] = util.Mod(theInitialsState.initials[theInitialsState.selected]+1, len(tokens))
	}

	return ScreenInitials
}

// DrawInitials draws one frame of the initials input screen.
func DrawInitials(count int, w, h int, screen *ebiten.Image) {
	drawPrompt(w, screen)
	drawInitials(w, screen)
}

// drawPrompt draws the input prompt to the top of the screen.
func drawPrompt(w int, screen *ebiten.Image) {
	prompt := "Enter your initials..."
	promptY := 4 * 16

	// Draw the title centered
	promptX := (w - len(prompt)*16) / 2
	text.Draw(screen, prompt, fonts.ArcadeFont16, promptX, promptY, color.White)
}

// drawInitials draws the initials selectors.
func drawInitials(w int, screen *ebiten.Image) {
	topCursorY := 8 * 16
	bottomCursorY := 12 * 16
	initialsY := 10 * 16
	initialsX := []int{
		(3 * (w - 16)) / 8,
		(w - 16) / 2,
		(5 * (w - 16)) / 8,
	}

	// Draw the initials
	for i, initial := range theInitialsState.initials {
		// Draw initial
		text.Draw(screen, tokens[initial], fonts.ArcadeFont16, initialsX[i], initialsY, color.White)

		if i == theInitialsState.selected {
			// Draw the cursors
			text.Draw(screen, "/\\", fonts.ArcadeFont16, initialsX[i]-8, topCursorY, color.White)
			text.Draw(screen, "\\/", fonts.ArcadeFont16, initialsX[i]-8, bottomCursorY, color.White)
		}
	}
}

// getInitials retrieves the initials chosen.
func getInitials() []string {
	return []string{
		tokens[theInitialsState.initials[0]],
		tokens[theInitialsState.initials[1]],
		tokens[theInitialsState.initials[2]],
	}
}
