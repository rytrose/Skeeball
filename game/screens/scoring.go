package screens

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/rytrose/soup-the-moon/game/animation"
	"github.com/rytrose/soup-the-moon/game/fonts"
	"github.com/rytrose/soup-the-moon/game/images"
	"github.com/rytrose/soup-the-moon/game/input"
)

// planetImageWidth is the width of the planet PNGs.
const planetImageWidth = 280

// mercuryImage is the loaded mercury PNG.
var mercuryImage *ebiten.Image

// earthImage is the loaded earth PNG.
var earthImage *ebiten.Image

// marsImage is the loaded mars PNG.
var marsImage *ebiten.Image

// jupiterImage is the loaded jupiter PNG.
var jupiterImage *ebiten.Image

// saturnImage is the loaded saturn PNG.
var saturnImage *ebiten.Image

// plutoImage is the loaded pluto PNG.
var plutoImage *ebiten.Image

func init() {
	mercuryImage = loadImage(images.Mercury_png)
	earthImage = loadImage(images.Earth_png)
	marsImage = loadImage(images.Mars_png)
	jupiterImage = loadImage(images.Jupiter_png)
	saturnImage = loadImage(images.Saturn_png)
	plutoImage = loadImage(images.Pluto_png)
}

// loadImage is a helper function to load an image.
func loadImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

// scoringState maintains all state needed for the scoring screen.
type scoringState struct {
	score          int
	confirmingBack bool
	planetImage    *ebiten.Image
	animatedText   *animation.TextScale
	nextStar       int
	numStars       int
	stars          map[int]*animation.Star
}

// theScoringState is the state of the scoring screen.
var theScoringState = &scoringState{
	stars: map[int]*animation.Star{},
}

// Star constants
const (
	scoringStarMaxSpawnLength = 60
	scoringStarMinDuration    = 2
	scoringStarMaxDuration    = 8
)

// UpdateScoring updates the scoring screen state before every frame.
func UpdateScoring(w, h int) ScreenID {
	if input.Back() {
		if theScoringState.confirmingBack {
			// Reset score
			theScoringState.score = 0

			// Stop animating planet
			theScoringState.planetImage = nil
			theScoringState.animatedText = nil

			// Clear confirmation flag
			theScoringState.confirmingBack = false

			return ScreenMenu
		}

		// Stop animating planet
		theScoringState.planetImage = nil
		theScoringState.animatedText = nil

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
		theScoringState.planetImage = mercuryImage
		theScoringState.animatedText = animation.NewTextScale(0, w/2, 3*h/4, "MERCURY", fonts.ArcadeFont32, color.White, 0.05, 5, 24, 1)
	}
	if input.Earth() {
		theScoringState.score += 250
		theScoringState.planetImage = earthImage
		theScoringState.animatedText = animation.NewTextScale(0, w/2, 3*h/4, "EARTH", fonts.ArcadeFont32, color.White, 1.5, 2, 20, 3)
	}
	if input.Mars() {
		theScoringState.score += 500
		theScoringState.planetImage = marsImage
		theScoringState.animatedText = animation.NewTextScale(0, w/2, 3*h/4, "MARS", fonts.ArcadeFont32, color.White, 2.0, 3, 10, 4)
	}
	if input.Jupiter() {
		theScoringState.score += 1000
		theScoringState.planetImage = jupiterImage
		theScoringState.animatedText = animation.NewTextScale(0, w/2, 3*h/4, "JUPITER", fonts.ArcadeFont32, color.White, 2.5, 4, 10, 3)
	}
	if input.Saturn() {
		theScoringState.score += 2000
		theScoringState.planetImage = saturnImage
		theScoringState.animatedText = animation.NewTextScale(0, w/2, 3*h/4, "SATURN", fonts.ArcadeFont64, color.White, 1.6, 3, 8, 5)
	}
	if input.Pluto() {
		theScoringState.score += 5000
		theScoringState.planetImage = plutoImage
		theScoringState.animatedText = animation.NewTextScale(0, w/2, 3*h/4, "PLUTO", fonts.ArcadeFont16, color.White, 12.0, 5, 8, 6)
	}

	return ScreenScoring
}

// DrawScoring draws one frame of the scoring screen.
func DrawScoring(count uint64, w, h int, screen *ebiten.Image) {
	drawScoringStars(w, h, screen)
	drawScoringPlayerInitials(screen)
	drawScoringScore(w, screen)
	drawScoringPlanetAnimation(w, h, screen)

	if theScoringState.confirmingBack {
		drawScoringConfirmBack(w, screen)
	}
}

// drawScoringPlayerInitials draws the current player's initials.
func drawScoringPlayerInitials(screen *ebiten.Image) {
	initialsY := 2 * 32
	initialsX := 32

	// Draw the intials space delimited
	initialsString := strings.Join(getInitials(), " ")
	text.Draw(screen, initialsString, fonts.ArcadeFont32, initialsX, initialsY, color.White)
}

// drawScoringScore draws the current player's score.
func drawScoringScore(w int, screen *ebiten.Image) {
	scoreY := 2 * 32
	scoreX := w - (7 * 32) - 32

	// Draw the score
	scoreString := fmt.Sprintf("%07d", theScoringState.score)
	text.Draw(screen, scoreString, fonts.ArcadeFont32, scoreX, scoreY, color.White)
}

// drawScoringConfirmBack draws a confirmation message for exiting.
func drawScoringConfirmBack(w int, screen *ebiten.Image) {
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

// drawScoringPlanetAnimation draws a planet image a text animation, if necessary.
func drawScoringPlanetAnimation(w, h int, screen *ebiten.Image) {
	if theScoringState.planetImage != nil && theScoringState.animatedText != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(w)/2-float64(planetImageWidth)/2, float64(h)/5)
		screen.DrawImage(theScoringState.planetImage, op)
		done := theScoringState.animatedText.Draw(w, h, screen)
		if done {
			theScoringState.planetImage = nil
			theScoringState.animatedText = nil
		}
	}
}

// drawScoringStars draws random starbursts in the background.
func drawScoringStars(w, h int, screen *ebiten.Image) {
	// Spawn a new star if we've waited long enough
	if theScoringState.nextStar == 0 {
		// Set a duration to wait before drawing the next new star
		theScoringState.nextStar = rand.Intn(scoringStarMaxSpawnLength) + 1

		// Create the new star
		newStar := animation.NewStar(
			theScoringState.numStars, // Monotonically increasing ID number
			rand.Intn(w),             // Randomly place the star on the x-axis
			rand.Intn(h),             // Randomly place the star on the y-axis
			rand.Intn(scoringStarMaxDuration)+scoringStarMinDuration, // Vary the speed
			rand.Intn(animation.StarNumFrames),                       // Vary the starting frame
			rand.Intn(2) == 1,                                        // Vary the direction of the animation
			1+rand.Float64()*2,                                       // Set the scaling factor of the star
		)
		theScoringState.numStars++

		// Set the new star to state
		theScoringState.stars[newStar.ID] = newStar
	}

	// Draw stars
	for _, star := range theScoringState.stars {
		done := star.Draw(w, h, screen)
		if done {
			// Remove star when done animating
			delete(theScoringState.stars, star.ID)
		}
	}

	// Decrement the next star counter
	theScoringState.nextStar--
}
