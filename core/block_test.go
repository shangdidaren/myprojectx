package core

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"my_projectx/types"
	"testing"
	"time"
)

func TestHeader_DecodeEncode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     78973434,
	}
	buf := new(bytes.Buffer)
	err := h.EncodeBinary(buf)
	assert.Nil(t, err)

	h2 := &Header{}
	err = h2.DecodeBinary(buf)
	assert.Nil(t, err)

	assert.Equal(t, h, h2)
}

func TestBlock_EncodeDecode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     78973434,
		},
		Transactions: nil,
	}
	buf := new(bytes.Buffer)
	assert.Nil(t, b.EncodeBinary(buf))
	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)
}

func TestBlockHash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     78973434,
		},
		Transactions: nil,
	}
	h := b.Hash()
	assert.False(t, h.IsZero())
}
