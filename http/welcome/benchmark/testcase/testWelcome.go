package testcase

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/http"
	"github.com/lordking/toolbox/log"
)

//RequestHello hello接口的测试案例
func (t *TestCase) RequestHello() {

	url := host + "/welcome/hello"

	var data = []byte(`{
		"name":"leking",
	   "content": {
	   		"Ye":"You are welcome"
	   }
	}`)

	b, _ := common.PrettyJSON(data)
	log.Debug("Request: %s", b)

	result, err := http.RequestJSON("POST", url, data)
	if err != nil {
		log.Error("Error: %s", err.Error())
	}

	s, _ := common.PrettyJSON(result)
	log.Debug("Response: %s", s)
}
