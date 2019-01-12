package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	rw := bufio.NewReadWriter(r, w)

	rw.WriteString("hello world, this is go basic IO method!!,  input something: ")
	rw.Flush()

	text, err := rw.ReadString('\n')
	if err == nil {
		rw.WriteString("input: " + text)
	} else {
		fmt.Println("error: ", err)
	}

}
