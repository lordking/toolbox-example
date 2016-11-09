package testcase

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/http"
	"github.com/lordking/toolbox/log"
)

//RequestCreate hello接口的测试案例
func (t *TestCase) RequestCreate() {

	url := host + "/person/new"

	k := 0
	for {

		if k > limit {
			break
		}

		for j := 0; j < count; j++ {

			if k > limit {
				break
			}

			go func(t int) {

				str := fmt.Sprintf(`{
					"name":"leking%d",
					"phone":"189aaaa%d"
				}`, t, t)

				data := []byte(str)
				b, _ := common.PrettyJSON(data)
				log.Debugf("Request: %s", b)

				result, err := http.RequestJSON("POST", url, data)
				if err != nil {
					log.Errorf("Error: %s", err.Error())
				}

				s, _ := common.PrettyJSON(result)
				log.Debugf("Response: %s", s)

			}(k)

			k++
		}

		time.Sleep(1 * time.Second)
	}

}

//RequestFind hello接口的测试案例
func (t *TestCase) RequestFind() {

	for {

		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		for j := 0; j < count; j++ {

			go func() {

				url := fmt.Sprintf("%s/person/leking%d", host, r.Intn(limit))

				data := []byte(`{}`)
				b, _ := common.PrettyJSON(data)
				log.Debugf("Request: %s", b)

				result, err := http.RequestJSON("GET", url, data)
				if err != nil {
					log.Errorf("Error: %s", err.Error())
				}

				s, _ := common.PrettyJSON(result)
				log.Debugf("Response: %s", s)

			}()

		}

		time.Sleep(1 * time.Second)
	}

}

//RequestUpdate hello接口的测试案例
func (t *TestCase) RequestUpdate() {

	for {

		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		for j := 0; j < count; j++ {

			go func() {

				url := fmt.Sprintf("%s/person/update/leking%d", host, r.Intn(limit))

				str := fmt.Sprintf(`{
					"phone":"189bbbb%d"
				}`, r.Intn(limit))

				data := []byte(str)
				b, _ := common.PrettyJSON(data)
				log.Debugf("Request: %s", b)

				result, err := http.RequestJSON("PUT", url, data)
				if err != nil {
					log.Errorf("Error: %s", err.Error())
				}

				s, _ := common.PrettyJSON(result)
				log.Debugf("Response: %s", s)

			}()

		}

		time.Sleep(1 * time.Second)
	}

}

//RequestDelete hello接口的测试案例
func (t *TestCase) RequestDelete() {

	for {

		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		for j := 0; j < count; j++ {

			go func() {

				url := fmt.Sprintf("%s/person/delete/leking%d", host, r.Intn(limit))

				data := []byte(`{}`)
				b, _ := common.PrettyJSON(data)
				log.Debugf("Request: %s", b)

				result, err := http.RequestJSON("DELETE", url, data)
				if err != nil {
					log.Errorf("Error: %s", err.Error())
				}

				s, _ := common.PrettyJSON(result)
				log.Debugf("Response: %s", s)

			}()

		}

		time.Sleep(1 * time.Second)
	}

}
