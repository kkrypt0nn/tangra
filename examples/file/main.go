package main

import (
	"github.com/kkrypt0nn/tangra"
	"os"
)

func main() {
	file, err := os.OpenFile("example.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		createdFile, err := os.Create("example.log")
		if err != nil {
			panic(err)
		}
		file = createdFile
	}
	l := tangra.NewLogger()
	l.SetLogFile(file)
	// The styling will get removed when writing into the file.
	l.Debug("${fg:red}${effect:blink}${effect:bold}${sys:username} says hello!")
}
