package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {
	var rawStr = "你好世界,hello world"

	for i := 0; i < len(rawStr); i++ {
		// Rune is int32, representing a unicode character
		r, rsize := utf8.DecodeRuneInString(rawStr)
		// var c rune = '\xDF'

		fmt.Printf("Rune: %c, byte: %x\n", r, byte(r))
		rawStr = rawStr[rsize:] // Skip one character

	}

	c := '中'
	fmt.Println(reflect.TypeOf(c).Name())

}
