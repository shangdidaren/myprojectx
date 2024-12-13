package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"
	"my_projectx/types"
)

type Header struct {
	Version   uint32
	PrevBlock types.Hash
	Timestamp int64
	Height    uint32
	Nonce     uint64
}

func (header *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &header.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &header.PrevBlock); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &header.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &header.Height); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, &header.Nonce)
}
func (header *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &header.Version); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &header.PrevBlock); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &header.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &header.Height); err != nil {
		return err
	}
	return binary.Read(r, binary.LittleEndian, &header.Nonce)
}

type Block struct {
	Header       Header
	Transactions []Transaction
	hash         types.Hash
}

func (block *Block) Hash() types.Hash {
	// 区块的Hash计算:根据区块的头进行计算!
	if !block.hash.IsZero() {
		return block.hash
	}
	// todo, 没动明白,但是源码这里是有改动的,因为还没有说明清楚
	buf := new(bytes.Buffer)
	if err := block.Header.EncodeBinary(buf); err != nil {
		panic(err)
	}
	block.hash = sha256.Sum256(buf.Bytes())
	return block.hash
}

func (block *Block) DecodeBinary(r io.Reader) error {
	if err := block.Header.DecodeBinary(r); err != nil {
		return err
	}
	for _, tx := range block.Transactions {
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}
	return nil
}

func (block *Block) EncodeBinary(w io.Writer) error {
	if err := block.Header.EncodeBinary(w); err != nil {
		return err
	}
	for _, tx := range block.Transactions {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}
	return nil
}
