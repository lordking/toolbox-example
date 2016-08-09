package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mongo"
	"github.com/lordking/toolbox/http"
	"github.com/lordking/toolbox/log"

	"github.com/lordking/toolbox-example/http/blog/api"
	"github.com/lordking/toolbox-example/http/blog/models"
)

var token *models.Token

func init() {
	log.SetLevel(log.DebugLevel)
	common.ConfigRuntime()
	gin.SetMode(gin.ReleaseMode)
}

func authorization(username, password string) error {

	log.Debug("auth: %s : %s", username, password)

	if username != "" {

		result, _ := token.Find(username)
		if result != nil && result.ExpireTime > time.Now().Unix() {
			log.Debug("%s auth ok", username)
			return nil
		}
	}

	str := fmt.Sprintf("%s auth failure", username)
	return common.NewError(401, str)
}

func main() {

	//读取环境变量
	httpConfigPath := flag.String("http", "./config/http.json", "http config path")
	dbConfigPath := flag.String("db", "./config/mongo.json", "database config path")
	flag.Parse()

	//创建一个数据库访问单例
	mongo := mongo.New()
	err := database.Configure(mongo, *dbConfigPath)
	defer common.CheckFatal(err)

	h := http.CreateServer(*httpConfigPath)
	h.Router.Use(static.Serve("/", static.LocalFile("./assets", false)))

	user, err := api.NewUser(mongo)
	defer common.CheckFatal(err)
	userGroup := h.Group("/user")
	{
		userGroup.POST("/login", user.Login)
	}

	token, err = models.NewToken(mongo)
	defer common.CheckFatal(err)

	blog, err := api.NewBlog(mongo)
	defer common.CheckFatal(err)
	blogGroup := h.Group("/blog", http.BasicAuth(authorization))
	{
		blogGroup.POST("/new", blog.Create)
		blogGroup.GET("/:start/:limit", blog.Find)
		blogGroup.PUT("/update/:id", blog.Update)
		blogGroup.DELETE("/delete/:id", blog.Delete)
	}

	h.RunServ()
}
