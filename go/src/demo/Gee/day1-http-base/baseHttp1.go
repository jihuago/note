package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	// http.ListenAndServe()启动Web服务，第一个参数是地址，:9999表示端口监听
	// 第二个参数则代表处理所有的HTTP请求的实例，nil表示使用标准库的实例处理（代码如下）
	/*
	func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
		handler := sh.srv.Handler
		if handler == nil {
			handler = DefaultServeMux
		}
		if req.RequestURI == "*" && req.Method == "OPTIONS" {
			handler = globalOptionsHandler{}
		}
		handler.ServeHTTP(rw, req)
	}
	 */
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request)  {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

