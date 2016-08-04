package main

import (
	"time"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/redis"
	"github.com/lordking/toolbox/log"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

type MainDelegate struct{}

func (d *MainDelegate) GetPerson(obj *PersonVO) error {
	log.Debug("Receive a message: %s", common.PrettyObject(obj))

	return nil
}

func main() {

	obj := &PersonVO{
		Name:  "leking",
		Phone: "18900000000",
	}

	redis := redis.New()
	err := database.CreateInstance(redis, "./redis.json")
	defer common.CheckFatal(err)

	delegate := &MainDelegate{}

	p := &Person{
		Delegate: delegate,
	}

	//设置或新增
	err = p.Set("leking", obj, 10)
	defer common.CheckFatal(err)

	//获取
	obj, err = p.Get("leking")
	defer common.CheckFatal(err)
	log.Debug("form: %s", common.PrettyObject(obj))

	//删除
	err = p.Delete("leking")
	defer common.CheckFatal(err)

	//订阅
	err = p.Subscribe("person")
	defer common.CheckFatal(err)

	//发布
	for {
		err = p.Publish("person", obj)
		defer common.CheckFatal(err)
		time.Sleep(1e9)
	}

}
