package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func main() {
	fmt.Println("it's loggin!")

	logger_config := zap.NewProductionConfig()
	logger_config.Encoding = "console"
	logger_config.OutputPaths = []string{
		"/tmp/test.log",
		"stdout",
	}
	logger, err := logger_config.Build()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar().With(
		zap.String("ProgramName", "logging"),
	)
	sugar.Infof("Failed to fetch URL: %s, attempt %d, backoff: %s", "localhost", 3, time.Second)
}
