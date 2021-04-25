package animation

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// TextScale maintains the state of a text scaling animation.
type TextScale struct {
	ID           int
	x            int
	y            int
	text         string
	font         font.Face
	color        color.Color
	scale        float64
	divisions    int
	divisionCtr  int
	speed        int
	speedCtr     int
	iterations   int
	iterationCtr int
}

// NewTextScale is a TextScale factory.
// ID can be used to identify multiple animations.
// x and y are the center point upon which the animation will run.
// font and color specify the text style.
// scale is used to define the change of size of the text.
// divisions is how many equal steps between the original size and the scaled size to animate.
// speed is how many frames each division should take to animate.
func NewTextScale(
	id int,
	x, y int,
	text string,
	font font.Face,
	color color.Color,
	scale float64,
	divisions int,
	speed int,
	iterations int,
) *TextScale {
	return &TextScale{
		ID:         id,
		x:          x,
		y:          y,
		text:       text,
		font:       font,
		color:      color,
		scale:      scale,
		divisions:  divisions,
		speed:      speed,
		iterations: iterations,
	}
}

// Draw draws a TextScale to the screen.
func (t *TextScale) Draw(w, h int, screen *ebiten.Image) bool {
	// This iteration is done animating
	if t.divisionCtr == t.divisions {
		if t.iterationCtr == t.iterations-1 {
			// No more frames to draw
			return true
		}
		t.divisionCtr = 0
		t.iterationCtr++
	}

	// Draw the original sized text to an image
	bounds := text.BoundString(t.font, t.text)
	textWidth, textHeight := int(bounds.Dx()), int(bounds.Dy())
	textImage := ebiten.NewImage(textWidth, textHeight)
	text.Draw(textImage, t.text, t.font, 0-bounds.Min.X, 0-bounds.Min.Y, t.color)

	// Determine scale factor based on division
	scaleFactor := 1 + (((t.scale - 1) / float64(t.divisions)) * float64(t.divisionCtr+1))
	scaledTextWidth := int(scaleFactor * float64(textWidth))
	scaledTextHeight := int(scaleFactor * float64(textHeight))

	// Scale and translate the text image depending on the state of the animation
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleFactor, scaleFactor)
	tx := float64(t.x - scaledTextWidth/2)
	ty := float64(t.y - scaledTextHeight/2)
	op.GeoM.Translate(tx, ty)

	// Draw animation frame
	screen.DrawImage(textImage, op)

	// Progress the animation after the division has been shown enough
	if t.speedCtr == t.speed {
		t.divisionCtr++
		t.speedCtr = 0
	}

	// Increment speed counter
	t.speedCtr++

	// Another frame to draw
	return false
}
