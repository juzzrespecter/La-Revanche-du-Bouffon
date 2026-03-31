package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"ft_otp/internal/utils"
	"os"
)

func encryptKey(plaintext, key []byte) ([]byte, error) {
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	ciphertext := make([]byte, len(plaintext)+padding+aes.BlockSize)
	plaintext = append(plaintext, make([]byte, padding)...)

	iv := ciphertext[:aes.BlockSize]
	rand.Read(iv)
	hash := sha512.Sum512(key)
	b, err := aes.NewCipher(hash[:32])
	if err != nil {
		return nil, err
	}

	m := cipher.NewCBCEncrypter(b, iv)
	m.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func decryptKey(ciphertext, key []byte) ([]byte, error) {
	hash := sha512.Sum512(key)
	b, err := aes.NewCipher(hash[:32])
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	m := cipher.NewCBCDecrypter(b, iv)
	m.CryptBlocks(ciphertext, ciphertext)
	return ciphertext, nil
}

// AES CBC
func EncryptKey(key string) error {
	const (
		keyFile = "ft_otp.key"
	)

	if len(key) < 64 {
		return utils.ErrKeyLength
	}
	dec, err := hex.DecodeString(key)
	if err != nil {
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
	encryptedKey, err := encryptKey(dec, passwd)
	if err != nil {
		return fmt.Errorf("Encrypt Key: %w", err)
	}
	if n, err := fd.Write(encryptedKey); err != nil || n != len(encryptedKey) {
		return fmt.Errorf("Encrypt Key: %w", err)
	}
	return nil
}

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
	plaintext, err := decryptKey(encKey, passwd)
	if err != nil {
		return nil, fmt.Errorf("Decrypt key: %w", err)
	}
	return plaintext, nil
}
