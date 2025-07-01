package logkit

import (
	"context"

	"go.uber.org/zap"
)

type BaseLogger interface {
	// Debug logs a message at DebugLevel.
	Debug(msg string, fields ...zap.Field)

	// Info logs a message at InfoLevel.
	Info(msg string, fields ...zap.Field)

	// Warn logs a message at WarnLevel.
	Warn(msg string, fields ...zap.Field)

	// Error logs a message at ErrorLevel.
	Error(msg string, fields ...zap.Field)

	// Fatal logs a message at FatalLevel and exits the application.
	Fatal(msg string, fields ...zap.Field)

	// Panic logs a message at PanicLevel and panics.
	Panic(msg string, fields ...zap.Field)

	// DPanic logs a message at DPanicLevel.
	// Panics in development mode, logs in production.
	DPanic(msg string, fields ...zap.Field)
}

type FormatLogger interface {
	// Debugf logs a formatted message at DebugLevel.
	Debugf(template string, args ...any)

	// Infof logs a formatted message at InfoLevel.
	Infof(template string, args ...any)

	// Warnf logs a formatted message at WarnLevel.
	Warnf(template string, args ...any)

	// Errorf logs a formatted message at ErrorLevel.
	Errorf(template string, args ...any)

	// Fatalf logs a formatted message at FatalLevel and exits the application.
	Fatalf(template string, args ...any)

	// Panicf logs a formatted message at PanicLevel and panics.
	Panicf(template string, args ...any)

	// DPanicf logs a formatted message at DPanicLevel.
	// Panics in development mode, logs in production.
	DPanicf(template string, args ...any)
}

// ILogger defines the logging interface that supports both structured and formatted logging,
// as well as context-aware operations and logger cleanup.
type ILogger interface {
	BaseLogger

	FormatLogger

	// FromContext returns a logger enriched with values from the given context.
	FromContext(ctx context.Context) ILogger

	// Sync flushes any buffered log entries.
	Sync() error
}
