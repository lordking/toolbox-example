package main

import (
	ws "goutils"

	"flag"
	"os"
	"reflect"

	"goutils-example/webseed/person_mysql/benchmark/testcase"
)

func init() {
	ws.SetLogger(ws.DebugLevel)
}

var methodName = flag.String("m", "", "test case name")

func main() {
	flag.Parse()

	if *methodName == "" {
		ws.LogError("Not found testcase!")
		os.Exit(0)
	}

	s := &testcase.TestCase{}
	v := reflect.ValueOf(s)

	v.MethodByName(*methodName).Call(nil)

}
