package main

import (
	ws "goutils"

	"flag"
	"os"
	"reflect"
	"time"

	"goutils-example/webseed/welcome/benchmark/testcase"
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

	for {

		for j := 0; j < 100; j++ {
			go v.MethodByName(*methodName).Call(nil)
		}

		time.Sleep(1 * time.Second)
	}

}
