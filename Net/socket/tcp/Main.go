package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		// os.Args[0] is the filename of the program
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}

	target := os.Args[1]

	conn, err := net.Dial("tcp", target)
	checkErr(err)

	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n")) //发送一个标准http请求行
	fmt.Println("tcp 消息发送完毕，长度： ", n)

	result, err := readFully(conn)
	fmt.Println("read over")
	checkErr(err)
	fmt.Println("收到tcp消息： ")
	fmt.Println(string(result))

	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		fmt.Println(result.Len())
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil // 无错误返回
}
