package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/redis"
	"github.com/lordking/toolbox/log"
)

type Reveiver struct{}

var (
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "config", "", "config file.")
}

func (d *Reveiver) GetPerson(obj *PersonVO) error {
	log.Debugf("Received a message: %s", common.PrettyObject(obj))

	return nil
}

func main() {

	common.InitConfig("redis_exmple", cfgFile)
	log.SetLogDefaults("log")

	obj := &PersonVO{
		Name:  "leking",
		Phone: "18900000000",
	}

	redis := redis.New()
	err := database.Configure("redis", redis)
	defer common.CheckFatal(err)

	receiver := &Reveiver{}

	p := NewPerson(redis, receiver)

	//设置或新增
	err = p.Set("leking", obj, 1000)
	defer common.CheckFatal(err)
	log.Debugf("set a person: %s", common.PrettyObject(obj))

	//获取
	obj, err = p.Get("leking")
	defer common.CheckFatal(err)
	log.Debugf("get a person: %s", common.PrettyObject(obj))

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
		obj.Phone = fmt.Sprintf("18%d", common.RandInt64(900000000, 999999999))
		log.Debugf("publish a `person`: %s", common.PrettyObject(obj))
		err = p.Publish("person", obj)
		defer common.CheckFatal(err)
		time.Sleep(1e9)
	}

}
