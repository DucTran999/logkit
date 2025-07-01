package logkit_test

import (
	"context"
	"sync"
	"testing"

	"github.com/DucTran999/logkit"
	"github.com/stretchr/testify/require"
)

func TestLogFromMultiGoRoutine(t *testing.T) {
	conf := logkit.Config{
		Environment: logkit.Production,
		LogToFile:   true,
		FilePath:    "logs/app.log",
	}

	logInst, err := logkit.NewLogger(conf)
	require.NoError(t, err)
	defer func() { _ = logInst.Sync() }()

	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func(logkit.ILogger) {
			ctx := context.Background()
			newCtx := context.WithValue(ctx, logkit.RequestIDKeyCtx, "123456")

			defer wg.Done()
			logInst.FromContext(newCtx).Error("example error log")
			logInst.FromContext(newCtx).Info("example info log")
			logInst.FromContext(newCtx).Debug("example error log")
			logInst.FromContext(newCtx).Warn("example warn log")

			logInst.FromContext(context.Background()).Infof("example info log %s", "test")
			logInst.FromContext(context.Background()).Debugf("example debug log %s", "test")
			logInst.FromContext(context.Background()).Errorf("example error log %s", "test")
			logInst.FromContext(context.Background()).Warnf("example warn log %s", "test")
		}(logInst)
	}
	wg.Wait()
}
