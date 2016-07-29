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
	  "name":"sunny",
	  "phone":"025-216549778"
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

	url := kHost + "/person/sunny"

	result, err := ws.RequestJSON("GET", url, nil)
	if err != nil {
		t.Error(err)
	}

	s, _ := ws.PrettyJSON(result)
	t.Logf("Response: %s", s)
}

func TestUpdate(t *testing.T) {

	url := kHost + "/person/update/sunny"

	var data = []byte(`{
	  "phone":"025-216549779"
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

	url := kHost + "/person/delete/sunny"

	result, err := ws.RequestJSON("DELETE", url, nil)
	if err != nil {
		t.Error(err)
	}

	s, _ := ws.PrettyJSON(result)
	t.Logf("Response: %s", s)

}
