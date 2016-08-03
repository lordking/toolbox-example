package main

import (
	"flag"
	"reflect"
	"time"

	"github.com/lordking/toolbox-example/http/welcome/benchmark/testcase"
	"github.com/lordking/toolbox/log"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {

	methodName := flag.String("m", "", "test case name")
	flag.Parse()

	if *methodName == "" {
		log.Fatal("Not found testcase!")
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
