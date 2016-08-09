package main

import (
	"database/sql"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database/mysql"
	"github.com/lordking/toolbox/log"
)

const (
	collectionName = "person"
)

type (
	Person struct {
		db *mysql.MySQL
	}

	//Person 用户数据对象
	PersonVO struct {
		Id    int
		Name  string `json:"name" bson:"name"`
		Phone string `json:"phone" bson:"phone"`
	}
)

func (p *Person) Insert(obj *PersonVO) {

	conn := (p.db.GetConnection()).(*sql.DB)

	stmt, err := conn.Prepare("INSERT INTO person(name, phone) VALUES(?, ?)")
	defer stmt.Close()
	defer common.CheckFatal(err)

	result, err := stmt.Exec(obj.Name, obj.Phone)
	lastId, err := result.LastInsertId()
	defer common.CheckFatal(err)

	log.Debug("Insert result: the `id` of a new row is `%d`", lastId)

}

func (p *Person) FindAll(name string) {

	conn := (p.db.GetConnection()).(*sql.DB)

	var result []PersonVO
	stmt, err := conn.Query("SELECT id, name, phone FROM person WHERE name = ?", name)
	defer stmt.Close()
	defer common.CheckFatal(err)

	for stmt.Next() {
		var obj PersonVO
		err := stmt.Scan(&(obj.Id), &(obj.Name), &(obj.Phone))
		defer common.CheckError(err)

		result = append(result, obj)
	}

	log.Debug("Find result: %s", common.PrettyObject(result))

}

func (p *Person) UpdateAll(name string, obj *PersonVO) {

	conn := (p.db.GetConnection()).(*sql.DB)

	stmt, err := conn.Prepare("UPDATE person SET phone=? where name=?")
	defer stmt.Close()
	defer common.CheckFatal(err)

	result, err := stmt.Exec(obj.Phone, name)
	rowsCount, err := result.RowsAffected()
	defer common.CheckFatal(err)

	log.Debug("Update result: the sum of effected rows is `%d`", rowsCount)

}

func (p *Person) RemoveAll(name string) {

	conn := (p.db.GetConnection()).(*sql.DB)

	stmt, err := conn.Prepare("DELETE FROM person WHERE name=?")
	defer stmt.Close()
	defer common.CheckFatal(err)

	result, err := stmt.Exec(name)
	rowsCount, err := result.RowsAffected()
	defer common.CheckFatal(err)

	log.Debug("Delete result: the sum of effected rows is `%d`", rowsCount)

}

func NewPerson(db *mysql.MySQL) (*Person, error) {

	err := db.Connect()

	return &Person{
		db: db,
	}, err
}
