package ui

import (
	"fmt"
	"log"
	"os"
)

var logFile *os.File
var logger *log.Logger

func InitializeLogger() error {
	logFileName := "logs.log"
	var err error
	logFile, err = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	logger = log.New(logFile, "FedoraInstaller: ", log.LstdFlags)
	logger.Println("--- Installer started ---") // Log start of installer session
	return nil
}

func CloseLogger() {
	if logFile != nil {
		logger.Println("--- Installer finished ---") // Log end of installer session
		logFile.Close()
	}
}

func Log(message string) {
	if logger != nil {
		logger.Println(message)
	} else {
		fmt.Println("Logger not initialized, message:", message) // Fallback if logger not setup
	}
}

func LogError(message string, err error) {
	if logger != nil {
		logger.Printf("ERROR: %s - %v\n", message, err)
	} else {
		fmt.Printf("ERROR: %s - %v\n", message, err) // Fallback if logger not setup
	}
}

func Logf(format string, v ...any) {
	if logger != nil {
		logger.Printf(format, v...)
	} else {
		fmt.Printf(format, v...) // Fallback if logger not setup
	}
}
