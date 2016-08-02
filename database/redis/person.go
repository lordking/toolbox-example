package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/redis"
)

const (
	key = "person"
)

//Person 用户数据对象
type (
	Person struct{}

	PersonForm struct {
		Name  string `json:"name" bson:"name"`
		Phone string `json:"phone" bson:"phone"`
	}
)

func (p *Person) Set(form *PersonForm) error {

	//获取单例
	db := (database.Instance).(*redis.Redis)

	if err := db.Connect(); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	if err := db.SetObject(key, form); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	if err := db.Close(); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return nil
}

func (p *Person) Get() (*PersonForm, error) {

	//获取单例
	db := (database.Instance).(*redis.Redis)

	if err := db.Connect(); err != nil {
		return nil, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	form := &PersonForm{}
	if err := db.GetObject(form, key); err != nil {
		return nil, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	if err := db.Close(); err != nil {
		return nil, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return form, nil
}
