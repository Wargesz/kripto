package main

import "crypto/ed25519"

type User struct {
	Name       string
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
}

func newUser(username string) User {
	pub, priv, err := ed25519.GenerateKey(nil)
	must(err)
	return User{Name: username, PublicKey: pub, PrivateKey: priv}

}
