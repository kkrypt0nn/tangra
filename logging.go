package tangra

import (
	"fmt"
	"os"
)

func (l Logger) log(level Level, message string) {
	if level < l.loggingLevel {
		return
	}

	out := l.output
	if out == nil {
		out = os.Stdout
	}

	rendered := l.render(level, message, l.styling)

	_, _ = fmt.Fprintln(out, rendered)
	if l.logFile != nil {
		_, _ = fmt.Fprintln(l.logFile, l.removeStyling(rendered))
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
func (l Logger) Fatal(msg string) {
	l.log(FATAL, msg)
	os.Exit(1)
}
func (l Logger) Fatalf(format string, args ...any) {
	l.log(FATAL, fmt.Sprintf(format, args...))
	os.Exit(1)
}
func (l Logger) Trace(msg string)                  { l.log(TRACE, msg) }
func (l Logger) Tracef(format string, args ...any) { l.log(TRACE, fmt.Sprintf(format, args...)) }
