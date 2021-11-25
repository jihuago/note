package main

import (
	"fmt"
	"log"
	"net/http"
)

// 使用自定义的

func main()  {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}

type Engine struct {}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 这样可以自定义路由映射的规则，统一添加一些处理逻辑，拥有统一的控制入口
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
