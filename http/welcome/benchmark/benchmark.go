package main

import (
	"flag"
	"reflect"
	"time"

	"github.com/lordking/toolbox-example/http/welcome/benchmark/testcase"
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/log"
)

var (
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "config", "", "config file.")
}

func main() {

	common.InitConfig("welcome", cfgFile)
	log.SetLogDefaults("log")

	methodName := flag.String("m", "", "test case name")
	flag.Parse()

	if *methodName == "" {
		log.Fatalf("Not found testcase!")
	}

	s := &testcase.TestCase{}
	v := reflect.ValueOf(s)

	for {

		for j := 0; j < 100; j++ {
			go v.MethodByName(*methodName).Call(nil)
		}

		time.Sleep(1 * time.Second)
	}

}
