package main

import (
	"database/sql"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mysql"
	"github.com/lordking/toolbox/log"
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

func (p *Person) Insert() {

	//获取单例
	db := (database.Instance).(*mysql.MySQL)

	err := db.Connect()
	defer common.CheckFatal(err)

	connection := (db.GetConnection()).(*sql.DB)

	stmt, err := connection.Prepare("INSERT INTO person(name, phone) VALUES(?, ?)")
	defer stmt.Close()
	defer common.CheckFatal(err)

	result, err := stmt.Exec(p.Name, p.Phone)
	lastId, err := result.LastInsertId()
	defer common.CheckFatal(err)

	log.Debug("Insert result: the `id` of a new row is `%d`", lastId)

	db.Close()
}

func (p *Person) FindAll(name string) {

	//获取单例
	db := (database.Instance).(*mysql.MySQL)

	err := db.Connect()
	defer common.CheckFatal(err)

	connection := (db.GetConnection()).(*sql.DB)

	var result []Person
	stmt, err := connection.Query("SELECT id, name, phone FROM person WHERE name = ?", name)
	defer stmt.Close()
	defer common.CheckFatal(err)

	for stmt.Next() {
		var person Person
		err := stmt.Scan(&(person.Id), &(person.Name), &(person.Phone))
		defer common.CheckError(err)

		result = append(result, person)
	}

	log.Debug("Find result: %s", common.PrettyObject(result))

	db.Close()
}

func (p *Person) UpdateAll(name string, phone string) {

	//获取单例
	db := (database.Instance).(*mysql.MySQL)

	err := db.Connect()
	defer common.CheckFatal(err)

	connection := (db.GetConnection()).(*sql.DB)

	stmt, err := connection.Prepare("UPDATE person SET phone=? where name=?")
	defer stmt.Close()
	defer common.CheckFatal(err)

	result, err := stmt.Exec(phone, name)
	rowsCount, err := result.RowsAffected()
	defer common.CheckFatal(err)

	log.Debug("Update result: the sum of effected rows is `%d`", rowsCount)

	db.Close()
}

func (p *Person) RemoveAll(name string) {

	//获取单例
	db := (database.Instance).(*mysql.MySQL)

	err := db.Connect()
	defer common.CheckFatal(err)

	connection := (db.GetConnection()).(*sql.DB)

	stmt, err := connection.Prepare("DELETE FROM person WHERE name=?")
	defer stmt.Close()
	defer common.CheckFatal(err)

	result, err := stmt.Exec(name)
	rowsCount, err := result.RowsAffected()
	defer common.CheckFatal(err)

	log.Debug("Delete result: the sum of effected rows is `%d`", rowsCount)

	db.Close()
}
