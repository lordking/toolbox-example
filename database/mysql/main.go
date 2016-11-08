package main

import (
	"fmt"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mysql"
	"github.com/lordking/toolbox/log"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
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
	mysql := mysql.New()
	err := database.ConfigureCfgKey(mysql, "database")
	defer common.CheckFatal(err)

	form := &PersonVO{
		Name:  "leking",
		Phone: "18900000000",
	}

	p, err := NewPerson(mysql)
	defer common.CheckFatal(err)

	//插入数据
	p.Insert(form)

	//查询数据
	p.FindAll(form.Name)

	//更新数据
	form.Phone = "13900001111"
	p.UpdateAll(form.Name, form)

	//删除数据
	p.RemoveAll(form.Name)

}
