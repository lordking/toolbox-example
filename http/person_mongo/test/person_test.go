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

func TestCreate(t *testing.T) {

	url := kHost + "/person/new"

	var data = []byte(`{
		"name":"leking",
	  "phone": "18987871818"
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

func TestFind(t *testing.T) {

	url := kHost + "/person/leking"

	var data = []byte(`{}`)

	b, _ := ws.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := ws.RequestJSON("GET", url, data)
	if err != nil {
		t.Error(err)
	}

	s, _ := ws.PrettyJSON(result)
	t.Logf("Response: %s", s)

}

func TestUpdate(t *testing.T) {

	url := kHost + "/person/update/leking"

	var data = []byte(`{
		"phone": "18987871111"
	}`)

	b, _ := ws.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := ws.RequestJSON("PUT", url, data)
	if err != nil {
		t.Error(err)
	}

	s, _ := ws.PrettyJSON(result)
	t.Logf("Response: %s", s)

}

func TestDelete(t *testing.T) {

	url := kHost + "/person/delete/leking"

	var data = []byte(`{}`)

	b, _ := ws.PrettyJSON(data)
	t.Logf("Request: %s", b)

	result, err := ws.RequestJSON("DELETE", url, data)
	if err != nil {
		t.Error(err)
	}

	s, _ := ws.PrettyJSON(result)
	t.Logf("Response: %s", s)

}
