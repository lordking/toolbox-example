package main

import (
	"flag"
	"reflect"

	"github.com/lordking/toolbox-example/http/person_mongo/benchmark/testcase"
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/log"
)

var (
	cfgFile    string
	methodName string
)

func init() {
	flag.StringVar(&cfgFile, "config", "", "config file.")
	flag.StringVar(&methodName, "m", "", "test case name")
	flag.Parse()
}

func main() {

	common.InitConfig("person", cfgFile)
	log.SetLogDefaults("log")

	if methodName == "" {
		log.Fatal("Not found testcase!")
	}

	s := &testcase.TestCase{}
	v := reflect.ValueOf(s)

	v.MethodByName(methodName).Call(nil)

}
