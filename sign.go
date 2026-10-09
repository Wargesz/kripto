package main

import "crypto/ed25519"

type Transaction struct {
	Msg           string `json:"msg"`
	UserPublicKey []byte `json:"userpublickey"`
	Signature     string `json:"signature"`
}

func generateKeys() (ed25519.PublicKey, ed25519.PrivateKey) {
	public, private, err := ed25519.GenerateKey(nil)
	must(err)
	return public, private
}

func (t *Transaction) sign(privateKey ed25519.PrivateKey) {
	t.Signature = string(ed25519.Sign(privateKey, []byte(t.Msg)))
}

func (t *Transaction) verify() bool {
	return ed25519.Verify(t.UserPublicKey, []byte(t.Msg), []byte(t.Signature))
}
