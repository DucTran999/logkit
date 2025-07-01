package logkit_test

import (
	"errors"
	"testing"

	"github.com/DucTran999/logkit"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestPanicf(t *testing.T) {
	conf := logkit.Config{
		Environment: logkit.Staging,
	}

	logInst, err := logkit.NewLogger(conf)
	if err != nil {
		t.Fatalf("Init logkit ERR=%v", err)
	}

	panicOccurred := false
	defer func() {
		if r := recover(); r != nil {
			logInst.Error("example panic log", zap.Any("stack", r))
			panicOccurred = true
		}
		_ = logInst.Sync()
		require.True(t, panicOccurred, "Expected panic did not occur in Staging environment")
	}()

	logInst.Panicf("example panic log %v", errors.New("panic test"))
}

func TestDPanicInDevelopment(t *testing.T) {
	conf := logkit.Config{
		Environment: logkit.Development,
	}

	logInst, err := logkit.NewLogger(conf)
	if err != nil {
		t.Fatalf("Init logkit ERR=%v", err)
	}

	panicOccurred := false
	defer func() {
		if r := recover(); r != nil {
			logInst.Error("example dpanic log", zap.Any("stack", r))
			panicOccurred = true
		}
		_ = logInst.Sync()
		require.True(t, panicOccurred, "Expected panic did not occur in Staging environment")
	}()

	logInst.DPanic("example dpanic log")
}

func TestDPanicNotInDevelopment(t *testing.T) {
	conf := logkit.Config{
		Environment: logkit.Production,
	}

	logInst, err := logkit.NewLogger(conf)
	if err != nil {
		t.Fatalf("Init logkit ERR=%v", err)
	}

	panicOccurred := false
	defer func() {
		if r := recover(); r != nil {
			logInst.Error("example dpanic log", zap.Any("stack", r))
			panicOccurred = true
		}
		_ = logInst.Sync()
		require.False(t, panicOccurred, "Expected panic did not occur in prod environment")
	}()

	logInst.DPanic("example dpanic log")
}

func TestDPanicfInDevelopment(t *testing.T) {
	conf := logkit.Config{
		Environment: logkit.Development,
	}

	logInst, err := logkit.NewLogger(conf)
	require.NoError(t, err)

	panicOccurred := false
	defer func() {
		if r := recover(); r != nil {
			logInst.Error("example dpanicf log", zap.Any("stack", r))
			panicOccurred = true
		}
		_ = logInst.Sync()
		require.True(t, panicOccurred, "Expected panic occur in Development environment")
	}()

	logInst.DPanicf("example dpanicf log err:%v", errors.New("panic test"))
}

func TestDPanicfNotInDevelopment(t *testing.T) {
	conf := logkit.Config{
		Environment: logkit.Staging,
	}

	logInst, err := logkit.NewLogger(conf)
	require.NoError(t, err)

	panicOccurred := false
	defer func() {
		if r := recover(); r != nil {
			logInst.Error("example dpanicf log", zap.Any("stack", r))
			panicOccurred = true
		}
		_ = logInst.Sync()
		require.False(t, panicOccurred, "Expected panic not occur in stage environment")
	}()

	logInst.DPanicf("example dpanicf log err:%v", errors.New("panic test"))
}
