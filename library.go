package main

import (
	"fmt"
	"os"
	"strings"
)

type Book struct {
	name   string
	author string
}

func loadBooks() {
	libraries = map[string][]Book{}
	data, err := os.ReadFile("books.txt")
	must(err)
	bookStrings := strings.Split(string(data), "\n")
	for i := range len(bookStrings) / 2 {
		libraries["Debrecen"] = append(libraries["Debrecen"], Book{name: bookStrings[i*2], author: bookStrings[i*2+1]})
	}
}

func showBooks() {
	for library, books := range libraries {
		fmt.Printf("%s:\n", library)
		for i, book := range books {
			fmt.Println("\t", i, book)
		}
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

func moveBook(from string, index int, to string) {
	libraries[to] = append(libraries[to], libraries[from][index])
	libraries[from] = append(libraries[from][:index], libraries[from][index+1:]...)
	if len(libraries[from]) == 0 {
		delete(libraries, from)
	}
}
