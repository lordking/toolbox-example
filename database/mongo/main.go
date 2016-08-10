package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mongo"
	"github.com/lordking/toolbox/log"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {

	//创建一个数据库访问单例
	mongo := mongo.New()
	err := database.Configure(mongo, "./mongo.json")
	defer common.CheckFatal(err)

	form := &PersonVO{
		Name:  "leking",
		Phone: "18900000000",
	}

	p, err := NewPerson(mongo)
	defer common.CheckFatal(err)

	//插入数据
	p.insert(form)
	defer common.CheckFatal(err)

	//查询数据
	p.findAll(form.Name)
	defer common.CheckFatal(err)

	//更新数据
	form.Phone = "13900001111"
	p.updateAll(form.Name, form)
	defer common.CheckFatal(err)

	//删除数据
	p.removeAll(form.Name)
	defer common.CheckFatal(err)
}
