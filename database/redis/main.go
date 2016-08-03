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

func (d *MainDelegate) GetPerson(form *PersonForm) error {
	log.Debug("form: %s", common.PrettyObject(form))

	return nil
}

func main() {

	form := &PersonForm{
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
	err = p.Set("leking", form, 10)
	defer common.CheckFatal(err)

	//获取
	form, err = p.Get("leking")
	defer common.CheckFatal(err)
	log.Debug("form: %s", common.PrettyObject(form))

	//删除
	err = p.Delete("leking")
	defer common.CheckFatal(err)

	//订阅
	err = p.Subscribe("person")
	defer common.CheckFatal(err)

	//发布
	for {
		err = p.Publish("person", form)
		defer common.CheckFatal(err)
		time.Sleep(1e9)
	}

}
