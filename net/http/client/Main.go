package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	httpRequest101()

}

func httpRequest101() {
	// 基本http请求
	resp, err := http.Get("http://www.linkinghack.com")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	defer resp.Body.Close()

	// 构建bytes数组来接收二进制数据
	body := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := resp.Body.Read(buf[0:])
		fmt.Println("read :", n)
		if err != nil {
			fmt.Println("error: ", err)
			break
		}
		body.Write(buf[0:n])
	}

	fmt.Println("result: ")
	fmt.Println(body.String())
}
