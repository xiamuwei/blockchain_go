package types

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Hash [32]uint8

func (h Hash) IsZero() bool {
	for _, value := range h {
		if value != 0 {
			return false
		}
	}
	return true
}
func (h Hash) ToSlice() []byte {
	buf := make([]byte, 32)
	for i := 0; i < 32; i++ {
		buf[i] = h[i]
	}
	return buf
}

func (h Hash) String() string {
	return hex.EncodeToString(h.ToSlice())
}

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		panic(fmt.Sprintf("The length of bytes %d should be 32", len(b)))
	}
	var value [32]uint8
	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}
	return Hash(value)
}

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
