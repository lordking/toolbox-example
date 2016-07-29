package main

import (
	"flag"
	ws "goutils"
	"goutils-example/webseed/person_mysql/api"

	"github.com/gin-gonic/gin"
)

var httpConfigPath = flag.String("http", "./config/http.json", "http config path")
var dbConfigPath = flag.String("db", "./config/db.json", "database config path")

func init() {
	ws.SetLogger(ws.DebugLevel)
	ws.ConfigRuntime()
	gin.SetMode(gin.ReleaseMode)

	ws.InitDataSource(*dbConfigPath)
}

func main() {
	flag.Parse()

	h := ws.CreateHTTPServer(*httpConfigPath)

	person := &api.Person{}
	personGroup := h.Group("/person")
	{
		personGroup.POST("/new", person.Create)
		personGroup.GET("/:name", person.Find)
		personGroup.PUT("/update/:name", person.Update)
		personGroup.DELETE("/delete/:name", person.Delete)
	}

	h.RunServ()
}
