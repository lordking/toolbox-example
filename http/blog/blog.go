package main

import (
	"flag"
	ws "goutils"
	"goutils-example/webseed/blog/api"
	m "goutils-example/webseed/blog/models"

	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var httpConfigPath = flag.String("http", "./config/http.json", "http config path")
var dbConfigPath = flag.String("db", "./config/db.json", "database config path")

func init() {
	ws.SetLogger(ws.DebugLevel)
	ws.ConfigRuntime()
	gin.SetMode(gin.DebugMode)

	ws.InitDataSource(*dbConfigPath)
}

func authorization(username, password string) error {

	ws.LogDebug("auth: %s : %s", username, password)

	if username != "" {

		token := &m.Token{}
		result, _ := token.Find(username)

		if result != nil && result.ExpireTime > time.Now().Unix() {
			ws.LogDebug("auth ok")
			return nil
		}
	}

	return ws.NewError(401, "wrong token")
}

func main() {
	flag.Parse()

	h := ws.CreateHTTPServer(*httpConfigPath)
	h.Router.Use(static.Serve("/", static.LocalFile("./assets", false)))

	user := &api.User{}
	userGroup := h.Group("/user")
	{
		userGroup.POST("/login", user.Login)
	}

	blog := &api.Blog{}
	blogGroup := h.Group("/blog", ws.BasicAuth(authorization))
	{
		blogGroup.POST("/new", blog.Create)
		blogGroup.GET("/:start/:limit", blog.Find)
		blogGroup.PUT("/update/:id", blog.Update)
		blogGroup.DELETE("/delete/:id", blog.Delete)
	}

	h.RunServ()
}
