package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User string `from:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//演示.ShouldBind()功能，基于请求自动提取JSON、form表单和QueryString类型的数据，并把值绑定到指定的结构体对象
func HandleLogin(c *gin.Context)  {
	var login Login

	if err := c.ShouldBind(&login); err == nil {
		fmt.Printf("login info:%#v\n", login)
		c.JSON(http.StatusOK, gin.H{
			"user": login.User,
			"password": login.Password,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
