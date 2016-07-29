package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	ws "goutils"
	"goutils-example/webseed/blog/models"
)

type (

	//Blog controller声明
	Blog struct{}

	//BlogCreateForm 创建日志的json协议
	BlogCreateForm struct {
		Subject string `json:"subject" binding:"required"`
		Blog    string `json:"blog" binding:"required"`
		Author  string `json:"author" binding:"required"`
	}

	//BlogUpdateForm 更新日志的json协议
	BlogUpdateForm struct {
		Subject string `json:"subject" `
		Blog    string `json:"blog" `
	}
)

//Create 创建日志
func (ctrl *Blog) Create(c *gin.Context) {

	var json BlogCreateForm

	err := c.BindJSON(&json)
	if err != nil {
		ws.JSONResponse(c, 403, err)
		return
	}

	b := models.Blog{Subject: json.Subject, Blog: json.Blog, Author: json.Author}
	err = b.Create()

	if err != nil {
		err := err.(*ws.Error)
		ws.JSONResponse(c, err.Code, err.Message)
		return
	}

	ws.JSONResponse(c, 200, b)
}

//Find 查找日志
func (ctrl *Blog) Find(c *gin.Context) {

	start, err := strconv.Atoi(c.Param("start"))
	if err != nil {
		ws.JSONResponse(c, 403, err)
		return
	}

	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		ws.JSONResponse(c, 403, err)
		return
	}

	b := models.Blog{}
	result, err := b.Find(start, limit)

	if err != nil {
		err := err.(*ws.Error)
		ws.JSONResponse(c, err.Code, err.Message)
		return
	}

	ws.JSONResponse(c, 200, result)
}

//Update 更新日志
func (ctrl *Blog) Update(c *gin.Context) {

	id := c.Param("id")
	var json BlogUpdateForm

	err := c.BindJSON(&json)
	if err != nil {
		ws.JSONResponse(c, 403, err)
		return
	}

	b := models.Blog{Subject: json.Subject, Blog: json.Blog}
	err = b.Update(id)

	if err != nil {
		err := err.(*ws.Error)
		ws.JSONResponse(c, err.Code, err.Message)
		return
	}

	ws.JSONResponse(c, 200, "ok")
}

//Delete 删除日志
func (ctrl *Blog) Delete(c *gin.Context) {

	id := c.Param("id")
	b := models.Blog{}
	err := b.Delete(id)

	if err != nil {
		err := err.(*ws.Error)
		ws.JSONResponse(c, err.Code, err.Message)
		return
	}

	ws.JSONResponse(c, 200, "ok")
}
