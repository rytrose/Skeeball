package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rytrose/soup-the-moon/game/util"
	"github.com/stianeikeland/go-rpio"
)

// downButtonPin is the raspberry pi GPIO pin number for the down button.
const downButtonPin = 23

// downPin is the pin for the down button.
var downPin rpio.Pin

// InitDown configures the raspberry pi GPIO pin.
func InitDown() {
	// Set up down button pin
	if util.IsRasPi() {
		downPin = rpio.Pin(downButtonPin)
		downPin.Input()
		downPin.PullUp()

		// Check pin on update loop
		RegisterPin(downPin)
	}
}

// Down returns true if the down button was pressed.
func Down() bool {
	if util.IsRasPi() {
		if IsRPIOButtonJustPressed(downPin) {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		return true
	}

	return false
}
