package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("tests/wordlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))
}
