package tangra

import (
	"fmt"
	"os"

	"github.com/kkrypt0nn/tangra/v2/terminal"
)

// Logger is represents a logger structure.
type Logger struct {
	// Styling describes whether the logger should style the logged message.
	Styling bool
	// Prefix is the prefix before the logged message.
	Prefix string
	// LogFile is the log file to write the messages into.
	LogFile *os.File

	loggingLevel   Level
	dateFormat     string
	datetimeFormat string
	timeFormat     string
	forceStyling   bool
}

// NewLogger creates a new logger.
func NewLogger() *Logger {
	return &Logger{
		Styling:        true,
		Prefix:         "${datetime} ${level:color}${level:name}${reset}: ",
		loggingLevel:   NONE,
		dateFormat:     "Jan 02, 2006",
		datetimeFormat: "Jan 02, 2006 15:04:05",
		timeFormat:     "15:04:05",
	}
}

// SetForceStyling sets to whether it should force the styling render.
func (l *Logger) SetForceStyling(forceStyling bool) {
	l.forceStyling = forceStyling
}

// SetLogFile will set the log file to write logs into.
func (l *Logger) SetLogFile(file *os.File) {
	l.LogFile = file
}

// SetDateFormat sets the format to use when logging the date.
func (l *Logger) SetDateFormat(format string) {
	l.dateFormat = format
}

// SetDatetimeFormat sets the format to use when logging the date and time.
func (l *Logger) SetDatetimeFormat(format string) {
	l.datetimeFormat = format
}

// SetTimeFormat sets the format to use when logging the time.
func (l *Logger) SetTimeFormat(format string) {
	l.timeFormat = format
}

// SetStyling sets the styling setting of the logger.
func (l *Logger) SetStyling(styling bool) {
	l.Styling = styling
}

// SetPrefix sets the prefix before the log message.
func (l *Logger) SetPrefix(prefix string) {
	l.Prefix = prefix
}

// SetLevelColor sets the color of the level.
func (l *Logger) SetLevelColor(id Level, color string) {
	LevelColors[id] = color
}

func (l *Logger) doLog(level Level, message string) {
	previousLevel := l.loggingLevel
	l.loggingLevel = level
	defer func() { l.loggingLevel = previousLevel }()

	formatted := l.Prefix + message + terminal.RESET
	formatted = l.formatMessage(formatted, true)

	fmt.Println(formatted)

	if l.LogFile != nil {
		if _, err := l.LogFile.WriteString(l.removeStyling(formatted) + "\n"); err != nil {
			panic(err)
		}
	}
}

// Debug prints a debug message.
func (l *Logger) Debug(message string) {
	l.doLog(DEBUG, message)
}

// Info prints an info message.
func (l *Logger) Info(message string) {
	l.doLog(INFO, message)
}

// Warn prints a warning message.
func (l *Logger) Warn(message string) {
	l.doLog(WARN, message)
}

// Error prints an error message.
func (l *Logger) Error(message string) {
	l.doLog(ERROR, message)
}

// Fatal prints a fatal error message.
func (l *Logger) Fatal(message string) {
	l.doLog(FATAL, message)
}

// Trace prints a tracing message.
func (l *Logger) Trace(message string) {
	l.doLog(TRACE, message)
}

// Log will log with the given level.
func (l *Logger) Log(level Level, message string) {
	l.doLog(level, message)
}

// SetLoggingLevel sets the logging level.
func (l *Logger) SetLoggingLevel(level Level) {
	l.loggingLevel = level
}

// Print simply prints the message, without logging level.
func (l *Logger) Print(message string) {
	fmt.Print(l.formatMessage(message, true) + terminal.RESET)
}

// Println simply prints the message with a new line, without logging level.
func (l *Logger) Println(message string) {
	fmt.Println(l.formatMessage(message, true) + terminal.RESET)
}

// formatMessage formats a message with or without styling
func (l *Logger) formatMessage(message string, withStyling bool) string {
	message = l.addVariables(message)
	if withStyling && ((l.Styling && terminal.AreColorsSupported()) || l.forceStyling) {
		return l.addStyling(message)
	}
	return l.removeStyling(message)
}
