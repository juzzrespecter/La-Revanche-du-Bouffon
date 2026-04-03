package totp

import (
	"encoding/binary"
	hmac "ft_otp/internal/HMAC"
	"strconv"
	"time"
)

func dt(H []byte) int {
	offset := H[19] & 0xf
	P := (int(H[offset])&0x7f)<<24 |
		(int(H[offset+1]))<<16 |
		(int(H[offset+2]))<<8 |
		(int(H[offset+3]))
	return P % 1000000
}

func TOTP(K []byte) string {
	t := time.Now().Unix() / 30
	T := make([]byte, 8)
	binary.BigEndian.PutUint64(T, uint64(t))
	hmac := hmac.HMAC(K, T)
	code := dt(hmac)
	return strconv.Itoa(code)
}
