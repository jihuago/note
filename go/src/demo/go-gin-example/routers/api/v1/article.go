package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/tmaio/go-gin-example/models"
	"github.com/tmaio/go-gin-example/pkg/e"
	"github.com/tmaio/go-gin-example/pkg/logging"
	"github.com/unknwon/com"
	"net/http"
)

//获取单个文章
func GetArticle(c *gin.Context)  {
}

func GetArticles(c *gin.Context)  {
	total := models.GetArticleTotal(map[string]interface{}{"title": 3333})

	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

func AddArticle(c *gin.Context)  {

}

func EditArticle(c *gin.Context)  {

}

// 删除文章
func DeleteArticle(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	
	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleById(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			//log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}
