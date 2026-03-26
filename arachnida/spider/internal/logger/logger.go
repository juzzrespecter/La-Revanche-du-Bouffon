package logger

import (
	"fmt"
	"os"
)

const (
	DEBUG   string = "\033[1;35m"
	INFO    string = "\033[1;34m"
	WARNING string = "\033[1;33m"
	ERROR   string = "\033[0;31m"
	RESET   string = "\033[0m"
)

func Debug(msg string) {
	fmt.Fprintf(os.Stderr, "%s [DEBUG]   %s %s\n", DEBUG, RESET, msg)
}

func Info(msg string) {
	fmt.Fprintf(os.Stderr, "%s [INFO]    %s %s\n", INFO, RESET, msg)
}

func Warning(msg string) {
	fmt.Fprintf(os.Stderr, "%s [WARNING] %s %s\n", WARNING, RESET, msg)
}

func Error(msg string) {
	fmt.Fprintf(os.Stderr, "%s [ERROR]	 %s %s\n", ERROR, RESET, msg)
}
