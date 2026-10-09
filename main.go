package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var tranChan chan Transaction
var blockchain []Block
var users []User
var libraries map[string][]Book

func main() {
	initializeBlockChain()
	tranChan = make(chan Transaction, 4)
	go minerWorker()
	r := bufio.NewReader(os.Stdin)
	if len(users) == 0 {
		fmt.Print("username: ")
		username, err := r.ReadString('\n')
		must(err)
		users = append(users, newUser(username))
	}
	loadBooks()
	for {
		print("\033[2J\033[H")
		showBooks()
		print(">")
		msg, err := r.ReadString('\n')
		must(err)
		params := strings.Split(msg[:len(msg)-1], " ")
		bookIndex, err := strconv.Atoi(params[1])
		must(err)
		transactionMessage := fmt.Sprintf("%s %s->%s", libraries[params[0]][bookIndex], params[0], params[2])
		moveBook(params[0], bookIndex, params[2])
		t := Transaction{Msg: transactionMessage, UserPublicKey: users[0].PublicKey}
		t.sign(users[0].PrivateKey)
		tranChan <- t
	}
}

func initializeUser() {
	r := bufio.NewReader(os.Stdin)
	if len(users) == 0 {
		fmt.Println("username:")
		username, err := r.ReadString('\n')
		must(err)
		users = append(users, newUser(username))
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
