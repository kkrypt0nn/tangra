package tangra

import (
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/kkrypt0nn/tangra/v2/terminal"
)

type FormatMode int

const (
	HumanFormat FormatMode = iota
	JSONFormat
)

func (l Logger) aliases() map[string]string {
	return map[string]string{
		// Variables
		"${date}":             "${now:date}",
		"${time}":             "${now:time}",
		"${datetime}":         "${now:datetime}",
		"${sys:arch}":         "${sys:architecture}",
		"${architecture}":     "${sys:architecture}",
		"${arch}":             "${sys:architecture}",
		"${sys:os}":           "${sys:operating_system}",
		"${operating_system}": "${sys:operating_system}",
		"${os}":               "${sys:operating_system}",
		"${hostname}":         "${sys:hostname}",
		"${username}":         "${sys:username}",
		"${groupid}":          "${sys:groupid}",
		"${userid}":           "${sys:userid}",
		"${function}":         "${caller:function}",
		"${shortfunction}":    "${caller:shortfunction}",
		"${file}":             "${caller:file}",
		"${line}":             "${caller:line}",

		// Styles
		"${black}":         "${fg:black}",
		"${red}":           "${fg:red}",
		"${green}":         "${fg:green}",
		"${yellow}":        "${fg:yellow}",
		"${blue}":          "${fg:blue}",
		"${purple}":        "${fg:purple}",
		"${cyan}":          "${fg:cyan}",
		"${white}":         "${fg:white}",
		"${gray}":          "${fg:gray}",
		"${brightred}":     "${fg:brightred}",
		"${brightgreen}":   "${fg:brightgreen}",
		"${brightyellow}":  "${fg:brightyellow}",
		"${brightblue}":    "${fg:brightblue}",
		"${brightpurple}":  "${fg:brightpurple}",
		"${brightcyan}":    "${fg:brightcyan}",
		"${brightwhite}":   "${fg:brightwhite}",
		"${bblack}":        "${bg:black}",
		"${bred}":          "${bg:red}",
		"${bgreen}":        "${bg:green}",
		"${byellow}":       "${bg:yellow}",
		"${bblue}":         "${bg:blue}",
		"${bpurple}":       "${bg:purple}",
		"${bcyan}":         "${bg:cyan}",
		"${bwhite}":        "${bg:white}",
		"${bgray}":         "${bg:gray}",
		"${bbrightred}":    "${bg:brightred}",
		"${bbrightgreen}":  "${bg:brightgreen}",
		"${bbrightyellow}": "${bg:brightyellow}",
		"${bbrightblue}":   "${bg:brightblue}",
		"${bbrightpurple}": "${bg:brightpurple}",
		"${bbrightcyan}":   "${bg:brightcyan}",
		"${bbrightwhite}":  "${bg:brightwhite}",
		"${bold}":          "${effect:bold}",
		"${dim}":           "${effect:dim}",
		"${underline}":     "${effect:underline}",
		"${blink}":         "${effect:blink}",
		"${inverse}":       "${effect:inverse}",
		"${strikethrough}": "${effect:strikethrough}",
		"${reset}":         "${effect:reset}",
	}
}

