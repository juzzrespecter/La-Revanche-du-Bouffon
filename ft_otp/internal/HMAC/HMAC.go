package hmac

import (
	"crypto/sha1"
)

const (
	ipadByte uint8 = 0x36
	opadByte uint8 = 0x5c
	B        int   = 64
	L        int   = 20
)

func shaSum(x []byte) []byte {
	tmp := sha1.Sum(x)
	return tmp[:]
}

func setMask(c uint8) []byte {
	a := make([]byte, B)
	for i := range B {
		a[i] = c
	}
	return a[:]
}

func xor(a []byte, b []byte) []byte {
	K := make([]byte, B)
	for i := range len(a) {
		K[i] = a[i] ^ b[i]
	}
	return K
}

func HMAC(K []byte, c []byte) []byte {
	if len(K) > int(B) {
		K = shaSum(K)
	}
	if len(K) < B {
		pad := B - len(K)%B
		K = append(K, make([]byte, pad)...)
	}
	ipad := setMask(ipadByte)
	opad := setMask(opadByte)
	K1 := shaSum(append(xor(K, ipad), c...))
	K2 := shaSum(append(xor(K, opad), K1...))
	return K2
}
