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

type Reveiver struct{}

func (d *Reveiver) GetPerson(obj *PersonVO) error {
	log.Debug("Receive a message: %s", common.PrettyObject(obj))

	return nil
}

func main() {

	obj := &PersonVO{
		Name:  "leking",
		Phone: "18900000000",
	}

	redis := redis.New()
	err := database.Configure(redis, "./redis.json")
	defer common.CheckFatal(err)

	receiver := &Reveiver{}

	p := NewPerson(redis, receiver)

	//设置或新增
	err = p.Set("leking", obj, 1000)
	defer common.CheckFatal(err)
	log.Debug("set a person: %s", common.PrettyObject(obj))

	//获取
	obj, err = p.Get("leking")
	defer common.CheckFatal(err)
	log.Debug("get a person: %s", common.PrettyObject(obj))

	// //删除
	err = p.Delete("leking")
	defer common.CheckFatal(err)
	log.Debug("delete a person")

	//订阅
	err = p.Subscribe("person")
	defer common.CheckFatal(err)
	log.Debug("subscribe `person`")

	//发布
	for {
		err = p.Publish("person", obj)
		defer common.CheckFatal(err)
		log.Debug("publish a `person`: %v", obj)
		time.Sleep(1e9)
	}

}
