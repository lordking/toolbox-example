package models

import (
	ws "goutils"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Id    int64  `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`
}

func (p *Person) Create() error {
	DataSourceInstance := ws.DataSourceInstance.(*(ws.MySQL))
	db, err := DataSourceInstance.GetConnection()
	mydb := db.(*sql.DB)

	if err != nil {
		return ws.ToError(ws.ErrCodeInternal, err)
	}

	stmt, err := mydb.Prepare("INSERT INTO person(name, phone) VALUES(?, ?)")
	defer stmt.Close()

	if err != nil {
		return ws.ToError(ws.ErrCodeInternal, err)
	}

	res, err := stmt.Exec(p.Name, p.Phone)

	if err != nil {
		return ws.ToError(ws.ErrCodeInternal, err)
	} else {
		p.Id, err = res.LastInsertId()
		if err != nil {
			return ws.ToError(ws.ErrCodeInternal, err)
		}

	}

	return nil
}

func (p *Person) Find(name string) ([]Person, error) {
	DataSourceInstance := ws.DataSourceInstance.(*(ws.MySQL))
	db, err := DataSourceInstance.GetConnection()
	mydb := db.(*sql.DB)

	if err != nil {
		return nil, ws.ToError(ws.ErrCodeInternal, err)
	}

	//查询
	var result []Person

	rows, err := mydb.Query("select id, name, phone from person where name = ?", name)

	if err != nil {
		return nil, ws.ToError(ws.ErrCodedParams, err)
	}

	defer rows.Close()
	i := 0
	for rows.Next() {
		var person Person

		err := rows.Scan(&(person.Id), &(person.Name), &(person.Phone))
		if err != nil {
			ws.ToError(ws.ErrCodedParams, err)
		}

		result = append(result, person)
		i++
	}

	err = rows.Err()
	if err != nil {
		return nil, ws.ToError(ws.ErrCodeInternal, err)
	}

	return result, nil
}

func (p *Person) Update(name string) (int64, error) {
	DataSourceInstance := ws.DataSourceInstance.(*(ws.MySQL))
	db, err := DataSourceInstance.GetConnection()
	mydb := db.(*sql.DB)

	if err != nil {
		return 0, ws.ToError(ws.ErrCodeInternal, err)
	}

	stmt, err := mydb.Prepare("UPDATE person SET phone=? where name=?")
	defer stmt.Close()

	if err != nil {
		return 0, ws.ToError(ws.ErrCodeInternal, err)
	}

	res, err := stmt.Exec(p.Phone, name)

	if err != nil {
		return 0, ws.ToError(ws.ErrCodeInternal, err)
	} else {
		num, err := res.RowsAffected()
		if err != nil {
			return 0, ws.ToError(ws.ErrCodeInternal, err)
		}
		return num, nil
	}
}

func (p *Person) Delete(name string) (int64, error) {
	DataSourceInstance := ws.DataSourceInstance.(*(ws.MySQL))
	db, err := DataSourceInstance.GetConnection()
	mydb := db.(*sql.DB)

	if err != nil {
		return 0, ws.ToError(ws.ErrCodeInternal, err)
	}

	stmt, err := mydb.Prepare("DELETE FROM person WHERE name=?")
	defer stmt.Close()

	if err != nil {
		return 0, ws.ToError(ws.ErrCodeInternal, err)
	}

	res, err := stmt.Exec(name)

	if err != nil {
		return 0, ws.ToError(ws.ErrCodeInternal, err)
	} else {
		num, err := res.RowsAffected()
		if err != nil {
			return 0, ws.ToError(ws.ErrCodeInternal, err)
		}
		return num, nil
	}
}
