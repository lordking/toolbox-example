package main

import (
	"github.com/lordking/toolbox-example/http/blog/cmd"
	"github.com/lordking/toolbox/common"
)

func main() {
	common.ConfigRuntime()
	cmd.Execute()
}
