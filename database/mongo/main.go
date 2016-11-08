package main

import (
	"fmt"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mongo"
	"github.com/lordking/toolbox/log"
	"github.com/spf13/viper"
)

func init() {
	initConfig()
	log.SetLogDefaults("log")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Read config file error: ", err)
	}
}

func main() {

	//创建一个数据库访问单例
	mongo := mongo.New()
	err := database.ConfigureCfgKey(mongo, "database")
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
