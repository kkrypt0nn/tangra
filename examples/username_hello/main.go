package main

import "github.com/kkrypt0nn/tangra"

func main() {
	l := tangra.NewLogger()
	l.Println("${fg:red}${effect:blink}${effect:bold}${sys:username} says hello!")
}
