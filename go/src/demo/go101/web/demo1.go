package web

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
}

/*
	Go的http包详解
		Go的http有两个核心功能：Conn、ServeMux

 */
func RunWeb()  {
	http.HandleFunc("/", sayHelloName)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


