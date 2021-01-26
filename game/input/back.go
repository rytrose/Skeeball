package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rytrose/soup-the-moon/game/util"
	"github.com/stianeikeland/go-rpio"
)

// backButtonPin is the raspberry pi GPIO pin number for the back button.
const backButtonPin = 6

// backPin is the pin for the back button.
var backPin rpio.Pin

// InitBack configures the raspberry pi GPIO pin.
func InitBack() {
	// Set up back button pin
	if util.IsRasPi() {
		backPin = rpio.Pin(backButtonPin)
		backPin.Input()
		backPin.PullUp()

		// Check pin on update loop
		RegisterPin(backPin)
	}
}

// Back returns true if the back button was pressed.
func Back() bool {
	if util.IsRasPi() {
		if IsRPIOButtonJustPressed(backPin) {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		return true
	}

	return false
}
