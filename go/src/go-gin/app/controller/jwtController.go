package controller

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// MyClaims自定义声明结构体并内嵌jwt.StandardClaims
// 额外记录一个username字段，所以要自定义结构体
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserInfo struct {
	Username string
	Password string
}

// 定义JWT的过期时间
const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("夏天过去")

func DemoJwt(c *gin.Context)  {
	// JWT 全称JSON Web Token 是一种跨域认证解决方案，属于一个开放的标准，规定了一种Token实现方式。多用于前后端分离项目和OAuth2.0业务场景
	// JWT就是一种基于Token的轻量级认证模式，服务端认证通过后，会生成一个JSON对象，经过签名后得到一个Token再发回给用户，用户后续请求只需要带上这个
	// Token，服务端解密之后就能获取该用户的相关信息了。

	// JWT服务器不保存session数据，所有数据保存在客户端。
	/*
		JWT的原理是，服务器认证后，生成一个JSON对象，发回给用户。
		以后，用户与服务器通信的时候，都要发回这个JSON对象。服务器完全只靠这个对象认定用户身份
		为了防止用户篡改数据，服务器在生成这个对象的时候，会加上签名。
		JWT的三个部分：
			header 头部
			Payload 负载
			Signature 签名
	*/

	var user UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":2001,
			"msg": "无效的参数",
		})
		return
	}

	if user.Username == "jack" && user.Password == "123456" {
		// 生成token
		tokenString, _ := genToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg": "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg": "鉴权失败",
	})
	return
}

// 生成JWT
func genToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		"username", // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer: "my-project", // 签发人
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 校验token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")

}
