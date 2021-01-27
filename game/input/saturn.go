package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rytrose/soup-the-moon/game/util"
	"github.com/stianeikeland/go-rpio"
)

// saturnButtonPin is the raspberry pi GPIO pin number for the saturn button.
const saturnButtonPin = 10

// saturnPin is the pin for the saturn button.
var saturnPin rpio.Pin

// InitSaturn configures the raspberry pi GPIO pin.
func InitSaturn() {
	// Set up saturn button pin
	if util.IsRasPi() {
		saturnPin = rpio.Pin(saturnButtonPin)
		saturnPin.Input()
		saturnPin.PullUp()

		// Check pin on update loop
		RegisterPin(saturnPin)
	}
}

// Saturn returns true if the saturn button was pressed.
func Saturn() bool {
	if util.IsRasPi() {
		if IsRPIOButtonJustPressed(saturnPin) {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		return true
	}

	return false
}
