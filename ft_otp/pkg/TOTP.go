package totp

import (
	"fmt"
	"time"
)

func TOTP(key []byte) {
	fmt.Println("key: ", key)

	T := time.Now().Unix() / 30
	fmt.Print(T)
}
