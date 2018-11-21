package main

import (
	"fmt"
	"strings"
)

type newt interface{}

func main() {
	sbuilder := strings.Builder{}
	sbuilder.Write([]byte{68, 78, 98, 112})
	fmt.Println(sbuilder.String())
}
