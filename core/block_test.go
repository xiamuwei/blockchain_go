package core

import (
	"blockchain_go/types"
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PreBlock:  types.RandomHash(),
		TimeStamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     999999,
	}
	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))
	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}

func TestBlock_Encode_Decode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PreBlock:  types.RandomHash(),
			TimeStamp: time.Now().UnixNano(),
			Height:    10,
		},
		Transactions: []Transaction{},
	}
	h := b.Hash()
	fmt.Println("hash = ", h)
	assert.False(t, h.IsZero())
}
