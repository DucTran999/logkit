package logkit

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// Infof logs a formatted message at InfoLevel.
func (l *logger) Infof(template string, args ...any) {
	l.sugarLogger.Infof(template, args...)
}

// Warnf logs a formatted message at WarnLevel.
func (l *logger) Warnf(template string, args ...any) {
	l.sugarLogger.Warnf(template, args...)
}

// Errorf logs a formatted message at ErrorLevel and includes a stack trace.
func (l *logger) Errorf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.logWithStack(zapcore.ErrorLevel, msg)
}

// Debugf logs a formatted message at DebugLevel.
func (l *logger) Debugf(template string, args ...any) {
	l.sugarLogger.Debugf(template, args...)
}

// Panicf logs a formatted message at PanicLevel and panics. Includes a stack trace.
func (l *logger) Panicf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.logWithStack(zapcore.PanicLevel, msg)
}

// DPanicf logs a formatted message at DPanicLevel. Includes a stack trace.
// Panics in development mode.
func (l *logger) DPanicf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.logWithStack(zapcore.DPanicLevel, msg)
}

// Fatalf logs a formatted message at FatalLevel and exits the program. Includes a stack trace.
func (l *logger) Fatalf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.logWithStack(zapcore.FatalLevel, msg)
}
