package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	fmt.Println("it's loggin!")

	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Encoding = "console" // Use console encoding
	loggerConfig.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:       "msg",
		ConsoleSeparator: " ", // Tab-separated fields
	}

	// Output logs to both console and a file
	loggerConfig.OutputPaths = []string{
		"/tmp/test.log", // Log to a file
		"stdout",        // Log to the console
	}

	// Build the logger
	logger, err := loggerConfig.Build()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer logger.Sync() // Flushes buffer, if any

	// Create a SugaredLogger
	sugar := logger.Sugar()

	// Log an example message
	pid := 344
	threads := 2
	sugar.Infof("[%s] [INFO] [%d] [%d] %s",
		time.Now().Format("20060102150405.000"),
		pid,     // PID
		threads, // Threads
		"Failed to fetch URL: localhost, attempt 3, backoff: 1s", // Log message
	)
}
