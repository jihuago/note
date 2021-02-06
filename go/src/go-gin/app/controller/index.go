package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {

	username := c.DefaultQuery("username", "小王子")

	c.JSON(http.StatusOK, gin.H{
		"status":  username,
	})
}