func (l Logger) styles() map[string]string {
	return map[string]string{
		// Foreground Colors
		"${fg:black}":        terminal.BLACK,
		"${fg:red}":          terminal.RED,
		"${fg:green}":        terminal.GREEN,
		"${fg:yellow}":       terminal.YELLOW,
		"${fg:blue}":         terminal.BLUE,
		"${fg:purple}":       terminal.PURPLE,
		"${fg:cyan}":         terminal.CYAN,
		"${fg:white}":        terminal.WHITE,
		"${fg:gray}":         terminal.GRAY,
		"${fg:brightred}":    terminal.BRIGHT_RED,
		"${fg:brightgreen}":  terminal.BRIGHT_GREEN,
		"${fg:brightyellow}": terminal.BRIGHT_YELLOW,
		"${fg:brightblue}":   terminal.BRIGHT_BLUE,
		"${fg:brightpurple}": terminal.BRIGHT_PURPLE,
		"${fg:brightcyan}":   terminal.BRIGHT_CYAN,
		"${fg:brightwhite}":  terminal.BRIGHT_WHITE,
		// Background Colors
		"${bg:black}":        terminal.BG_BLACK,
		"${bg:red}":          terminal.BG_RED,
		"${bg:green}":        terminal.BG_GREEN,
		"${bg:yellow}":       terminal.BG_YELLOW,
		"${bg:blue}":         terminal.BG_BLUE,
		"${bg:purple}":       terminal.BG_PURPLE,
		"${bg:cyan}":         terminal.BG_CYAN,
		"${bg:white}":        terminal.BG_WHITE,
		"${bg:gray}":         terminal.BG_GRAY,
		"${bg:brightred}":    terminal.BG_BRIGHT_RED,
		"${bg:brightgreen}":  terminal.BG_BRIGHT_GREEN,
		"${bg:brightyellow}": terminal.BG_BRIGHT_YELLOW,
		"${bg:brightblue}":   terminal.BG_BRIGHT_BLUE,
		"${bg:brightpurple}": terminal.BG_BRIGHT_PURPLE,
		"${bg:brightcyan}":   terminal.BG_BRIGHT_CYAN,
		"${bg:brightwhite}":  terminal.BG_BRIGHT_WHITE,
		// Special Effects
		"${effect:bold}":          terminal.BOLD,
		"${effect:dim}":           terminal.DIM,
		"${effect:underline}":     terminal.UNDERLINE,
		"${effect:blink}":         terminal.BLINK,
		"${effect:inverse}":       terminal.INVERSE,
		"${effect:strikethrough}": terminal.STRIKETHROUGH,
		"${effect:reset}":         terminal.RESET,
	}
}

func (l Logger) variables(level Level) map[string]func() string {
	return map[string]func() string{
		// Caller
		"${caller:function}": func() string {
			pc, _, _, ok := runtime.Caller(4)
			details := runtime.FuncForPC(pc)
			if ok && details != nil {
				return details.Name()
			}
			return ""
		},
		"${caller:shortfunction}": func() string {
			pc, _, _, ok := runtime.Caller(4)
			details := runtime.FuncForPC(pc)
			split := strings.Split(details.Name(), ".")
			if ok && details != nil && len(split) >= 2 {
				return split[len(split)-1]
			}
			return ""
		},
		"${caller:file}": func() string {
			_, file, _, ok := runtime.Caller(4)
			split := strings.Split(file, "/")
			if ok && len(split) >= 1 {
				return split[len(split)-1]
			}
			return ""
		},
		"${caller:line}": func() string {
			_, _, line, ok := runtime.Caller(4)
			if ok {
				return strconv.Itoa(line)
			}
			return "0"
		},
		// Logging Level
		"${level:color}": func() string {
			return GetLevelColor(level, l.forceStyling)
		},
		"${level:lowername}": func() string {
			return strings.ToLower(GetLevelName(level))
		},
		"${level:name}": func() string {
			return GetLevelName(level)
		},
		"${level:shortname}": func() string {
			return GetLevelShortName(level)
		},
		// Date & Time Now
		"${now:date}": func() string {
			return time.Now().Format(l.dateFormat)
		},
		"${now:time}": func() string {
			return time.Now().Format(l.timeFormat)
		},
		"${now:datetime}": func() string {
			return time.Now().Format(l.datetimeFormat)
		},
		// System
		"${sys:architecture}": func() string {
			return runtime.GOARCH
		},
		"${sys:hostname}": func() string {
			hostname, err := os.Hostname()
			if err != nil {
				return ""
			}
			return hostname
		},
		"${sys:operating_system}": func() string {
			return runtime.GOOS
		},
		"${sys:username}": func() string {
			user, err := user.Current()
			if err != nil {
				return ""
			}
			return user.Username
		},
		"${sys:groupid}": func() string {
			user, err := user.Current()
			if err != nil {
				return ""
			}
			return user.Gid
		},
		"${sys:userid}": func() string {
			user, err := user.Current()
			if err != nil {
				return ""
			}
			return user.Uid
		},
	}
}

func (l Logger) replaceAliases(message string) string {
	for a, v := range l.aliases() {
		message = strings.ReplaceAll(message, a, v)
	}
	return message
}

func (l Logger) addStyling(message string) string {
	for e, v := range l.styles() {
		message = strings.ReplaceAll(message, e, v)
	}
	return message
}

func (l Logger) removeStyling(message string) string {
	for e, v := range l.styles() {
		message = strings.ReplaceAll(message, e, "")
		message = strings.ReplaceAll(message, v, "")
	}
	return message
}

func (l Logger) addVariables(level Level, message string) string {
	message = l.replaceAliases(message)
	for variable, v := range l.variables(level) {
		message = strings.ReplaceAll(message, variable, v())
	}
	return message
}
