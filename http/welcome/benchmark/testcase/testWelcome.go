package testcase

import (
	ws "goutils"
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

	b, _ := ws.PrettyJSON(data)
	ws.LogDebug("Request: %s", b)

	result, err := ws.RequestJSON("POST", url, data)
	if err != nil {
		ws.LogError("Error: %q", err)
	}

	s, _ := ws.PrettyJSON(result)
	ws.LogDebug("Response: %s", s)
}
