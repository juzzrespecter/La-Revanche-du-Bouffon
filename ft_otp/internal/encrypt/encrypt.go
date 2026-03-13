package encrypt

import (
	"bufio"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"ft_otp/internal/utils"
	"os"
)

func encryptKey(key string, passwd []byte) ([]byte, error) {
	c, err := aes.NewCipher(passwd)
	if err != nil {
		return nil, err
	}
	padding := aes.BlockSize - len(key)%aes.BlockSize
	bytesKey = make([]byte, len(key)+padding)
}

func decryptKey(cypher, passwd []byte) (string, error) {

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
		return fmt.Errorf("Encrypt Key: %w", err)
	}
	defer fd.Close()
	rd := bufio.NewReader(os.Stdin)
	fmt.Println("Introduce encryption key:")
	passwd, err := rd.ReadBytes(byte('\n'))
	if err != nil {
		return fmt.Errorf("Encrypt Key: %w", err)
	}
	encryptedKey, err := encryptKey(key, passwd)
	if err != nil {
		return fmt.Errorf("Encrypt Key: %w", err)
	}
	if n, err := fd.Write(encryptedKey); err != nil || n != len(encryptedKey) {

	}
	return nil
}

func DecryptKey() {

}
