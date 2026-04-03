package utils

import (
	"os"
	"os/user"
)

func GetCurrentUserHome() {
	u, err := user.Current()
	if err != nil {
		home := os.Getenv("HOME")
	}
}
