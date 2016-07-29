package api

import (
	ws "goutils"

	"github.com/gin-gonic/gin"

	m "goutils-example/webseed/blog/models"
)

type (
	//User controller声明
	User struct{}

	//UserLoginForm 登录的json协议
	UserLoginForm struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)

//Login 用户登录
func (u *User) Login(c *gin.Context) {

	var json UserLoginForm

	err := c.BindJSON(&json)
	if err != nil {
		ws.JSONResponse(c, 403, err)
		return
	}

	if json.Username != "admin" && json.Password != "admin" {
		ws.JSONResponse(c, 401, "用户名或密码错误")
		return
	}

	token := m.Token{}
	t, _ := token.Create()

	t.ClearExpireTokens() //清除已退休的token

	ws.JSONResponse(c, 200, t)

}
