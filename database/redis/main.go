package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/redis"

	"github.com/lordking/toolbox/log"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {

	redis := redis.New()
	err := database.CreateInstance(redis, "./redis.json")
	defer common.CheckFatal(err)

	p := &Person{}
	form := &PersonForm{
		Name:  "leking",
		Phone: "18900000000",
	}

	err = p.Set(form)
	defer common.CheckFatal(err)

	f2, err := p.Get()
	defer common.CheckFatal(err)

	log.Debug("f2:%s", f2)
}
