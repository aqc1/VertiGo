package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// printable characters people actually use...
// No \n or \t nonsense
var (
	printable string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	builder   strings.Builder
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

	// Append Characters
	appended := appendCharacter(words)

	// Prepend Characters
	prepended := prependCharacter(words)

	// Toggle Case
	toggled := toggleCase(words)

	// Add newly created words to original
	for _, word := range appended {
		words = append(words, word)
	}
	for _, word := range prepended {
		words = append(words, word)
	}
	for _, word := range toggled {
		words = append(words, word)
	}

	// Print out to STDOUT
	for _, word := range words {
		fmt.Println(word)
	}
}

// Append a characters to each word
// :param words: Words to have characters appended to
// :return: slice containing newly created words
func appendCharacter(words []string) []string {
	// Keep track of new entries
	newWords := []string{}

	// Iterate over words
	for _, word := range words {
		// Each printable character is possible
		for _, character := range printable {

			builder.WriteString(word)
			builder.WriteRune(character)
			newWords = append(newWords, builder.String())
			builder.Reset()
		}
	}

	return newWords
}

// Prepend a characters to each word
// :param words: Words to have characters prepended to
// :return: slice containing newly created words
func prependCharacter(words []string) []string {
	// Keep track of new entries
	newWords := []string{}

	// Iterate over words
	for _, word := range words {
		// Each printable character is possible
		for _, character := range printable {
			builder.WriteRune(character)
			builder.WriteString(word)
			newWords = append(newWords, builder.String())
			builder.Reset()
		}
	}

	return newWords
}

// Toggle case of a word
// :param words: Word to toggle the case of
// :return: slice containing toggled words
func toggleCase(words []string) []string {
	// Keep track of new entries
	newWords := []string{}

	// Iterate over words
	for _, word := range words {
		// Each printable character is possible
		for _, character := range word {
			if unicode.IsUpper(character) {
				character = rune(unicode.ToLower(rune(character)))
			} else if unicode.IsLower(character) {
				character = rune(unicode.ToUpper(rune(character)))
			}
			builder.WriteString(string(character))
		}
		newWords = append(newWords, builder.String())
		builder.Reset()
	}

	return newWords
}
