package totp

import (
	"encoding/binary"
	"fmt"
	hmac "ft_otp/internal/HMAC"
	"time"
)

func dt(K []byte) uint32 {
	offsetBits := K[:4]
	offset := binary.LittleEndian.Uint32(offsetBits)
	// return
}

func TOTP(K []byte) {
	fmt.Println("key: ", K)

	t := time.Now().Unix() / 30
	T := []byte{
		uint8(t),
		uint8(t << 1),
		uint8(t << 2),
		uint8(t << 3),
		uint8(t << 4),
		uint8(t << 5),
	}
	fmt.Print(T)
	hmac := hmac.HMAC(K, T)
	fmt.Println(hmac)

	//		Let Snum  = StToNum(Sbits)   // Convert S to a number in
	//	                                   0...2^{31}-1
	//	  Return D = Snum mod 10^Digit //  D is a number in the range
	//	                                   0...10^{Digit}-1
}
