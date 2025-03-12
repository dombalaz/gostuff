package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hi 👧 and 👦"
	fmt.Printf("str: '%v', len: %v, rune count: %v\n", message, len(message), utf8.RuneCountInString(message))
	// Get rune at specific place
	r, i := utf8.DecodeRuneInString(message[3:])
	fmt.Printf("4th rune: %c, size: %v\n", r, i)
}
