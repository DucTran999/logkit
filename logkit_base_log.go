package logkit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Info logs a message at InfoLevel.
func (l *logger) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel.
func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.zapLogger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel and includes a stack trace.
func (l *logger) Error(msg string, fields ...zap.Field) {
	l.logWithStack(zapcore.ErrorLevel, msg, fields...)
}

// Debug logs a message at DebugLevel.
func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

// Panic logs a message at PanicLevel and then panics. Includes a stack trace.
func (l *logger) Panic(msg string, fields ...zap.Field) {
	l.logWithStack(zapcore.PanicLevel, msg, fields...)
}

// DPanic logs a message at DPanicLevel.
// In development, it panics. In production, it only logs. Includes a stack trace.
func (l *logger) DPanic(msg string, fields ...zap.Field) {
	l.logWithStack(zapcore.DPanicLevel, msg, fields...)
}

// Fatal logs a message at FatalLevel and then calls os.Exit(1). Includes a stack trace.
func (l *logger) Fatal(msg string, fields ...zap.Field) {
	l.logWithStack(zapcore.FatalLevel, msg, fields...)
}
