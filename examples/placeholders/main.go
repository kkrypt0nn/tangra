package main

import "github.com/kkrypt0nn/tangra/v2"

func main() {
	l := tangra.NewLogger()
	l.Info("${bold}${underline}Styling:")
	l.Info("${bold}Foreground colors:")
	l.Info("${fg:black}Black${reset}\t${fg:red}Red${reset}\t\t${fg:green}Green${reset}\t\t${fg:yellow}Yellow${reset}\t\t${fg:blue}Blue${reset}\t\t${fg:purple}Purple${reset}\t\t${fg:cyan}Cyan${reset}\t\t${fg:white}White")
	l.Info("${fg:gray}Gray${reset}\t${fg:brightred}Bright red${reset}\t${fg:brightgreen}Bright green${reset}\t${fg:brightyellow}Bright yellow\t${fg:brightblue}Bright blue${reset}\t${fg:brightpurple}Bright purple${reset}\t${fg:brightcyan}Bright cyan${reset}\t${fg:brightwhite}Bright white")
	l.Info("${bold}Background colors:")
	l.Info("${bg:black}Black${reset}\t${bg:red}Red${reset}\t\t${bg:green}Green${reset}\t\t${bg:yellow}Yellow${reset}\t\t${bg:blue}Blue${reset}\t\t${bg:purple}Purple${reset}\t\t${bg:cyan}Cyan${reset}\t\t${bg:white}White")
	l.Info("${bg:gray}Gray${reset}\t${bg:brightred}Bright red${reset}\t${bg:brightgreen}Bright green${reset}\t${bg:brightyellow}Bright yellow\t${bg:brightblue}Bright blue${reset}\t${bg:brightpurple}Bright purple${reset}\t${bg:brightcyan}Bright cyan${reset}\t${bg:brightwhite}Bright white")
	l.Info("${bold}Special Effects:")
	l.Info("${effect:bold}Bold${reset}\t${effect:dim}Dim${reset}\t\t${effect:underline}Underline${reset}\t${effect:blink}Blink${reset}\t\t${effect:inverse}Inverse${reset}\t\t${effect:strikethrough}Strikethrough${reset}")

	l.Info("\n${bold}${underline}Variables:")
	l.Info("${bold}Caller:")
	l.Info("Function: ${caller:function}\tShort function: ${caller:shortfunction}\t\tFile: ${caller:file}\t\t\t\tLine: ${caller:line}")
	l.Info("${bold}Logging Level:")
	l.Fatal("Level Color: ${level:color}Color${reset}\tLevel Name: ${level:name}\t\tLevel Short Name: ${level:shortname}")
	l.Info("${bold}Date & Time Now:")
	l.Info("Date: ${now:date}\tTime: ${now:time}\t\t\tDate & Time: ${now:datetime}")
	l.Info("${bold}System:")
	l.Info("Architecture: ${sys:architecture}\tHostname: ${sys:hostname}\tOperating System: ${sys:operating_system}\t\tUsername: ${username}")
	l.Info("Group ID: ${sys:groupid}\t\tUser ID: ${sys:userid}")
}
