package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/http"
	"github.com/lordking/toolbox/log"

	"github.com/lordking/toolbox-example/http/welcome/api"
)

func init() {
	log.SetLevel(log.DebugLevel)
	common.ConfigRuntime()
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	httpConfigPath := flag.String("http", "./config/http.json", "http config path")
	flag.Parse()

	h := http.CreateServer(*httpConfigPath)

	// 一个rest服务的简单范例
	welcome := &api.Welcome{}
	group := h.Group("/welcome")
	{
		group.POST("/hello", welcome.Hello)
	}

	h.RunServ()
}
