package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/log"
)

func init() {
	log.SetLogDefaults("./log.json")
}

func main() {
	log.Debug("This is a `test`")
	log.Debugf("This is a '%s'", "test")

	log.Info("This is a `test`")
	log.Infof("This is a '%s'", "test")

	log.Warn("This is a `test`")
	log.Warnf("This is a '%s'", "test")

	err := common.NewError(common.ErrCodeInternal, "a error message!")

	log.Error("This is a eror: ", err)
	log.Errorf("This is a eror: %s", err)

	log.Fatalf("This is a fatal: %s", err)

	log.Panicf("This is a panic: %s", err)
}
