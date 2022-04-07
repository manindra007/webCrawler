package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Crawl(api string) (string, error) {
	resp, err := http.Get(api)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(body), nil
}
