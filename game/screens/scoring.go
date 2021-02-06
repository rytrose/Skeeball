package screens

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/rytrose/soup-the-moon/game/fonts"
	"github.com/rytrose/soup-the-moon/game/input"
)

// scoringState maintains all state needed for the scoring screen.
type scoringState struct {
	score          int
	confirmingBack bool
}

// theScoringState is the state of the scoring screen.
var theScoringState = &scoringState{}

// UpdateScoring updates the scoring screen state before every frame.
func UpdateScoring() ScreenID {
	if input.Back() {
		if theScoringState.confirmingBack {
			// Reset score
			theScoringState.score = 0

			// Clear confirmation flag
			theScoringState.confirmingBack = false

			return ScreenMenu
		}

		// Set confirmation flag
		theScoringState.confirmingBack = true

		return ScreenScoring
	}

	if input.Enter() {
		if theScoringState.confirmingBack {
			// Clear confirmation flag
			theScoringState.confirmingBack = false
		}
	}

	// Update score
	if input.Mercury() {
		theScoringState.score += -250
	}
	if input.Earth() {
		theScoringState.score += 250
	}
	if input.Mars() {
		theScoringState.score += 500
	}
	if input.Jupiter() {
		theScoringState.score += 1000
	}
	if input.Saturn() {
		theScoringState.score += 2000
	}
	if input.Pluto() {
		theScoringState.score += 5000
	}

	return ScreenScoring
}

// DrawScoring draws one frame of the scoring screen.
func DrawScoring(count int, w, h int, screen *ebiten.Image) {
	drawPlayerInitials(screen)
	drawScore(w, screen)

	if theScoringState.confirmingBack {
		drawConfirmBack(w, screen)
	}
}

// drawPlayerInitials draws the current player's initials.
func drawPlayerInitials(screen *ebiten.Image) {
	initialsY := 2 * 32
	initialsX := 32

	// Draw the intials space delimited
	initialsString := strings.Join(getInitials(), " ")
	text.Draw(screen, initialsString, fonts.ArcadeFont32, initialsX, initialsY, color.White)
}

// drawScore draws the current player's score.
func drawScore(w int, screen *ebiten.Image) {
	scoreY := 2 * 32
	scoreX := w - (7 * 32) - 32

	// Draw the score
	scoreString := fmt.Sprintf("%07d", theScoringState.score)
	text.Draw(screen, scoreString, fonts.ArcadeFont32, scoreX, scoreY, color.White)
}

// drawConfirmBack draws a confirmation message for exiting.
func drawConfirmBack(w int, screen *ebiten.Image) {
	confirmationStringL1 := "Return to base?"
	confirmationStringL2 := "Press back to quit."
	confirmationStringL3 := "Press enter"
	confirmationStringL4 := "to resume."

	confirmationYL1 := 12 * 16
	confirmationXL1 := (w - len(confirmationStringL1)*32) / 2
	confirmationYL2 := 16 * 16
	confirmationXL2 := (w - len(confirmationStringL2)*32) / 2
	confirmationYL3 := 20 * 16
	confirmationXL3 := (w - len(confirmationStringL3)*32) / 2
	confirmationYL4 := 23 * 16
	confirmationXL4 := (w - len(confirmationStringL4)*32) / 2

	// Draw the confirmation message
	text.Draw(screen, confirmationStringL1, fonts.ArcadeFont32, confirmationXL1, confirmationYL1, color.White)
	text.Draw(screen, confirmationStringL2, fonts.ArcadeFont32, confirmationXL2, confirmationYL2, color.White)
	text.Draw(screen, confirmationStringL3, fonts.ArcadeFont32, confirmationXL3, confirmationYL3, color.White)
	text.Draw(screen, confirmationStringL4, fonts.ArcadeFont32, confirmationXL4, confirmationYL4, color.White)
}
