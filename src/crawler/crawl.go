package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Responsebody struct {
	Url  string
	Data string
}

func crawl(api string) (string, error) {
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

func Webcrawler() ([]Responsebody, error) {
	var resp []Responsebody

	links := []string{"https://www.google.com", "https://en.wikipedia.org/wiki/Ant-Man_(film)"}
	for _, link := range links {
		data, _ := crawl(link)
		resp = append(resp, Responsebody{
			Url:  link,
			Data: data,
		})
	}
	if len(resp) == 0 {
		return resp, fmt.Errorf("none of the data can be fetched")
	}

	return resp, nil
}
