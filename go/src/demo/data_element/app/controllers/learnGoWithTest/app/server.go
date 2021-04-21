package app

import (
	"fmt"
	"net/http"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func PlayerServer(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "20")
}