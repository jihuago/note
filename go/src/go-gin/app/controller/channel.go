package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 同时开几个协程做同样的事情，哪个先搞定，用那个的结果
func GetResult(c *gin.Context)  {
	txtRes := make(chan string, 5)

	go func() {txtRes <- getTxt("api1.baidu.com")}()
	go func() {txtRes <- getTxt("api2.baidu.com")}()
	go func() {txtRes <- getTxt("api3.baidu.com")}()
	go func() {txtRes <- getTxt("api4.baidu.com")}()
	go func() {txtRes <- getTxt("api5.baidu.com")}()

	c.JSON(http.StatusOK, gin.H{
		"result": <- txtRes,
	})
}

func getTxt(host string) string {
	if host == "api4.baidu.com" {
		time.Sleep(3e9)
	} else {
		time.Sleep(2e9)
	}

	return host + ": 模拟结果"
}