package types

import "encoding/hex"

// Address 为什么定义为20长度?
// BTC就是20长度,
type Address [20]uint8

func (addr Address) ToSlice() []byte {
	return addr[:20]
}

func AddressFromBytes(addrByte []byte) Address {
	var b [20]uint8
	copy(b[:], addrByte[:20])
	return b
}

func (addr Address) String() string {
	return hex.EncodeToString(addr.ToSlice())
}
