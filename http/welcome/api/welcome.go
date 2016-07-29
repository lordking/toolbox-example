package api

import (
	ws "goutils"

	"github.com/gin-gonic/gin"
)

type (
	//Welcome 类声明
	Welcome struct{}

	//HelloForm 请求的json协议声明
	HelloForm struct {
		Name    string                 `json:"name" binding:"required"`
		Content map[string]interface{} `json:"content" binding:"required"`
	}
)

//Hello rest服务范例
func (w *Welcome) Hello(c *gin.Context) {

	var json HelloForm
	if c.BindJSON(&json) == nil {
		ws.LogDebug("Request: %s", json)

		ws.JSONResponse(c, 200, gin.H{"hello": json.Name, "extra": json.Content})
	} else {
		ws.JSONResponse(c, 400, "failure to parse json string")
	}
}
