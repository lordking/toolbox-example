package main

import (
	"database/sql"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mysql"
	"github.com/lordking/toolbox/log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	collectionName = "person"
)

//Person 用户数据对象
type Person struct {
	Id    int
	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`
}

func init() {
	log.SetLevel(log.DebugLevel)
}

func insert(p *Person) {

	//获取单例
	db := (database.Instance).(*mysql.MySQL)

	err := db.Connect()
	defer common.CheckFatal(err)

	connection := (db.GetConnection()).(*sql.DB)

	stmt, err := connection.Prepare("INSERT INTO person(name, phone) VALUES(?, ?)")
	defer common.CheckError(err)
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, p.Phone)
	defer common.CheckError(err)

	log.Debug("Insert result: %s", common.PrettyObject(result))

	db.Close()
}

func main() {
	//创建一个全景的数据库访问单例
	mysql := mysql.New()
	err := database.CreateInstance(mysql, "./db.json")
	common.CheckFatal(err)

	// p := &Person{
	// 	Name:  "leking",
	// 	Phone: "18900000000",
	// }
	//
	// //插入数据
	// insert(p)
}
