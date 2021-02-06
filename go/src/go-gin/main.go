package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/routers"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("systemt:", err)
			return
		}
	}()

	r := gin.Default()

	routers.LoadApiRouter(r)
	routers.LoadBlogRouter(r)

	// 计划引流A、职业指导培训B、知识付费C、实习推荐
	// 根据设定的计划推荐C
	if err := r.Run(); err != nil {
		fmt.Println("statuup service failed, err:%v\n", err)
	}
}
