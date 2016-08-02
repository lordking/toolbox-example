package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mongo"
	"github.com/lordking/toolbox/log"
)

const (
	collectionName = "person"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	//创建一个全景的数据库访问单例
	mongo := mongo.New()
	err := database.CreateInstance(mongo, "./db.json")
	defer common.CheckFatal(err)

	p := &Person{
		Name:  "leking",
		Phone: "18900000000",
	}

	//插入数据
	p.insert()

	//查询数据
	p.findAll(p.Name)

	//更新数据
	p.updateAll(p.Name, "13900000")

	//删除数据
	p.removeAll(p.Name)
}
