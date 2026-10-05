package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Block struct {
	header struct {
		prevHash   string
		merkleRoot string
		nonce      int
		difficulty int
		version    int
	}
	merkleRoot string
	merkleTree any
}

func (b *Block) hash() [32]byte {
	data, err := json.Marshal(fmt.Sprintf("%v", b.header))
	if err != nil {
		must(err)
	}
	return sha256.Sum256(data)
}

func (b *Block) mine() {
	for !b.isDifficult() {
		b.header.nonce++
	}
}

func (b *Block) isDifficult() bool {
	hash := b.hash()
	if b.header.difficulty < 0 || b.header.difficulty > len(hash) {
		return false
	}
	for i := 0; i < b.header.difficulty; i++ {
		if hash[i] != 0 {
			return false
		}
	}
	return true
}
