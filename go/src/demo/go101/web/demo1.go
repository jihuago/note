package web

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
}

func login(w http.ResponseWriter, r *http.Request)  {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./web/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func upload(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		currentTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("./web/upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./web/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

/*
	Go的http包详解
		Go的http有两个核心功能：Conn、ServeMux

 */
func RunWeb()  {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


