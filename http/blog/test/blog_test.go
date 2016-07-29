package test

/**
测试之前必须先启动http服务
*/

import (
	"encoding/base64"
	"testing"

	ws "goutils"
)

var auths string

func init() {
	b := token + ":123"
	s := base64.StdEncoding.EncodeToString([]byte(b))
	auths = "Basic " + s
}

func TestCreate(t *testing.T) {

	url := host + "/blog/new"

	var data = []byte(`{
   "subject":"标题",
   "blog":"测试内容xxx",
   "author":"leking"
   }`)

	b, _ := ws.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := ws.RequestJSON("POST", url, data, ws.Header{"Authorization", auths})
	if err != nil {
		t.Error(err)
	}

	s, _ := ws.PrettyJSON(result)
	t.Logf("Response: %s", s)
}

func TestFind(t *testing.T) {

	url := host + "/blog/0/10"

	var data = []byte(`{}`)
	b, _ := ws.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := ws.RequestJSON("GET", url, data, ws.Header{"Authorization", auths})
	if err != nil {
		t.Error(err)
	}

	s, _ := ws.PrettyJSON(result)
	t.Logf("Response: %s", s)
}

func TestUpdate(t *testing.T) {

	url := host + "/blog/update/" + update_id

	var data = []byte(`{
   "subject":"标题",
   "blog":"测试内容xxx"
   }`)

	b, _ := ws.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := ws.RequestJSON("PUT", url, data, ws.Header{"Authorization", auths})
	if err != nil {
		t.Error(err)
	}

	s, _ := ws.PrettyJSON(result)
	t.Logf("Response: %s", s)
}

func TestDelete(t *testing.T) {

	url := host + "/blog/delete/" + delete_id

	var data = []byte(`{}`)
	b, _ := ws.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := ws.RequestJSON("DELETE", url, data, ws.Header{"Authorization", auths})
	if err != nil {
		t.Error(err)
	}

	s, _ := ws.PrettyJSON(result)
	t.Logf("Response: %s", s)
}
