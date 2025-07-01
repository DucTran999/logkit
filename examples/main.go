package main

import (
	"log"

	"github.com/DucTran999/logkit"
)

func main() {
	conf := logkit.Config{
		Environment: logkit.Production,
		LogToFile:   true,
		FilePath:    "logs/app.log",
	}

	logInst, err := logkit.NewLogger(conf)
	if err != nil {
		log.Fatalln("Init logger ERR", err)
	}
	defer logInst.Sync()

	// Log at different levels
	logInst.Debug("Debug log")
	logInst.Info("Info log")
	logInst.Warn("Warning log")
	logInst.Error("Error log")
	// Note: Fatal() exits the program and Panic() causes a panic
	// Use these methods only when appropriate for your application
}
