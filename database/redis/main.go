package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/redis"
	"github.com/lordking/toolbox/log"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	//创建一个全景的数据库访问单例
	redis := redis.New()
	err := database.CreateInstance(redis, "./redis.json")
	defer common.CheckFatal(err)

}
