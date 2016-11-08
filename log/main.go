package main

import (
	"fmt"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/log"
	"github.com/spf13/viper"
)

func init() {
	initConfig()
	log.SetLogDefaults("log")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Read config file error: ", err)
	}
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
