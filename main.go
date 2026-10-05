package main

import (
	"fmt"
	"log"
)

func main() {
	b := Block{}
	b.header.difficulty = 2
	b.mine()
	fmt.Println(b)
	libraries = make(map[string][]Book)
	loadBooks()
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
