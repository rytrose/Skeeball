package io

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func testRPIO() {
	// Start RPIO
	RPIOClient.Start()
	defer RPIOClient.Stop()

	pin := rpio.Pin(4)
	pin.Input()
	pin.PullUp()

	callback := func(edge rpio.Edge) {
		fmt.Println(fmt.Sprintf("edge detected: %v (%v)", edge, pin.Read()))
	}

	RPIOClient.RegisterEdgeDetection(pin, rpio.FallEdge, callback)

	<-time.After(10 * time.Second)
}
