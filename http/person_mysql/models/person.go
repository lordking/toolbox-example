package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database/mysql"
)

type (
	Person struct {
		db *mysql.MySQL
	}

	PersonVO struct {
		Id    int64  `json:"id" bson:"_id"`
		Name  string `json:"name" bson:"name"`
		Phone string `json:"phone" bson:"phone"`
	}
)

func (p *Person) Create(obj *PersonVO) error {

	conn := (p.db.GetConnection()).(*sql.DB)

	stmt, err := conn.Prepare("INSERT INTO person(name, phone) VALUES(?, ?)")
	defer stmt.Close()
	if err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	result, err := stmt.Exec(obj.Name, obj.Phone)
	if err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	obj.Id = lastId

	return nil
}

func (p *Person) Find(name string) ([]PersonVO, error) {

	conn := (p.db.GetConnection()).(*sql.DB)

	var result []PersonVO
	stmt, err := conn.Query("SELECT id, name, phone FROM person WHERE name = ?", name)
	defer stmt.Close()
	if err != nil {
		return nil, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	for stmt.Next() {
		var obj PersonVO
		err := stmt.Scan(&(obj.Id), &(obj.Name), &(obj.Phone))
		defer common.CheckError(err)

		result = append(result, obj)
	}

	return result, nil
}

func (p *Person) Update(name string, obj *PersonVO) (int64, error) {

	conn := (p.db.GetConnection()).(*sql.DB)

	stmt, err := conn.Prepare("UPDATE person SET phone=? where name=?")
	defer stmt.Close()
	if err != nil {
		return -1, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	result, err := stmt.Exec(obj.Phone, name)
	rowsCount, err := result.RowsAffected()
	if err != nil {
		return -1, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return rowsCount, nil
}

func (p *Person) Delete(name string) (int64, error) {

	conn := (p.db.GetConnection()).(*sql.DB)

	stmt, err := conn.Prepare("DELETE FROM person WHERE name=?")
	defer stmt.Close()
	if err != nil {
		return -1, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	result, err := stmt.Exec(name)
	rowsCount, err := result.RowsAffected()
	if err != nil {
		return -1, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return rowsCount, nil
}

func NewPerson(db *mysql.MySQL) (*Person, error) {

	err := db.Connect()
	if err != nil {
		err = common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return &Person{db: db}, err
}
