package test

/**
测试之前必须先启动http服务
*/

import (
	"testing"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/http"
)

func TestHello(t *testing.T) {

	url := kHost + "/welcome/hello"

	var data = []byte(`{
		"name":"leking",
	   "content": {
	   		"Ye":"You are welcome"
	   }
	}`)

	b, _ := common.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := http.RequestJSON("POST", url, data)
	if err != nil {
		t.Error(err)
	}

	s, _ := common.PrettyJSON(result)
	t.Logf("Response: %s", s)

}
