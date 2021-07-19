package main

import "net/http"

type MyHandler struct {
	
}

func (myHan *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello World"))
}

func main()  {
	mH := &MyHandler{}
	server := http.Server{
		Addr: ":8080",
		Handler: mH,
	}

	server.ListenAndServe()
}
