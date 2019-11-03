package gin

import (
	"dashboard/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestApiPing(t *testing.T) {
	pingUrl := "http://" + config.Conf.ApiAddress + "/ping"
	resp, _ := http.Get(pingUrl)
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Get request result:" + string(body))
}
func TestApiUserRoutes(t *testing.T) {
	pingUrl := "http://" + config.Conf.ApiAddress + "/user/routes?token=admin-token"
	resp, _ := http.Get(pingUrl)
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Get request result:" + string(body))
}
