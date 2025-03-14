package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type LogLevel string

const (
	INFO  LogLevel = "INFO"
	ERROR LogLevel = "ERROR"
	DEBUG LogLevel = "DEBUG"
)

type Logger struct {
	mu       sync.Mutex
	logFile  *os.File
	filename string
}


var GlobalLogger *Logger

func InitLogger(filename string) {
	GlobalLogger = NewLogger(filename)
}



func NewLogger(filename string) *Logger {
	dir := "logs"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("failed to create log directory: %v", err))
	}


	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return &Logger{
		logFile:  file,
		filename: filename,
	}
}
func (l *Logger) Log(level LogLevel, message string, fields ...map[string]interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	logEntry := map[string]interface{}{
		"time":    time.Now().Format(time.RFC3339),
		"level":   level,
		"message": message,
	}

	if len(fields) > 0 && fields[0] != nil {
		for k, v := range fields[0] {
			logEntry[k] = v
		}
	}

	jsonData, _ := json.Marshal(logEntry)
	fmt.Fprintln(l.logFile, string(jsonData))
}

func (l *Logger) Close() {
	l.logFile.Close()
}
