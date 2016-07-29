package main

import "github.com/lordking/toolbox/log"

func main() {

	log.SetLevel(log.DebugLevel)

	log.Debug("This is a '%s'", "test")

	log.Info("This is a '%s'", "test")

	log.Warn("This is a '%s'", "test")

	log.Error("This is a '%s'", "test")

	log.Fatal("This is a '%s'", "test")

	log.Panic("This is a '%s'", "test")

}
