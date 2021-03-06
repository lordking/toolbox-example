package main

import (
	"flag"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mongo"
	"github.com/lordking/toolbox/log"
)

var (
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "config", "", "config file.")
}

func main() {

	common.InitConfig("mongo_exmple", cfgFile)
	log.SetLogDefaults("log")

	//创建一个数据库访问单例
	mongo := mongo.New()
	err := database.Configure("database", mongo)
	common.CheckFatal(err)

	form := &PersonVO{
		Name:  "leking",
		Phone: "18900000000",
	}

	p, err := NewPerson(mongo)
	common.CheckFatal(err)

	//插入数据
	p.insert(form)
	common.CheckFatal(err)

	//查询数据
	p.findAll(form.Name)
	common.CheckFatal(err)

	//更新数据
	form.Phone = "13900001111"
	p.updateAll(form.Name, form)
	common.CheckFatal(err)

	//删除数据
	p.removeAll(form.Name)
	common.CheckFatal(err)
}
