package io

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func TestRPIO() {
	// Start RPIO
	RPIOClient.Start()
	RPIOClient.Poll()
	defer RPIOClient.StopPolling()
	defer RPIOClient.Stop()

	pinNames := map[rpio.Pin]string{
		2:  "Pluto",
		3:  "Saturn",
		4:  "Jupiter",
		17: "Mars",
		27: "Earth",
		22: "Mercury",
	}

	generateCallback := func(pinNumber rpio.Pin, name string) func(rpio.Edge) {
		return func(edge rpio.Edge) {
			fmt.Printf("%s (%d) edge: %v\n", name, pinNumber, edge)
		}
	}

	for pinNumber, name := range pinNames {
		fmt.Printf("Setting up %s (%d)\n", name, pinNumber)
		pin := rpio.Pin(pinNumber)
		pin.Input()
		pin.PullUp()

		RPIOClient.RegisterEdgeDetection(pin, rpio.AnyEdge, generateCallback(pinNumber, name))
	}

	stop := time.After(60 * time.Second)
	<-stop

	// for {
	// 	select {
	// 	case <-stop:
	// 		fmt.Println("Exiting...")
	// 		return
	// 	default:
	// 		// Read pin
	// 		<-time.After(100 * time.Millisecond)
	// 	}
	// }
}
