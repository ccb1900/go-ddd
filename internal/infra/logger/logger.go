package logger

import (
	"goddd/pkg/config"
	"io"
	"log/slog"
	"os"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)
var logOnce sync.Once
func NewLogger(cfgApp config.AppConfig) *slog.Logger {
	cfg := cfgApp.Log
    // 日志文件切割器
    fileWriter := &lumberjack.Logger{
        Filename:   cfg.Filename,
        MaxSize:    cfg.MaxSize,
        MaxAge:     cfg.MaxAge,
        MaxBackups: cfg.MaxBackups,
        Compress:   cfg.Compress,
    }

    // 控制台输出
    consoleWriter := os.Stdout

    // 多路输出
    multiWriter := io.MultiWriter(fileWriter, consoleWriter)

    // 日志级别
    var level slog.Level
    switch cfg.Level {
    case "debug":
        level = slog.LevelDebug
    case "info":
        level = slog.LevelInfo
    case "warn":
        level = slog.LevelWarn
    case "error":
        level = slog.LevelError
    default:
        level = slog.LevelInfo
    }

    handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
        Level: level,
    })

    logger := slog.New(handler)

    // 设置为默认 logger
    slog.SetDefault(logger)

    return logger
}