package main

import (
	"os"
	"strings"
)

type Book struct {
	name   string
	author string
}

var libraries map[string][]Book

func loadBooks() {
	data, err := os.ReadFile("books.txt")
	must(err)
	bookStrings := strings.Split(string(data), "\n")
	for i := range len(bookStrings) / 2 {
		libraries["Debrecen"] = append(libraries["Debrecen"], Book{name: bookStrings[i*2], author: bookStrings[i*2+1]})
	}
}

func findBook(name string) string {
	for library, books := range libraries {
		for _, book := range books {
			if book.name == name {
				return library
			}
		}
	}
	return ""
}
