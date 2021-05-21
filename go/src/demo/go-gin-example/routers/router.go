package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tmaio/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is new test",
		})
	})

	return r
}
