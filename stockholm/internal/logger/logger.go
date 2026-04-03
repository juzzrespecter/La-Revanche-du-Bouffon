package logger

import (
	"fmt"
	"os"
	"sync"
)

const (
	DEBUG   string = "\033[1;35m"
	SUCCESS string = "\033[1;32m"
	INFO    string = "\033[1;34m"
	WARNING string = "\033[1;33m"
	ERROR   string = "\033[0;31m"
	RESET   string = "\033[0m"
)

var lock = &sync.Mutex{}

type logger struct {
	silent bool
}

var Logger *logger

func NewLogger(silent bool) *logger {
	lock.Lock()
	defer lock.Unlock()
	if Logger == nil {
		return nil
	}
	Logger = &logger{
		silent: silent,
	}
	return Logger
}

func (l *logger) Debug(msg string) {
	if !l.silent {
		fmt.Fprintf(os.Stderr, "%s [DEBUG]   %s %s\n", DEBUG, RESET, msg)
	}
}

func (l *logger) Success(msg string) {
	if !l.silent {
		fmt.Fprintf(os.Stderr, "%s [SUCCESS] %s %s\n", SUCCESS, RESET, msg)
	}
}

func (l *logger) Info(msg string) {
	if !l.silent {
		fmt.Fprintf(os.Stderr, "%s [INFO]    %s %s\n", INFO, RESET, msg)
	}
}

func (l *logger) Warning(msg string) {
	if !l.silent {
		fmt.Fprintf(os.Stderr, "%s [WARNING] %s %s\n", WARNING, RESET, msg)
	}
}

func (l *logger) Error(msg string) {
	if !l.silent {
		fmt.Fprintf(os.Stderr, "%s [ERROR]	 %s %s\n", ERROR, RESET, msg)
	}
}
