package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/redis"
	"github.com/lordking/toolbox/log"
)

//Person 用户数据对象
type (
	PersonDelegate interface {
		GetPerson(obj *PersonVO) error
	}

	Person struct {
		Delegate PersonDelegate
	}

	PersonVO struct {
		Name  string `json:"name" bson:"name"`
		Phone string `json:"phone" bson:"phone"`
	}
)

func (p *Person) Set(key string, obj *PersonVO, expire int) error {

	//获取单例
	db := (database.Instance).(*redis.Redis)

	if err := db.Connect(); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	if err := db.SetObject(key, obj, expire); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	if err := db.Close(); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return nil
}

func (p *Person) Get(key string) (*PersonVO, error) {

	obj := new(PersonVO)

	//获取单例
	db := (database.Instance).(*redis.Redis)

	if err := db.Connect(); err != nil {
		return nil, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	if err := db.GetObject(obj, key); err != nil {
		return nil, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	if err := db.Close(); err != nil {
		return nil, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return obj, nil
}

func (p *Person) Delete(key string) error {

	//获取单例
	db := (database.Instance).(*redis.Redis)

	if err := db.Connect(); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	if err := db.DeleteObject(key); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	if err := db.Close(); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return nil
}

func (p *Person) Publish(channel string, obj *PersonVO) error {

	//获取单例
	db := (database.Instance).(*redis.Redis)

	if err := db.Connect(); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	if err := db.PublishObject(channel, obj); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	if err := db.Close(); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return nil
}

func (p *Person) Subscribe(channel string) error {

	//获取单例
	db := (database.Instance).(*redis.Redis)

	if err := db.Connect(); err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	psc, err := db.Subscribe(channel)
	if err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	db.Receive(psc)

	go func() {

		for {
			data := <-db.ReceiveQueue

			if p.Delegate != nil {
				obj := new(PersonVO)
				common.ReadJSON(obj, data)
				if err := p.Delegate.GetPerson(obj); err != nil {
					log.Error("Receive error:", err)
				}
			}

		}

	}()

	return nil
}
