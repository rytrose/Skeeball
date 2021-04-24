package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rytrose/soup-the-moon/game/util"
	"github.com/stianeikeland/go-rpio"
)

// upButtonPin is the raspberry pi GPIO pin number for the up button.
const upButtonPin = 24

// upPin is the pin for the up button.
var upPin rpio.Pin

// InitUp configures the raspberry pi GPIO pin.
func InitUp() {
	// Set up up button pin
	if util.IsRasPi() {
		upPin = rpio.Pin(upButtonPin)
		upPin.Input()
		upPin.PullUp()

		// Check pin on update loop
		RegisterPin(upPin)
	}
}

// Up returns true if the up button was pressed.
func Up() bool {
	if util.IsRasPi() {
		if IsRPIOButtonJustPressed(upPin) {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		return true
	}

	return false
}
