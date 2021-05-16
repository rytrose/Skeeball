package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rytrose/soup-the-moon/game/util"
	"github.com/stianeikeland/go-rpio/v4"
)

// earthButtonPin is the raspberry pi GPIO pin number for the earth button.
const earthButtonPin = 27

// earthPin is the pin for the earth button.
var earthPin rpio.Pin

// InitEarth configures the raspberry pi GPIO pin.
func InitEarth() {
	// Set up earth button pin
	if util.IsRasPi() {
		earthPin = rpio.Pin(earthButtonPin)
		earthPin.Input()
		earthPin.PullUp()

		// Check pin on update loop
		RegisterPin(earthPin)
	}
}

// Earth returns true if the earth button was pressed.
func Earth() bool {
	if util.IsRasPi() {
		if IsRPIOButtonJustPressed(earthPin) {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		return true
	}

	return false
}
