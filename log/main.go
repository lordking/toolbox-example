package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/log"
)

func main() {

	log.SetLevel(log.DebugLevel)

	log.Debug("This is a '%s'", "test")

	log.Info("This is a '%s'", "test")

	log.Warn("This is a '%s'", "test")

	err := common.NewError(common.ErrCodeInternal, "a error message!")
	log.Error("This is a eror: %s", err)

	log.Fatal("This is a fatal: %s", err)

	log.Panic("This is a panic: %s", err)

}
