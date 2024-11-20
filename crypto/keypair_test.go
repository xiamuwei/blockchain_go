package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	address := publicKey.Address()
	fmt.Println(address)

	msg := []byte("hello blockchain")
	sign, err := privateKey.Sign(msg)
	assert.Nil(t, err)
	fmt.Println(sign)

	res := sign.Verify(publicKey, msg)
	assert.True(t, res)

}
