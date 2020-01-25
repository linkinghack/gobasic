package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var rawStr = "你好世界,hello world"

	for i := 0; i < len(rawStr); i++ {
		// Rune is int32, representing a unicode character
		r, rsize := utf8.DecodeRuneInString(rawStr)
		// var c rune = '\xDF'

		fmt.Printf("Rune: %x, byte: %x\n", r, byte(r))
		rawStr = rawStr[rsize:]
	}

}
