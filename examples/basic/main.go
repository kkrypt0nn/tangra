package main

import (
	"github.com/kkrypt0nn/tangra/v2"
)

func main() {
	l := tangra.NewLogger()
	l.Debug("Debug message example")
	world := "world"
	l.Infof("Hello %s", world)
}
