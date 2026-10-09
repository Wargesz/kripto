package main

import "testing"

func TestSignAndVerifySuccess(t *testing.T) {
	pub, priv := generateKeys()
	transaction := Transaction{Msg: "transaction message", UserPublicKey: pub}
	transaction.sign(priv)
	if !transaction.verify() {
		t.Errorf("transaction verification failed")
	}
}

func TestSignAndVerifyFail(t *testing.T) {
	pub, priv := generateKeys()
	transaction := Transaction{Msg: "transaction message", UserPublicKey: pub}
	transaction.sign(priv)
	transaction.Msg = "tampered message"
	if transaction.verify() {
		t.Errorf("transaction verification failed")
	}
}
