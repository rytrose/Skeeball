package animation

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rytrose/soup-the-moon/game/images"
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

// Star constants
const (
	StarNumFrames = 9
	StarW         = 32
	StarH         = 32
)

// Star maintains state for a single star.
type Star struct {
	ID        int
	x         int
	y         int
	speed     int
	speedCtr  int
	frame     int
	direction bool
}

// NewStar is a Star factory.
func NewStar(id int, x int, y int, speed int, frame int, direction bool) *Star {
	return &Star{
		ID:        id,
		x:         x,
		y:         y,
		speed:     speed,
		frame:     frame,
		direction: direction,
	}
}

// Draw draws a star to the screen.
func (s *Star) Draw(w, h int, screen *ebiten.Image) bool {
	// Star is done animating
	if s.frame < 0 || s.frame == StarNumFrames {
		// No more frames to draw
		return true
	}

	// Draw star animation frame
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.x), float64(s.y))
	sx, sy := s.frame*StarW, 0
	screen.DrawImage(starburstImage.SubImage(image.Rect(sx, sy, sx+StarW, sy+StarH)).(*ebiten.Image), op)

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

	// Another frame to draw
	return false
}
