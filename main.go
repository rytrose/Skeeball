package main

import (
	"github.com/rytrose/soup-the-moon/game"
	"github.com/rytrose/soup-the-moon/game/util"
	"github.com/rytrose/soup-the-moon/io"
)

func init() {
	if util.IsRasPi() {
		// Open GPIO on init
		io.RPIOClient.Start()
	}
}

func main() {
	if util.IsRasPi() {
		// Close GPIO on exit
		defer io.RPIOClient.Stop()
	}

	game.Run()
}

// func main() {
// 	io.TestRPIO()
// }
