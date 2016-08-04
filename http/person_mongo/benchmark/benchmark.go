package main

import (
	"flag"
	"reflect"

	"github.com/lordking/toolbox-example/http/person_mongo/benchmark/testcase"
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

	v.MethodByName(*methodName).Call(nil)

}
