package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	wordArray := strings.Fields(s)
	for _, word := range wordArray {
		value, exist := m[word]
		if exist {
			m[word] = value + 1
		} else {
			m[word] = 1
		}
	}
	return m
}

func main() {
	// m := WordCount("hello world hahah hello")
	// fmt.Println(m)

	a, b := 1, 2
	fmt.Println("A: ", a, "B: ", b)
}
