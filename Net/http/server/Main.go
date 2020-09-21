package main

import (
	"net"
	"net/http"
	"time"
)

// 第一种处理器,实现http.Handler {} 接口
type fooHandler struct{}

func (h fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bar := [5]byte{'b', 'a', 'r'}
	w.Write(bar[:])
}

func main() {
	//basicHttpServer()
	defaultServer()
}

func defaultServer() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello, go web"))
	})

	http.ListenAndServe(":8089", nil)
}

func basicHttpServer() {
	s := &http.Server{
		Addr:           ":9000",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 2^20 GB
	}

	// 向DefaultServeMux 中添加路由处理规则
	http.Handle("/foo", *new(fooHandler))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte{'i', 'n', 'd', 'e', 'x'})
	})

	listener, _ := net.Listen("tcp4", "127.0.0.1:9000")
	s.Serve(listener)
}
