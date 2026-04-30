package tangra

import (
	"encoding/json"
	"io"
	"time"

	"github.com/kkrypt0nn/tangra/v2/terminal"
)

type Logger struct {
	dateFormat     string
	datetimeFormat string
	forceStyling   bool
	formatMode     FormatMode
	logFile        io.Writer
	loggingLevel   Level
	output         io.Writer
	prefix         string
	styling        bool
	timeFormat     string
}

func NewLogger() Logger {
	return Logger{
		dateFormat:     "Jan 02, 2006",
		datetimeFormat: "Jan 02, 2006 15:04:05",
		formatMode:     HumanFormat,
		loggingLevel:   NONE,
		prefix:         "${datetime} ${level:color}${level:name}${reset}: ",
		styling:        true,
		timeFormat:     "15:04:05",
	}
}

func (l Logger) WithDateFormat(format string) Logger {
	l.dateFormat = format
	return l
}

func (l Logger) WithDatetimeFormat(format string) Logger {
	l.datetimeFormat = format
	return l
}

func (l Logger) WithForceStyling(force bool) Logger {
	l.forceStyling = force
	return l
}

func (l Logger) WithFormatMode(mode FormatMode) Logger {
	l.formatMode = mode
	return l
}

func (l Logger) WithLogFile(w io.Writer) Logger {
	l.logFile = w
	return l
}

func (l Logger) WithLoggingLevel(loggingLevel Level) Logger {
	l.loggingLevel = loggingLevel
	return l
}

func (l Logger) WithOutput(w io.Writer) Logger {
	l.output = w
	return l
}

func (l Logger) WithPrefix(prefix string) Logger {
	l.prefix = prefix
	return l
}

func (l Logger) WithStyling(styling bool) Logger {
	l.styling = styling
	return l
}

func (l Logger) WithTimeFormat(format string) Logger {
	l.timeFormat = format
	return l
}

func (l Logger) formatMessage(level Level, message string, withStyling bool) string {
	message = l.addVariables(level, message)
	if withStyling && ((l.styling && terminal.AreColorsSupported()) || l.forceStyling) {
		return l.addStyling(message)
	}
	return l.removeStyling(message)
}

func (l Logger) render(level Level, message string, withStyling bool) string {
	if l.formatMode == JSONFormat {
		return l.formatJSON(level, message)
	}
	msg := l.prefix + message + terminal.RESET
	return l.formatMessage(level, msg, withStyling)
}

type logEntry struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

func (l Logger) formatJSON(level Level, msg string) string {
	entry := logEntry{
		Time:    time.Now().Format(time.RFC3339),
		Level:   GetLevelName(level),
		Message: msg,
	}

	b, err := json.Marshal(entry)
	if err != nil {
		return `{"error":"json marshal failed"}`
	}
	return string(b)
}
