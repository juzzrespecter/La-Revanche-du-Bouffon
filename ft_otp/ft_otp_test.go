package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/hex"
	"testing"
	"time"

	h "ft_otp/internal/HMAC"
	to "ft_otp/pkg/totp"

	"github.com/pquerna/otp/totp"
)

// HMAC generation using same key
func TestHMAC(t *testing.T) {
	key := make([]byte, 64)
	msg := make([]byte, 8)

	rand.Read(msg)
	rand.Read(key)
	key = []byte(hex.EncodeToString(key))

	ourMac := h.HMAC(key, msg)
	theirMac := hmac.New(sha1.New, key)
	theirMac.Write(msg)
	expectedMac := theirMac.Sum(nil)
	if !bytes.Equal(ourMac, expectedMac) {
		t.Errorf("Expected MAC value: %v, got %v\n", expectedMac, ourMac)
	}
}

// TOTP code generation
func TestTOTP(t *testing.T) {
	K := make([]byte, 64)
	rand.Read(K)

	K_2 := base32.StdEncoding.EncodeToString(K)
	theirOTP, _ := totp.GenerateCode(K_2, time.Now())
	ourOTP := to.TOTP(K)
	if theirOTP != ourOTP {
		t.Errorf("Expected OTP value: %s, got %s\n", theirOTP, ourOTP)
	}

}
