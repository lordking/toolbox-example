package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mysql"
	"github.com/lordking/toolbox/http"
	"github.com/lordking/toolbox/log"

	"github.com/lordking/toolbox-example/http/person_mysql/api"
)

func init() {
	log.SetLevel(log.DebugLevel)
	common.ConfigRuntime()
	gin.SetMode(gin.ReleaseMode)
}

func main() {

	//读取环境变量
	httpConfigPath := flag.String("http", "./config/http.json", "http config path")
	dbConfigPath := flag.String("db", "./config/mysql.json", "database config path")
	flag.Parse()

	//创建一个数据库访问单例
	mysql := mysql.New()
	err := database.CreateInstance(mysql, *dbConfigPath)
	defer common.CheckFatal(err)

	//创建web服务
	h := http.CreateServer(*httpConfigPath)

	person, err := api.NewPerson()
	defer common.CheckError(err)

	group := h.Group("/person")
	{
		group.POST("/new", person.Create)
		group.GET("/:name", person.Find)
		group.PUT("/update/:name", person.Update)
		group.DELETE("/delete/:name", person.Delete)
	}

	h.RunServ()
}
