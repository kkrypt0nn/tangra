package main

import "github.com/kkrypt0nn/tangra"

func main() {
	l := tangra.NewLogger()
	l.Println("${bold}${underline}Styling:")
	l.Println("${bold}Foreground colors:")
	l.Println("${fg:black}Black${reset}\t${fg:red}Red${reset}\t\t${fg:green}Green${reset}\t\t${fg:yellow}Yellow${reset}\t\t${fg:blue}Blue${reset}\t\t${fg:purple}Purple${reset}\t\t${fg:cyan}Cyan${reset}\t\t${fg:white}White")
	l.Println("${fg:gray}Gray${reset}\t${fg:brightred}Bright red${reset}\t${fg:brightgreen}Bright green${reset}\t${fg:brightyellow}Bright yellow\t${fg:brightblue}Bright blue${reset}\t${fg:brightpurple}Bright purple${reset}\t${fg:brightcyan}Bright cyan${reset}\t${fg:brightwhite}Bright white")
	l.Println("${bold}Background colors:")
	l.Println("${bg:black}Black${reset}\t${bg:red}Red${reset}\t\t${bg:green}Green${reset}\t\t${bg:yellow}Yellow${reset}\t\t${bg:blue}Blue${reset}\t\t${bg:purple}Purple${reset}\t\t${bg:cyan}Cyan${reset}\t\t${bg:white}White")
	l.Println("${bg:gray}Gray${reset}\t${bg:brightred}Bright red${reset}\t${bg:brightgreen}Bright green${reset}\t${bg:brightyellow}Bright yellow\t${bg:brightblue}Bright blue${reset}\t${bg:brightpurple}Bright purple${reset}\t${bg:brightcyan}Bright cyan${reset}\t${bg:brightwhite}Bright white")
	l.Println("${bold}Special Effects:")
	l.Println("${effect:bold}Bold${reset}\t${effect:dim}Dim${reset}\t\t${effect:underline}Underline${reset}\t${effect:blink}Blink${reset}\t\t${effect:inverse}Inverse${reset}\t\t${effect:strikethrough}Strikethrough${reset}")

	l.Println("\n${bold}${underline}Variables:")
	l.Println("${bold}Caller:")
	l.Println("Function: ${caller:function}\tShort function: ${caller:shortfunction}\t\tFile: ${caller:file}\t\t\t\tLine: ${caller:line}")
	l.Println("${bold}Logging Level:")
	l.SetLoggingLevel(tangra.FATAL)
	l.Println("Level Color: ${level:color}Color${reset}\tLevel Name: ${level:name}\t\tLevel Short Name: ${level:shortname}")
	l.SetLoggingLevel(tangra.NONE)
	l.Println("${bold}Date & Time Now:")
	l.Println("Date: ${now:date}\tTime: ${now:time}\t\t\tDate & Time: ${now:datetime}")
	l.Println("${bold}System:")
	l.Println("Architecture: ${sys:architecture}\tHostname: ${sys:hostname}\tOperating System: ${sys:operating_system}\t\tUsername: ${username}")
	l.Println("Group ID: ${sys:groupid}\t\tUser ID: ${sys:userid}")
}
