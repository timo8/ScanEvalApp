package logging

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger      *slog.Logger
	errorLogger *slog.Logger
	once        sync.Once
)

func InitLogger() {
	once.Do(func() {
		logsDir := "logs"

		if err := os.MkdirAll(logsDir, os.ModePerm); err != nil {
			log.Fatalf("CRITICAL: Nepodarilo sa vytvoriť priečinok logs: %v", err)
		}

		appLogWriter := &lumberjack.Logger{
			Filename:   filepath.Join(logsDir, "app.log"),
			MaxSize:    5,    
			MaxBackups: 3,    
			MaxAge:     7,    
			Compress:   true, 
		}

		errorLogWriter := &lumberjack.Logger{
			Filename:   filepath.Join(logsDir, "error.log"),
			MaxSize:    5,    
			MaxBackups: 3,    
			MaxAge:     7,    
			Compress:   true, 
		}

		appHandler := slog.NewTextHandler(appLogWriter, &slog.HandlerOptions{
			Level: slog.LevelDebug, 
		})

		errorHandler := slog.NewTextHandler(errorLogWriter, &slog.HandlerOptions{
			Level: slog.LevelError, 
		})

		logger = slog.New(appHandler)
		errorLogger = slog.New(errorHandler)
	})
}

func GetLogger() *slog.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}

func GetErrorLogger() *slog.Logger {
	if errorLogger == nil {
		InitLogger()
	}
	return errorLogger
}
