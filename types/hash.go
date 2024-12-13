package types

import (
	"crypto/rand"
	"encoding/hex"
)

// Hash SHA256生成 256bit=32byte=32uint8
type Hash [32]uint8

func (hash Hash) IsZero() bool {
	for _, b := range hash {
		if b != 0 {
			return false
		}
	}
	return true
}

// SHA-256（Secure Hash Algorithm 256 bit）是一种加密哈希函数
// 它会产生一个256位（32字节）的哈希值
// 这个哈希值通常表示为一个64位的十六进制数
func (hash Hash) String() string {
	return hex.EncodeToString(hash.HashToBytes())
}

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		panic("wrong hash size")
	}
	return Hash(b)
}
func (hash Hash) HashToBytes() []byte {
	//b := make([]byte, 32)
	//for i, h := range hash {
	//	b[i] = h
	//}
	//return b
	return hash[:]
}

func RandomBytes(size int) []byte {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func RandomHash() Hash {
	//return Hash(RandomBytes(32))
	return HashFromBytes(RandomBytes(32))
}
