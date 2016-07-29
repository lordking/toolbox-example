package test

import (
	"testing"

	ws "goutils"
)

func TestLogin(t *testing.T) {

	url := host + "/user/login"

	data := []byte(`{
		"username": "admin",
		"password": "admin"
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
