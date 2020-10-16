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

	pin := rpio.Pin(4)
	pin.Input()

	callback := func(edge rpio.Edge) {
		fmt.Println(fmt.Sprintf("edge detected: %v", edge))
	}

	io.RPIOClient.RegisterEdgeDetection(pin, rpio.AnyEdge, callback)

	<-time.After(10 * time.Second)
}
