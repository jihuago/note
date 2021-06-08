package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/tmaio/go-gin-example/models"
	"github.com/tmaio/go-gin-example/pkg/logging"
	"github.com/tmaio/go-gin-example/pkg/setting"
	"github.com/tmaio/go-gin-example/routers"
	"log"
	"syscall"
)


func main() {

/*	router := gin.Default()
	router.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "test",
		})
	})
*/
	setting.Setup()
	models.Setup()
	logging.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err:%v", err)
	}

	/*
	router := routers.InitRouter()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()*/
}
