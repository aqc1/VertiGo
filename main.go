package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Temporary test file
	// This will eventually be a CLI arg
	handle, err := os.Open("tests/wordlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Grab words from wordlist and add to a slice
	words := []string{}
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// Process slice here
	appended := addAfter.appendCharacter(words)
	mangled := append(words, appended)

	for _, word := range mangled {
		fmt.Println(word)
	}
}
