package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rytrose/soup-the-moon/game/util"
	"github.com/stianeikeland/go-rpio/v4"
)

// mercuryButtonPin is the raspberry pi GPIO pin number for the mercury button.
const mercuryButtonPin = 22

// mercuryPin is the pin for the mercury button.
var mercuryPin rpio.Pin

// InitMercury configures the raspberry pi GPIO pin.
func InitMercury() {
	// Set up mercury button pin
	if util.IsRasPi() {
		mercuryPin = rpio.Pin(mercuryButtonPin)
		mercuryPin.Input()
		mercuryPin.PullUp()

		// Check pin on update loop
		RegisterPin(mercuryPin)
	}
}

// Mercury returns true if the mercury button was pressed.
func Mercury() bool {
	if util.IsRasPi() {
		if IsRPIOButtonJustPressed(mercuryPin) {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		return true
	}

	return false
}
