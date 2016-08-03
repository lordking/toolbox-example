package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/http"
	"github.com/lordking/toolbox/log"
)

type (
	//Welcome 类声明
	Welcome struct{}

	//HelloForm 请求的json协议声明
	WelcomeForm struct {
		Name    string                 `json:"name" binding:"required"`
		Content map[string]interface{} `json:"content" binding:"required"`
	}
)

//Hello rest服务范例
func (w *Welcome) Hello(c *gin.Context) {

	var form WelcomeForm
	if c.BindJSON(&form) == nil {
		log.Debug("Request: %s", common.PrettyObject(form))

		http.JSONResponse(c, 200, gin.H{"hello": form.Name, "extra": form.Content})
	} else {
		http.JSONResponse(c, 400, "failure to parse json string")
	}
}
