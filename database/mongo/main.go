package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database/mongo"
)

func main() {

	config = &mongo.Config{}
	if configPath != "" {
		data, err = common.GetFileData(configPath)
		common.CheckFatal(err)

		err = ReadJSON(config, data)
		common.CheckFatal(err)
	}
	err = ReadOSEnv(config)
	common.CheckFatal(err)
}
