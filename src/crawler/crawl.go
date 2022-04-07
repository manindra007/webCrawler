package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Responsebody for crawler request.
type Responsebody struct {
	Url  string
	Data string
}

//crawl: function to crawl over link
//input: link
//output: Responsebody,error
func crawl(link string) (string, error) {
	resp, err := http.Get(link)
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

//Webcrawler: function to handle multiple links
//input: string array of links.
//output: responsebody arr and error.
func Webcrawler(links []string) ([]Responsebody, error) {
	var resp []Responsebody

	for _, link := range links {
		data, err := crawl(link)
		if err != nil {
			continue
		}
		resp = append(resp, Responsebody{
			Url:  link,
			Data: data,
		})
	}
	if len(resp) == 0 {
		return resp, fmt.Errorf("none of the data can be fetched")
	} else if len(resp) < len(links) {
		return resp, fmt.Errorf("some of the data can be fetched")
	}

	return resp, nil
}
