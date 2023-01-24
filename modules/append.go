package modules

import (
	"strings"
)

// Printable characters
const printable string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

func appendCharacter(words []string) []string {
	// Keep track of new entries
	newWords := []string{}
	var builder strings.Builder

	// Iterate over words
	for _, word := range words {
		for _, character := range printable {

			builder.WriteString(word)
			builder.WriteRune(character)
			newWords = append(newWords, builder.String())
			builder.Reset()
		}
	}

	return newWords
}
