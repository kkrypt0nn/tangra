package main

import "github.com/kkrypt0nn/tangra/v2"

func main() {
	l := tangra.NewLogger()
	l.Info("${fg:red}${effect:blink}${effect:bold}${sys:username} says hello!")
}
