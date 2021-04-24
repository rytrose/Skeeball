package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rytrose/soup-the-moon/game/util"
	"github.com/stianeikeland/go-rpio"
)

// enterButtonPin is the raspberry pi GPIO pin number for the enter button.
const enterButtonPin = 18

// enterPin is the pin for the enter button.
var enterPin rpio.Pin

// InitEnter configures the raspberry pi GPIO pin.
func InitEnter() {
	// Set up enter button pin
	if util.IsRasPi() {
		enterPin = rpio.Pin(enterButtonPin)
		enterPin.Input()
		enterPin.PullUp()

		// Check pin on update loop
		RegisterPin(enterPin)
	}
}

// Enter returns true if the enter button was pressed.
func Enter() bool {
	if util.IsRasPi() {
		if IsRPIOButtonJustPressed(enterPin) {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return true
	}

	return false
}
