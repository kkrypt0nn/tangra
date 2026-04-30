package tangra

import (
	"fmt"
	"os"

	"github.com/kkrypt0nn/tangra/v2/terminal"
)

func (l Logger) log(level Level, message string) {
	if level < l.loggingLevel {
		return
	}

	out := l.out
	if out == nil {
		out = os.Stdout
	}

	raw := l.prefix + message + terminal.RESET
	formatted := l.formatMessage(level, raw, true)

	_, _ = fmt.Fprintln(out, formatted)

	if l.logFile != nil {
		_, _ = fmt.Fprintln(l.logFile, l.removeStyling(formatted))
	}
}

func (l Logger) Debug(msg string)                  { l.log(DEBUG, msg) }
func (l Logger) Debugf(format string, args ...any) { l.log(DEBUG, fmt.Sprintf(format, args...)) }
func (l Logger) Info(msg string)                   { l.log(INFO, msg) }
func (l Logger) Infof(format string, args ...any)  { l.log(INFO, fmt.Sprintf(format, args...)) }
func (l Logger) Warn(msg string)                   { l.log(WARN, msg) }
func (l Logger) Warnf(format string, args ...any)  { l.log(WARN, fmt.Sprintf(format, args...)) }
func (l Logger) Error(msg string)                  { l.log(ERROR, msg) }
func (l Logger) Errorf(format string, args ...any) { l.log(ERROR, fmt.Sprintf(format, args...)) }
func (l Logger) Fatal(msg string)                  { l.log(FATAL, msg) }
func (l Logger) Fatalf(format string, args ...any) { l.log(FATAL, fmt.Sprintf(format, args...)) }
func (l Logger) Trace(msg string)                  { l.log(TRACE, msg) }
func (l Logger) Tracef(format string, args ...any) { l.log(TRACE, fmt.Sprintf(format, args...)) }
