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
	score int
}

// theScoringState is the state of the scoring screen.
var theScoringState = &scoringState{}

// UpdateScoring updates the scoring screen state before every frame.
func UpdateScoring() ScreenID {
	if input.Back() {
		// Reset score
		theScoringState.score = 0

		// TODO: consider a back confirmation page
		return ScreenInitials
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
}

// drawPlayerInitials draws the current player's initials.
func drawPlayerInitials(screen *ebiten.Image) {
	initialsY := 4 * 16
	initialsX := 4 * 14

	// Draw the intials space delimited
	initialsString := strings.Join(getInitials(), " ")
	text.Draw(screen, initialsString, fonts.ArcadeFont16, initialsX, initialsY, color.White)
}

// drawScore draws the current player's score.
func drawScore(w int, screen *ebiten.Image) {
	scoreY := 4 * 16
	scoreX := w - (11 * 16)

	// Draw the score
	scoreString := fmt.Sprintf("%07d", theScoringState.score)
	text.Draw(screen, scoreString, fonts.ArcadeFont16, scoreX, scoreY, color.White)
}
