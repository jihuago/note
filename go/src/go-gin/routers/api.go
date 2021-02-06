package routers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-gin/app/controller"
	"net/http"
)

func LoadApiRouter(r *gin.Engine) {
	r.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"name": "jack",
		})
	})

	r.GET("/test", controller.Index)
	
	//获取json参数
	r.POST("/json", func(context *gin.Context) {
		b, _ := context.GetRawData() // 从context.Request.Body读取请求数据
		var m map[string]interface{}

		// 反序列化
		_ = json.Unmarshal(b, &m)
		context.JSON(http.StatusOK, m)
	})

	// 获取path参数：请求的参数通过URL路径传递，例如/user/search/小jack
	r.GET("/user/search/:username", func(c *gin.Context) {
		username := c.Param("username")

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"username2": username,
			"status": username,
			"a": "a",
		})
	})
}
