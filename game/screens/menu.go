package screens

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/rytrose/soup-the-moon/game/fonts"
	"github.com/rytrose/soup-the-moon/game/input"
)

// options are the various menu options.
var options = []string{
	"New Game",
	"Leaderboard",
}

// menuState maintains all state needed for the menu screen.
type menuState struct {
	selected int
	nextStar int
	numStars int
	stars    map[int]*Star
}

// theMenuState is the state of the menu screen.
var theMenuState = &menuState{
	stars: map[int]*Star{},
}

// Star constants
const (
	starMaxSpawnLength = 60
	starMaxDuration    = 5
	starNumFrames      = 9
	starW              = 32
	starH              = 32
)

// Star maintains state for a starburst animation.
type Star struct {
	id        int
	x         int
	y         int
	speed     int
	speedCtr  int
	frame     int
	direction bool
}

// Draw draws a star to the screen.
func (s *Star) Draw(screen *ebiten.Image) {
	// Star is done animating
	if s.frame < 0 || s.frame == starNumFrames {
		delete(theMenuState.stars, s.id)
		return
	}

	// Draw star animation frame
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.x), float64(s.y))

	// Progress the animation after the frame has been shown enough
	if s.speedCtr == s.speed {
		if s.direction {
			s.frame++
		} else {
			s.frame--
		}
		s.speedCtr = 0
	}

	// Increment speed counter
	s.speedCtr++
}

// UpdateMenu updates menu screen state before every frame.
func UpdateMenu() ScreenID {
	// Scroll down menu
	if input.Down() {
		next := theMenuState.selected + 1
		if next == len(options) {
			next = 0
		}
		theMenuState.selected = next
	}

	// Scroll up menu
	if input.Up() {
		next := theMenuState.selected - 1
		if next == -1 {
			next = len(options) - 1
		}
		theMenuState.selected = next
	}

	return ScreenMenu
}

// DrawMenu draws one frame of the menu screen.
func DrawMenu(count int, w, h int, screen *ebiten.Image) {
	drawTitle(w, screen)
	drawOptions(w, screen)
}

// drawTitle draws the title to the top of the screen.
func drawTitle(w int, screen *ebiten.Image) {
	title := "Shoot the Moon"

	// Draw the title centered
	titleX := (w - len(title)*32) / 2
	text.Draw(screen, title, fonts.ArcadeFont32, titleX, 4*16, color.White)
}

// drawOptions draws the menu options.
func drawOptions(w int, screen *ebiten.Image) {
	startingY := 8 * 16
	startingX := w / 4
	tab := 32

	// Draw all options
	for i, option := range options {
		text.Draw(screen, option, fonts.ArcadeFont16, startingX+tab, startingY+i*32, color.White)
	}

	// Draw cursor
	text.Draw(screen, ">", fonts.ArcadeFont16, startingX, startingY+theMenuState.selected*32, color.White)
}

// drawStars draws random starbursts in the background.
func drawStars(c int, w, h int, screen *ebiten.Image) {
	// Spawn a new star if it's time
	if theMenuState.nextStar == 0 {
		theMenuState.nextStar = rand.Intn(starMaxSpawnLength) + 1

		newStar := &Star{
			id:        theMenuState.numStars,
			x:         rand.Intn(w),
			y:         rand.Intn(h),
			speed:     rand.Intn(starMaxDuration) + 1,
			frame:     rand.Intn(starNumFrames),
			direction: rand.Intn(2) == 1,
		}
		theMenuState.numStars++

		theMenuState.stars[newStar.id] = newStar
	}

	// Draw stars
	for _, star := range theMenuState.stars {
		star.Draw(screen)
	}

	theMenuState.nextStar--
}
