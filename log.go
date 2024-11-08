package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	filePath string
}

// NewLogger creates a new logger with the specified directory and filename
func NewLogger(directory, filename string) (*Logger, error) {
	// Creates the directory if it doesn't exist
	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		return nil, fmt.Errorf("error creating directory: %w", err)
	}

	// Constructs the full file path
	filePath := filepath.Join(directory, filename)

	// Creates the log file (overwrites if it already exists)
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("error creating log file: %w", err)
	}
	defer file.Close()

	fmt.Println("Log file created:", filePath)
	return &Logger{filePath: filePath}, nil
}

// AppendLog adds a new message to the log file with a timestamp
func (l *Logger) AppendLog(message string) error {
	// Opens the log file in append mode
	file, err := os.OpenFile(l.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening log file: %w", err)
	}
	defer file.Close()

	// Creates the log entry string with a timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("%s - %s\n", timestamp, message)

	// Writes the log entry to the file
	if _, err := file.WriteString(logEntry); err != nil {
		return fmt.Errorf("error writing to log file: %w", err)
	}

	fmt.Println("Log added:", logEntry)
	return nil
}
