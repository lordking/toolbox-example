package test

/**
测试之前必须先启动http服务
*/

import (
	"testing"

	ws "goutils"
)

func init() {

}

func TestHello(t *testing.T) {

	url := kHost + "/welcome/hello"

	var data = []byte(`{
		"name":"leking",
	   "content": {
	   		"Ye":"You are welcome"
	   }
	}`)

	b, _ := ws.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := ws.RequestJSON("POST", url, data)
	if err != nil {
		t.Error(err)
	}

	s, _ := ws.PrettyJSON(result)
	t.Logf("Response: %s", s)

}
