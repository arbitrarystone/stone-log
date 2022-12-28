package main

import (
	slog "github.com/arbitrarystone/stone-log"
	"go.uber.org/zap"
)

func main() {
	ops := slog.Options{
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Out:        "file",
		FileName:   "./slog-test.log",
		LogLevel:   "debug",
	}
	slog.Init(&ops)
	slog.Info("Success..",
		zap.String("statusCode", "x200"),
		zap.String("url", "url-test"))
}
