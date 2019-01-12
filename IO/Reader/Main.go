package main

import (
	"fmt"
	"io"
	"strings"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := len(b); i > 0; i-- {
		b[i-1] = 'A'
	}

	return 1, nil
}

func main() {
	r := strings.NewReader("Hello, Reader!") // 创建一个string Reader

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF { // stream尾返回io.EOF
			break
		}
	}

	myreader := MyReader{}
	v, _ := myreader.Read(b)
	fmt.Printf("Read from MyReader: %v, bytes: %d", b, v)
}
