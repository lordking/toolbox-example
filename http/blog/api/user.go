package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/http"

	"github.com/lordking/toolbox-example/http/blog/models"
)

type (
	//User controller声明
	User struct {
		token *models.Token
	}

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
		http.JSONResponse(c, 403, err)
		return
	}

	if json.Username != "admin" && json.Password != "admin" {
		http.JSONResponse(c, 401, "用户名或密码错误")
		return
	}

	obj := &models.TokenVO{}
	err = u.token.Create(obj)
	if err != nil {
		err := err.(*common.Error)
		http.JSONResponse(c, err.Code, err.Message)
	}

	u.token.ClearExpireTokens() //清除已退休的token
	http.JSONResponse(c, 200, obj)
}

func NewUser() (*User, error) {

	token, err := models.NewToken()

	ctrl := &User{
		token: token,
	}

	return ctrl, err
}
