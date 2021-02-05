package common_package

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

//Go内置的 net/http包提供了HTTP客户端和服务端的实现

func DemoGet()  {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Panicln(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("get resp failed, err:%v", err)
	}

	//fmt.Println(string(body))

	file, err := os.OpenFile("./public/baidu.html", os.O_CREATE|os.O_WRONLY, 0666)

	if err  != nil{
		fmt.Println(err)
		return
	}

	defer file.Close()

	file.WriteString(string(body))

}

func DemoTestWithParam()  {
	apiurl := "http://baidu.com/s"

	data := url.Values{}
	data.Set("wd", "test")
	
	u, err := url.ParseRequestURI(apiurl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err%v\n", err)
		return
	}

	u.RawQuery = data.Encode()
	//fmt.Println(u.String()) // u.String() 获取请求url

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v \n", err)
		return
	}

	file, err := os.OpenFile("./public/" + Date(time.Now(), "20060102030405") + ".html", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open file failed:%v", err)
		return
	}

	defer file.Close()

	file.WriteString(string(b))

}
