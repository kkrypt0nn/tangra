package terminal

import "os"

const RESET = "\033[0m"

func AreColorsSupported() bool {
	// Usual way to check
	hasTerm := os.Getenv("TERM") != ""
	// Windows Terminal may have a session variable set. Not always as it's inconsistent, see GitHub issues about it...
	hasWTSession := os.Getenv("WT_SESSION") != ""
	// When the Command Prompt is started within Windows Terminal, there is this environment variable set
	hasSessionNameConsole := os.Getenv("SESSIONNAME") == "Console"

	// If any of the conditions are met, it most likely can handle colors
	return hasTerm || hasWTSession || hasSessionNameConsole
}
