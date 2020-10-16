package main

import (
	"fmt"
	"os"
	"os/signal"

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
