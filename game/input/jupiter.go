package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rytrose/soup-the-moon/game/util"
	"github.com/stianeikeland/go-rpio/v4"
)

// jupiterButtonPin is the raspberry pi GPIO pin number for the jupiter button.
const jupiterButtonPin = 4

// jupiterPin is the pin for the jupiter button.
var jupiterPin rpio.Pin

// InitJupiter configures the raspberry pi GPIO pin.
func InitJupiter() {
	// Set up jupiter button pin
	if util.IsRasPi() {
		jupiterPin = rpio.Pin(jupiterButtonPin)
		jupiterPin.Input()
		jupiterPin.PullUp()

		// Check pin on update loop
		RegisterPin(jupiterPin)
	}
}

// Jupiter returns true if the jupiter button was pressed.
func Jupiter() bool {
	if util.IsRasPi() {
		if IsRPIOButtonJustPressed(jupiterPin) {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		return true
	}

	return false
}
