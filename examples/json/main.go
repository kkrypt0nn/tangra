package main

import (
	"github.com/kkrypt0nn/tangra/v2"
)

func main() {
	l := tangra.NewLogger().WithFormatMode(tangra.JSONFormat)
	l.Debug("Debug message example")
}
