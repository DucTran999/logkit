//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/DucTran999/logkit"
	"go.uber.org/zap"
)

func main() {
	loggerInst, err := logkit.NewLogger(logkit.Config{
		Environment: logkit.Production,
		LogToFile:   false,
		FilePath:    "logs/app.log",
	})
	if err != nil {
		log.Fatalln("failed to initialize logger:", err)
	}
	defer loggerInst.Sync()

	loggerInst.Error("example error log", zap.Int("user_id", 166))
}
