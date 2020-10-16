package io

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// RPIOClient is the RPIO singleton.
var RPIOClient *rPIO

// DefaultPollFreq is the default pin polling frequency.
const DefaultPollFreq = 100 * time.Millisecond

func init() {
	// Instatiate RPIO client singleton
	RPIOClient = &rPIO{
		open: false,
		poller: &rpioPoller{
			ticker:         time.NewTicker(DefaultPollFreq),
			registeredPins: make(map[rpio.Pin]pinRegistration),
			newPin:         make(chan pinRegistration),
			removePin:      make(chan rpio.Pin),
			newPollFreq:    make(chan time.Duration),
			stop:           make(chan struct{}),
		},
		registeredPins: make(map[rpio.Pin]bool),
	}
}

// rPIO is a wrapper interfacing with Raspberry Pi GPIO.
type rPIO struct {
	open           bool              // open maintains state of GPIO.
	poller         *rpioPoller       // poller manages polling pins for edge detection.
	registeredPins map[rpio.Pin]bool // registeredPins keeps track of what pins are registered.
}

// Start opens the GPIO pins and starts polling.
func (r *rPIO) Start() {
	if r.open {
		// Only attempt to open once
		return
	}

	// Open GPIO
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprintf("unable to open GPIO: %s", err))
	}

	// Start polling
	go r.poller.poll()

	r.open = true
}

// Stop closes GPIO and stops polling.
func (r *rPIO) Stop() {
	if !r.open {
		// Don't attempt to stop if not started
		return
	}

	// Stop polling
	r.poller.stop <- struct{}{}

	// Close GPIO
	err := rpio.Close()
	if err != nil {
		panic(fmt.Sprintf("unable to close GPIO: %s", err))
	}

	r.open = false
}

// RegisterEdgeDetection registers a callback for a detected edge on a specified pin.
func (r *rPIO) RegisterEdgeDetection(pin rpio.Pin, edge rpio.Edge, callback func(rpio.Edge)) error {
	if !r.open {
		return fmt.Errorf("GPIO is not yet open")
	}

	_, exists := r.registeredPins[pin]
	if exists {
		return fmt.Errorf("pin is already registered, call RemoveEdgeDetectionRegistration before attempting a new registration")
	}

	fmt.Println(fmt.Sprintf("registering pin %d, edge %v", pin, edge))

	// Only one registration per pin
	r.registeredPins[pin] = true

	// Setup detection
	rpio.DetectEdge(pin, edge)

	// Register with poller
	r.poller.newPin <- pinRegistration{
		pin:      pin,
		edge:     edge,
		callback: callback,
	}

	return nil
}

// RemoveEdgeDetectionRegistration removes an edge detection registration for a specified pin.
func (r *rPIO) RemoveEdgeDetectionRegistration(pin rpio.Pin) error {
	if !r.open {
		return fmt.Errorf("GPIO is not yet open")
	}

	_, exists := r.registeredPins[pin]
	if !exists {
		return fmt.Errorf("pin is not yet registered")
	}

	fmt.Println(fmt.Sprintf("removing registration for pin %d", pin))

	// Remove pin registration
	delete(r.registeredPins, pin)

	// Clear detection
	rpio.DetectEdge(pin, rpio.NoEdge)

	// Remove registration with poller
	r.poller.removePin <- pin

	return nil
}

// UpdatePollFreq changes the polling frequency of edge detection.
func (r *rPIO) UpdatePollFreq(d time.Duration) error {
	if !r.open {
		return fmt.Errorf("polling has not yet started")
	}

	// Update the poller frequency
	r.poller.newPollFreq <- d

	return nil
}

// pinRegistration is a registration for a callback when an edge is detected for a pin.
type pinRegistration struct {
	pin      rpio.Pin        // pin is the pin to monitor for edge detection.
	edge     rpio.Edge       // edge is the type of edge to run the callback on.
	callback func(rpio.Edge) // callback is the function to run when an edge is detected.
}

// rpioPoller manages polling pins for edge detection.
type rpioPoller struct {
	ticker         *time.Ticker                 // ticker manages the polling period.
	registeredPins map[rpio.Pin]pinRegistration // registeredPins contains which pins should be polled for what edge detection.
	newPin         chan pinRegistration         // newPins allows a new pin to be incorporated into polling.
	removePin      chan rpio.Pin                // removePin allows a pin to be removed from polling.
	newPollFreq    chan time.Duration           // newPollFreq updates the polling frequency.
	stop           chan struct{}                // stop ends polling.
}

// poll starts the pin polling routine.
func (p *rpioPoller) poll() {
pollLoop:
	for {
		select {
		case <-p.ticker.C:
			fmt.Println("poller: polling")
			// Read pins and handle edge detection
			for pin, registration := range p.registeredPins {
				if pin.EdgeDetected() {
					go registration.callback(registration.edge)
				}
			}
		case newRegistration := <-p.newPin:
			fmt.Println("poller: new pin")
			// Add pin registration to pins to poll
			p.registeredPins[newRegistration.pin] = newRegistration
		case registrationToRemove := <-p.removePin:
			fmt.Println("poller: removing pin")
			// Remove pin registration from pins to poll
			delete(p.registeredPins, registrationToRemove)
		case newPollFreq := <-p.newPollFreq:
			fmt.Println("poller: updating polling frequency")
			// Update the ticker polling frequency
			p.ticker.Reset(newPollFreq)
		case <-p.stop:
			break pollLoop
		}
	}
}
