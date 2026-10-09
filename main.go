package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

var tranChan chan Transaction
var blockchain []Block
var users []User

func main() {
	initializeBlockChain()
	tranChan = make(chan Transaction, 4)
	go minerWorker()
	r := bufio.NewReader(os.Stdin)
	if len(users) == 0 {
		fmt.Println("username:")
		username, err := r.ReadString('\n')
		must(err)
		users = append(users, newUser(username))
	}
	for {
		msg, err := r.ReadString('\n')
		must(err)
		t := Transaction{Msg: msg, UserPublicKey: users[0].PublicKey}
		t.sign(users[0].PrivateKey)
		tranChan <- t
	}
}

func initializeBlockChain() {
	h := Header{
		PrevHash:   "0",
		MerkleRoot: "0",
		Nonce:      0,
		Difficulty: 0,
		Version:    1,
	}
	b := Block{
		Header:     h,
		MerkleTree: tree{},
	}
	blockchain = append(blockchain, b)
}

func minerWorker() {
	var transactions [4]Transaction
	head := 0
	for {
		transactions[head] = <-tranChan
		head++
		if head == 4 {
			extendBlockchain(transactions)
			head = 0
		}
	}
}

func extendBlockchain(transactions [4]Transaction) {
	var transactionString []string
	fmt.Println(transactions)
	for _, transaction := range transactions {
		if !transaction.verify() {
			must(errors.New("failed to verify transaction"))
		}
		data, err := json.Marshal(transaction)
		must(err)
		transactionString = append(transactionString, string(data))
	}
	tree := buildMerkleTree(transactionString)
	prevHash := fmt.Sprintf("%x", blockchain[len(blockchain)-1].hash())
	h := Header{
		PrevHash:   prevHash,
		MerkleRoot: tree[0][0],
		Difficulty: 2,
		Version:    1,
	}
	b := Block{
		Header:     h,
		MerkleTree: tree,
	}
	b.mine()
	blockchain = append(blockchain, b)
	data, err := json.Marshal(blockchain)
	must(err)
	writeBlockhain(data)
}

func writeBlockhain(data []byte) {
	err := os.WriteFile("blockchain.json", data, 0644)
	must(err)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
