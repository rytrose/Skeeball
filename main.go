package main

import (
	"fmt"
	"time"

	"github.com/rytrose/skeeball/io"
	"github.com/stianeikeland/go-rpio"
)

func main() {
	// Start RPIO
	io.RPIOClient.Start()
	defer io.RPIOClient.Stop()

	pin := rpio.Pin(27)
	pin.Input()
	pin.PullUp()

	callback := func(edge rpio.Edge) {
		fmt.Println(fmt.Sprintf("edge detected: %v (%v)", edge, pin.Read()))
	}

	io.RPIOClient.RegisterEdgeDetection(pin, rpio.RiseEdge, callback)

	<-time.After(10 * time.Second)
}
