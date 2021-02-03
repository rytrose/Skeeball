package input

import (
	"sync"

	"github.com/stianeikeland/go-rpio"
)

type rpioButtonState struct {
	buttonDurations     map[rpio.Pin]int
	prevButtonDurations map[rpio.Pin]int

	m sync.RWMutex
}

var theRPIOButtonState = &rpioButtonState{
	buttonDurations:     map[rpio.Pin]int{},
	prevButtonDurations: map[rpio.Pin]int{},
}

// update reads registered pins and updates internal state.
func (s *rpioButtonState) update() {
	s.m.Lock()
	defer s.m.Unlock()

	// Read pins, update durations
	for pin := range s.buttonDurations {
		// Update previous duration counter
		s.prevButtonDurations[pin] = s.buttonDurations[pin]

		// Read pin and update duration
		if rpio.ReadPin(pin) == rpio.Low {
			s.buttonDurations[pin]++
		} else {
			s.buttonDurations[pin] = 0
		}
	}
}

// RPIOButtonUpdate should be called in the game update loop.
func RPIOButtonUpdate() {
	theRPIOButtonState.update()
}

// RegisterPin configures a pin to be read on game update loop.
// Assumes buttons are configured such that a pressed button reads as rpio.Low.
func RegisterPin(pin rpio.Pin) {
	theRPIOButtonState.m.Lock()
	defer theRPIOButtonState.m.Unlock()

	// Set pin durations to 0
	theRPIOButtonState.prevButtonDurations[pin] = 0
	theRPIOButtonState.buttonDurations[pin] = 0
}

// DeregisterPin stops reading a pin on game update loop.
func DeregisterPin(pin rpio.Pin) {
	theRPIOButtonState.m.Lock()
	defer theRPIOButtonState.m.Unlock()

	// Remove pin from duration states
	delete(theRPIOButtonState.prevButtonDurations, pin)
	delete(theRPIOButtonState.buttonDurations, pin)
}

// IsRPIOButtonJustPressed returns a boolean value indicating
// whether the given RPIO button was pressed just in the current frame.
//
// IsRPIOButtonJustPressed is concurrency-safe.
func IsRPIOButtonJustPressed(pin rpio.Pin) bool {
	return RPIOButtonDuration(pin) == 1
}

// IsRPIOButtonJustReleased returns a boolean value indicating
// whether the given RPIO button was released just in the current frame.
//
// IsRPIOButtonJustReleased is concurrency-safe.
func IsRPIOButtonJustReleased(pin rpio.Pin) bool {
	theRPIOButtonState.m.RLock()
	r := theRPIOButtonState.buttonDurations[pin] == 0 &&
		theRPIOButtonState.prevButtonDurations[pin] > 0
	theRPIOButtonState.m.RUnlock()
	return r
}

// RPIOButtonDuration returns how long the RPIO button has been pressed in frames.
//
// RPIOButtonDuration is concurrency-safe.
func RPIOButtonDuration(pin rpio.Pin) int {
	theRPIOButtonState.m.RLock()
	frames := theRPIOButtonState.buttonDurations[pin]
	theRPIOButtonState.m.RUnlock()
	return frames
}
