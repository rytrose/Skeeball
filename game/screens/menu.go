package screens

import (
	"bytes"
	"image"
	"image/color"

	// Import png decoding
	_ "image/png"

	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/rytrose/soup-the-moon/game/fonts"
	"github.com/rytrose/soup-the-moon/game/images"
	"github.com/rytrose/soup-the-moon/game/input"
	"github.com/rytrose/soup-the-moon/game/util"
)

// starburstImage is the loaded startburst sprite PNG.
var starburstImage *ebiten.Image

func init() {
	// Read image from byte slice
	img, _, err := image.Decode(bytes.NewReader(images.Starburst_png))
	if err != nil {
		log.Fatal(err)
	}

	// Populate image
	starburstImage = ebiten.NewImageFromImage(img)
}

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
	starMaxSpawnLength = 30
	starMinDuration    = 2
	starMaxDuration    = 10
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
func (s *Star) Draw(w, h int, screen *ebiten.Image) {
	// Star is done animating
	if s.frame < 0 || s.frame == starNumFrames {
		delete(theMenuState.stars, s.id)
		return
	}

	// Draw star animation frame
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.x), float64(s.y))
	sx, sy := s.frame*starW, 0
	screen.DrawImage(starburstImage.SubImage(image.Rect(sx, sy, sx+starW, sy+starH)).(*ebiten.Image), op)

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
		theMenuState.selected = util.Mod(theMenuState.selected+1, len(options))
	}

	// Scroll up menu
	if input.Up() {
		theMenuState.selected = util.Mod(theMenuState.selected-1, len(options))
	}

	// Select option
	if input.Enter() {
		selectedOption := options[theMenuState.selected]
		switch selectedOption {
		case NewGame:
			return ScreenInitials
		}
	}

	return ScreenMenu
}

// DrawMenu draws one frame of the menu screen.
func DrawMenu(count int, w, h int, screen *ebiten.Image) {
	drawTitle(w, screen)
	drawOptions(w, screen)
	drawStars(count, w, h, screen)
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
		text.Draw(screen, string(option), fonts.ArcadeFont16, startingX+tab, startingY+i*32, color.White)
	}

	// Draw cursor
	text.Draw(screen, ">", fonts.ArcadeFont16, startingX, startingY+theMenuState.selected*32, color.White)
}

// drawStars draws random starbursts in the background.
func drawStars(c int, w, h int, screen *ebiten.Image) {
	// Spawn a new star if we've waited long enough
	if theMenuState.nextStar == 0 {
		// Set a duration to wait before drawing the next new star
		theMenuState.nextStar = rand.Intn(starMaxSpawnLength) + 1

		// Create the new star
		newStar := &Star{
			id:        theMenuState.numStars,
			x:         rand.Intn(w),
			y:         rand.Intn(h),
			speed:     rand.Intn(starMaxDuration) + starMinDuration,
			frame:     rand.Intn(starNumFrames),
			direction: rand.Intn(2) == 1,
		}
		theMenuState.numStars++

		// Set the new star to state
		theMenuState.stars[newStar.id] = newStar
	}

	// Draw stars
	for _, star := range theMenuState.stars {
		star.Draw(w, h, screen)
	}

	// Decrement the next star counter
	theMenuState.nextStar--
}
