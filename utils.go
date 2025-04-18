package main

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func init() {
	log.New()
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		PadLevelText:  true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return "", frame.File + ":" + strconv.Itoa(frame.Line)
		},
	})
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
}

// logrusWriter implements io.Writer for Gin logs and maps Gin log levels to logrus levels.
type logrusWriter struct{}

func (w logrusWriter) Write(p []byte) (n int, err error) {
	msg := string(p)
	trimmed := strings.TrimSpace(msg)
	// Map Gin log level prefix to logrus
	switch {
	case strings.HasPrefix(trimmed, "[INFO]"):
		log.Info(strings.TrimPrefix(trimmed, "[INFO]"))
	case strings.HasPrefix(trimmed, "[WARN]"):
		log.Warn(strings.TrimPrefix(trimmed, "[WARN]"))
	case strings.HasPrefix(trimmed, "[WARNING]"):
		log.Warn(strings.TrimPrefix(trimmed, "[WARNING]"))
	case strings.HasPrefix(trimmed, "[ERROR]"):
		log.Error(strings.TrimPrefix(trimmed, "[ERROR]"))
	case strings.HasPrefix(trimmed, "[FATAL]"):
		log.Fatal(strings.TrimPrefix(trimmed, "[FATAL]"))
	case strings.HasPrefix(trimmed, "[PANIC]"):
		log.Panic(strings.TrimPrefix(trimmed, "[PANIC]"))
	case strings.HasPrefix(trimmed, "[DEBUG]"):
		log.Debug(strings.TrimPrefix(trimmed, "[DEBUG]"))
	default:
		log.Debug(trimmed)
	}
	return len(p), nil
}
