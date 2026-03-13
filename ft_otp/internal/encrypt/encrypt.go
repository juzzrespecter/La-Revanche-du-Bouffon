package encrypt

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"ft_otp/internal/utils"
	"os"
)

func encryptKey(key string, passwd string) {

}

func EncryptKey(key string) error {
	const (
		keyFile = "ft_otp.key"
	)
	if len(key) < 64 {
		return utils.ErrKeyLength
	}
	if _, err := hex.DecodeString(key); err != nil {
		fmt.Println(err)
		return utils.ErrKeyEncode
	}
	fd, err := os.OpenFile(keyFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		// error de permisos, not a file, etc
		fmt.Println(err)
		os.Exit(1)
	}
	defer fd.Close()
	rd := bufio.NewReader(os.Stdin)
	fmt.Println("Introduce encryption key:")
	passwd, err := rd.ReadString(byte('\n'))
	if err != nil {
		// broken pipe

	}
	fmt.Println(passwd)

	return nil
}
