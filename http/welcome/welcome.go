package main

import (
	"flag"
	ws "goutils"

	"goutils-example/webseed/welcome/api"

	"github.com/gin-gonic/gin"
)

func init() {
	ws.SetLogger(ws.DebugLevel)
	ws.ConfigRuntime()
	gin.SetMode(gin.ReleaseMode)
}

var httpConfigPath = flag.String("http", "./config/http.json", "http config path")

func main() {
	flag.Parse()

	h := ws.CreateHTTPServer(*httpConfigPath)

	// 一个rest服务的简单范例
	welcome := &api.Welcome{}
	welcomeGroup := h.Group("/welcome")
	{
		welcomeGroup.POST("/hello", welcome.Hello)
	}

	h.RunServ()
}
