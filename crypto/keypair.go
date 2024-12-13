package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"my_projectx/types"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func (pk PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, pk.key, data)
	if err != nil {
		panic(err)
	}
	return &Signature{
		r, s,
	}, nil
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return PrivateKey{key}
}

func (pk PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		key: &pk.key.PublicKey,
	}
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (pk PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(pk.key.Curve, pk.key.X, pk.key.Y)
}

// Address 为什么饶了很远的路,还写了一个AddressFromBytes方法?
// 要返回一个Address即[20]uint8, uint8=byte, 即[]byte类型->[20]uint8
func (pk PublicKey) Address() types.Address {

	// 比特币地址是通过公钥经过 SHA-256、RIPEMD-160 哈希，再加上网络前缀、校验和并进行 Base58 编码生成的。
	// 这种设计保证了地址的唯一性和完整性，同时便于验证与传输。
	// 但是当前演示的仅为简单的实现
	h := sha256.Sum256(pk.ToSlice())
	return types.AddressFromBytes(h[:20])
}

type Signature struct {
	r, s *big.Int
}

func (sig Signature) Verify(publicKey PublicKey, data []byte) bool {
	return ecdsa.Verify(publicKey.key, data, sig.r, sig.s)
}
