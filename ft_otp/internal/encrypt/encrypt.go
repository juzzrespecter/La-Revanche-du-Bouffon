package encrypt

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"ft_otp/internal/utils"
	"os"
)

func encryptKey(key, passwd []byte) ([]byte, error) {
	c, err := aes.NewCipher(passwd)
	if err != nil {
		return nil, err
	}
	padding := aes.BlockSize - len(key)%aes.BlockSize
	cypher := make([]byte, len(key)+padding)
	cypher = append(cypher, make([]byte, padding)...)
	c.Encrypt()

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
	passwd, err := utils.ReadInput()
	if err != nil {
		return fmt.Errorf("Encrypt Key: %w", err)
	}
	encryptedKey, err := encryptKey([]byte(key), passwd)
	if err != nil {
		return fmt.Errorf("Encrypt Key: %w", err)
	}
	if n, err := fd.Write(encryptedKey); err != nil || n != len(encryptedKey) {

	}
	return nil
}

// Bla bla bla test
func DecryptKey(key string) ([]byte, error) {
	encKey, err := os.ReadFile(key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	passwd, err := utils.ReadInput()
	if err != nil {
		return nil, fmt.Errorf("Decrypt Key: %w", err)
	}
}
