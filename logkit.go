package logkit

import (
	"context"
	"runtime"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger is an implementation of the ILogger interface,
// wrapping both a structured zap.Logger and a sugared logger for formatted output.
type logger struct {
	zapLogger   *zap.Logger
	sugarLogger *zap.SugaredLogger
}

// NewLogger creates and returns a new ILogger instance configured
// according to the provided Config.
//
// The logger uses a custom zapcore and applies sampling in production
// environments to limit the volume of logs. In development mode,
// it enables development-specific options (like DPanic triggering).
//
// Parameters:
//   - conf: Config struct defining the logging environment, output settings, etc.
//
// Returns:
//   - ILogger: a structured and sugared logger instance
//   - error: if any error occurs during logger initialization
func NewLogger(conf Config) (ILogger, error) {
	// Create the base zap core
	core := newZapCore(conf)

	// Apply sampling only for production environment
	if conf.Environment == Production {
		core = zapcore.NewSamplerWithOptions(
			core,
			time.Second, // interval
			100,         // log first 100 entries
			100,         // thereafter log zero entries within the interval
		)
	}

	zapLog := zap.New(core)
	if conf.Environment == Development {
		zapLog = zapLog.WithOptions(zap.Development())
	}

	return &logger{
		zapLogger:   zapLog,
		sugarLogger: zapLog.Sugar(),
	}, nil
}

// FromContext retrieves data from the context and returns a logger with those fields
func (l *logger) FromContext(ctx context.Context) ILogger {
	// Extract the request ID from the context
	requestID, ok := ctx.Value(RequestIDKeyCtx).(string)

	// If a non-empty request ID exists, attach it to the logger
	if ok && requestID != "" {
		newLogger := l.zapLogger.With(zap.String(string(RequestIDKeyCtx), requestID))
		return &logger{zapLogger: newLogger, sugarLogger: newLogger.Sugar()}
	}

	// Return the original logger if no valid request ID is found
	return l
}

func (l *logger) Sync() error {
	return l.zapLogger.Sync()
}

// logWithStack logs a message with optional stack trace information for error-level logs and above.
func (l *logger) logWithStack(level zapcore.Level, msg string, fields ...zap.Field) {
	// For error and above, capture caller information
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			stacktrace := zap.Dict("source",
				zap.String("path", file),
				zap.Int("line", line),
				zap.String("func", fn.Name()),
			)
			fields = append(fields, stacktrace)
		}
	}

	// Log with stack trace
	if ce := l.zapLogger.Check(level, msg); ce != nil {
		ce.Write(fields...)
	}
}
