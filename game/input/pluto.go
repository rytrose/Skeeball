package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rytrose/soup-the-moon/game/util"
	"github.com/stianeikeland/go-rpio"
)

// plutoButtonPin is the raspberry pi GPIO pin number for the pluto button.
const plutoButtonPin = 11

// plutoPin is the pin for the pluto button.
var plutoPin rpio.Pin

// InitPluto configures the raspberry pi GPIO pin.
func InitPluto() {
	// Set up pluto button pin
	if util.IsRasPi() {
		plutoPin = rpio.Pin(plutoButtonPin)
		plutoPin.Input()
		plutoPin.PullUp()

		// Check pin on update loop
		RegisterPin(plutoPin)
	}
}

// Pluto returns true if the pluto button was pressed.
func Pluto() bool {
	if util.IsRasPi() {
		if IsRPIOButtonJustPressed(plutoPin) {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		return true
	}

	return false
}
