package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mysql"
	"github.com/lordking/toolbox/log"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	//创建一个全景的数据库访问单例
	mysql := mysql.New()
	err := database.CreateInstance(mysql, "./db.json")
	defer common.CheckFatal(err)

	p := &Person{
		Name:  "leking",
		Phone: "18900000000",
	}

	//插入数据
	p.Insert()

	//查询数据
	p.FindAll(p.Name)

	//更新数据
	p.UpdateAll(p.Name, "13900000")

	//删除数据
	p.RemoveAll(p.Name)

}
