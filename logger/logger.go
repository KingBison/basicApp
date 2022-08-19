package logger

import (
	"fmt"
)

func Info(msg string) {
	fmt.Println("INFO: " + msg)
}

func Debug(msg string) {
	fmt.Println("DEBUG: " + msg)
}

func Error(msg string) {
	fmt.Println("ERROR: " + msg)
}
