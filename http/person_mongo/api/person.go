package api

import (
	ws "goutils"
	"goutils-example/webseed/person_mongo/models"
	"github.com/gin-gonic/gin"
)

type (
	//Person 类声明
	Person struct{}

	//PersonCreateForm 请求的创建person的json表单
	PersonCreateForm struct {
		Name  string `json:"name" binding:"required"`
		Phone string `json:"phone" binding:"required"`
	}

	//PersonUpdateForm 请求的更新person的json表单
	PersonUpdateForm struct {
		Phone string `json:"phone" binding:"required"`
	}
)

//Create 创建用户
func (ctrl *Person) Create(c *gin.Context) {

	var json PersonCreateForm

	err := c.BindJSON(&json)
	if err != nil {
		ws.JSONResponse(c, 403, err)
		return
	}

	p := models.Person{Name: json.Name, Phone: json.Phone}
	err = p.Create()

	if err != nil {
		err := err.(*ws.Error)
		ws.JSONResponse(c, err.Code, err.Message)
		return
	}

	ws.JSONResponse(c, 200, p)
}

//Find 查询用户
func (ctrl *Person) Find(c *gin.Context) {

	name := c.Param("name")
	p := models.Person{}
	result, err := p.Find(name)

	if err != nil {
		err := err.(*ws.Error)
		ws.JSONResponse(c, err.Code, err.Message)
		return
	}

	ws.JSONResponse(c, 200, result)
}

//Update 更新用户
func (ctrl *Person) Update(c *gin.Context) {

	name := c.Param("name")

	var json PersonUpdateForm

	err := c.BindJSON(&json)
	if err != nil {
		ws.JSONResponse(c, 403, err)
		return
	}

	p := models.Person{Phone: json.Phone}
	result, err := p.Update(name)

	if err != nil {
		err := err.(*ws.Error)
		ws.JSONResponse(c, err.Code, err.Message)
		return
	}

	ws.JSONResponse(c, 200, result)
}

//Delete 删除用户
func (ctrl *Person) Delete(c *gin.Context) {

	name := c.Param("name")
	p := models.Person{}
	result, err := p.Delete(name)

	if err != nil {
		err := err.(*ws.Error)
		ws.JSONResponse(c, err.Code, err.Message)
		return
	}

	ws.JSONResponse(c, 200, result)
}
