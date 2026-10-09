package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Header struct {
	PrevHash   string `json:"prevHash"`
	MerkleRoot string `json:"merkleRoot"`
	Nonce      int    `json:"nonce"`
	Difficulty int    `json:"difficulty"`
	Version    int    `json:"version"`
}

type Block struct {
	Header     Header
	MerkleTree tree   `json:"merkleTree"`
}

func (b *Block) hash() [32]byte {
	data, err := json.Marshal(fmt.Sprintf("%v", b.Header))
	if err != nil {
		must(err)
	}
	return sha256.Sum256(data)
}

func (b *Block) mine() {
	for !b.isDifficult() {
		b.Header.Nonce++
	}
}

func (b *Block) isDifficult() bool {
	hash := b.hash()
	if b.Header.Difficulty < 0 || b.Header.Difficulty > len(hash) {
		return false
	}
	for i := 0; i < b.Header.Difficulty; i++ {
		if hash[i] != 0 {
			return false
		}
	}
	return true
}
