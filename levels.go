package tangra

import "github.com/kkrypt0nn/tangra/terminal"

// Level represents the level of logging - https://logging.apache.org/log4j/2.x/log4j-api/apidocs/org/apache/logging/log4j/Level.html
type Level int

const (
	// TRACE level designates finer-grained informational events than the DEBUG.
	TRACE Level = iota
	// DEBUG level designates fine-grained informational events that are most useful to debug an application.
	DEBUG
	// INFO level designates informational messages that highlight the progress of the application at coarse-grained level.
	INFO
	// WARN level designates potentially harmful situations.
	WARN
	// ERROR level designates error events that might still allow the application to continue running.
	ERROR
	// FATAL level designates very severe error events that will presumably lead the application to abort.
	FATAL
	// NONE will reset the logging level to nothing.
	NONE Level = 1337
)

var (
	// LevelColors convert a logging level to its color.
	LevelColors = map[Level]string{
		TRACE: terminal.GRAY,
		DEBUG: terminal.GRAY + terminal.BOLD,
		INFO:  terminal.BLUE + terminal.BOLD,
		WARN:  terminal.YELLOW + terminal.BOLD,
		ERROR: terminal.RED + terminal.BOLD,
		FATAL: terminal.BRIGHT_WHITE + terminal.BOLD + terminal.BG_RED,
	}
	// LevelNames convert a logging level to its name.
	LevelNames = map[Level]string{
		TRACE: "TRACE",
		DEBUG: "DEBUG",
		INFO:  "INFO",
		WARN:  "WARN",
		ERROR: "ERROR",
		FATAL: "FATAL",
	}
	// ShortLevelNames convert a logging level to its short name.
	ShortLevelNames = map[Level]string{
		TRACE: "tra",
		DEBUG: "dbg",
		INFO:  "inf",
		WARN:  "war",
		ERROR: "err",
		FATAL: "fat",
	}
)

// GetLevelColor returns the color of the level.
func GetLevelColor(id Level) string {
	if terminal.AreColorsSupported() || ForceStyling {
		return LevelColors[id]
	}
	return ""
}

// GetLevelName returns the name of the level.
func GetLevelName(id Level) string {
	return LevelNames[id]
}

// GetLevelShortName returns the short name of the level.
func GetLevelShortName(id Level) string {
	return ShortLevelNames[id]
}
