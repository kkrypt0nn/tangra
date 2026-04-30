package tangra

import (
	"io"

	"github.com/kkrypt0nn/tangra/v2/terminal"
)

type Logger struct {
	prefix         string
	out            io.Writer
	logFile        io.Writer
	loggingLevel   Level
	dateFormat     string
	datetimeFormat string
	timeFormat     string
	forceStyling   bool
	styling        bool
}

func NewLogger() Logger {
	return Logger{
		prefix:         "${datetime} ${level:color}${level:name}${reset}: ",
		loggingLevel:   NONE,
		dateFormat:     "Jan 02, 2006",
		datetimeFormat: "Jan 02, 2006 15:04:05",
		timeFormat:     "15:04:05",
		styling:        true,
	}
}

func (l Logger) WithPrefix(prefix string) Logger {
	l.prefix = prefix
	return l
}

func (l Logger) WithLevel(level Level) Logger {
	l.loggingLevel = level
	return l
}

func (l Logger) WithForceStyling(force bool) Logger {
	l.forceStyling = force
	return l
}

func (l Logger) WithLogFile(w io.Writer) Logger {
	l.logFile = w
	return l
}

func (l Logger) WithOutput(w io.Writer) Logger {
	l.out = w
	return l
}

func (l Logger) WithDateFormat(format string) Logger {
	l.dateFormat = format
	return l
}

func (l Logger) WithDatetimeFormat(format string) Logger {
	l.datetimeFormat = format
	return l
}

func (l Logger) WithTimeFormat(format string) Logger {
	l.timeFormat = format
	return l
}

func (l Logger) WithStyling(styling bool) Logger {
	l.styling = styling
	return l
}

func (l Logger) formatMessage(level Level, message string, withStyling bool) string {
	message = l.addVariables(level, message)
	if withStyling && ((l.styling && terminal.AreColorsSupported()) || l.forceStyling) {
		return l.addStyling(message)
	}
	return l.removeStyling(message)
}
