package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin/app/controller"
	"net/http"
)

func LoadBlogRouter(r *gin.Engine)  {
	r.GET("/channel", controller.GetResult)

	// 参数绑定
	r.POST("/loginJSON", controller.HandleLogin)

	// JWT
	r.GET("/jwt", controller.DemoJwt)

	// 获取querystring参数
	r.GET("/querystring", controller.ResponseUserInfo)

	// 首页，展示一个表单
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "test",
		})
	})
}