package logkit_test

import (
	"testing"

	"github.com/DucTran999/logkit"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestLogInfo(t *testing.T) {
	conf := logkit.Config{
		Environment: logkit.Development,
	}

	logInst, err := logkit.NewLogger(conf)
	require.NoError(t, err)
	defer func() { _ = logInst.Sync() }()

	logInst.Info("example info log")
}

func TestLogWarn(t *testing.T) {
	conf := logkit.Config{
		Environment: logkit.Development,
	}

	logInst, err := logkit.NewLogger(conf)
	require.NoError(t, err)
	defer func() { _ = logInst.Sync() }()

	logInst.Warn("example warn log")
}

func TestLogError(t *testing.T) {
	conf := logkit.Config{
		Environment: logkit.Development,
	}

	logInst, err := logkit.NewLogger(conf)
	require.NoError(t, err)
	defer func() { _ = logInst.Sync() }()

	logInst.Error("example error log")
}

func TestLogDebug(t *testing.T) {
	conf := logkit.Config{
		Environment: logkit.Development,
	}

	logInst, err := logkit.NewLogger(conf)
	require.NoError(t, err)
	defer func() { _ = logInst.Sync() }()

	logInst.Debug("example debug log")
}

func TestPanic(t *testing.T) {
	conf := logkit.Config{
		Environment: logkit.Production,
	}

	logInst, err := logkit.NewLogger(conf)
	require.NoError(t, err)

	panicOccurred := false
	defer func() {
		if r := recover(); r != nil {
			logInst.Error("example panic log", zap.Any("stack", r))
			panicOccurred = true
		}
		_ = logInst.Sync()
		require.True(t, panicOccurred, "Expected panic occur in prod environment")
	}()

	logInst.Panic("example panic log", zap.String("stack", "stack trace"))
}
