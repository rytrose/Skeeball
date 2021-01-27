package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rytrose/soup-the-moon/game/util"
	"github.com/stianeikeland/go-rpio"
)

// marsButtonPin is the raspberry pi GPIO pin number for the mars button.
const marsButtonPin = 8

// marsPin is the pin for the mars button.
var marsPin rpio.Pin

// InitMars configures the raspberry pi GPIO pin.
func InitMars() {
	// Set up mars button pin
	if util.IsRasPi() {
		marsPin = rpio.Pin(marsButtonPin)
		marsPin.Input()
		marsPin.PullUp()

		// Check pin on update loop
		RegisterPin(marsPin)
	}
}

// Mars returns true if the mars button was pressed.
func Mars() bool {
	if util.IsRasPi() {
		if IsRPIOButtonJustPressed(marsPin) {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		return true
	}

	return false
}
