package main

import (
	"fmt"
	crwl "webcrawler/src/crawler"
)

type Responsebody struct {
	Url  string
	Data string
}

func main() {
	fmt.Println("hello world")
	links := []string{"https://www.google.com", "https://www.youtube.com"}
	var resp []Responsebody
	for _, link := range links {
		data, _ := crwl.Crawl(link)
		resp = append(resp, Responsebody{
			Url:  link,
			Data: data,
		})
	}
	fmt.Println(resp)
}
