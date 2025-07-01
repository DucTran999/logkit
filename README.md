# Logkit

[![Go Report Card](https://goreportcard.com/badge/github.com/DucTran999/logkit)](https://goreportcard.com/report/github.com/DucTran999/logkit)
[![Go](https://img.shields.io/badge/Go-1.23-blue?logo=go)](https://golang.org)
[![CI](https://github.com/DucTran999/logkit/actions/workflows/ci.yml/badge.svg)](https://github.com/DucTran999/logkit/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/DucTran999/logkit/branch/master/graph/badge.svg)](https://codecov.io/gh/DucTran999/logkit)
[![License](https://img.shields.io/github/license/DucTran999/logkit)](LICENSE)

This package wraps Uber's zap library to simplify configuration and logging in Go applications.

## Installation

```bash
go get github.com/DucTran999/logkit
```

## 🚀 Features

- 🪵 Configurable log levels: Debug, Info, Warn, Error, Fatal, Panic, Dpanic
- 🎛️ JSON or human-readable log formats (console).
- 🔁 Flexible output targets: stdout, file, or any io.Writer.
- 🕒 RFC3339 timestamp formatting.
- 🛠️ Development and production-friendly logging.
- 🧵 Log level filtering and structured context-aware logging.

---

## 📦 Usage

Here’s an example of how to use the logkit in your application.

```go
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
```

## 🧪 Testing

```sh
# run: go test ./... if you not install task
task test
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Submit a pull request

---

## 📄 License

This project is licensed under the MIT License.

---
