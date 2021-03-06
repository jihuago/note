package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin/app/controller"
)

func LoadBlogRouter(r *gin.Engine)  {
	r.GET("/channel", controller.GetResult)

	// 参数绑定
	r.POST("/loginJSON", controller.HandleLogin)

	// JWT
	r.GET("/jwt", controller.DemoJwt)
}