package main

import (
	"fmt"
	"time"
	"net/http"
	"log"
)

func main() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)
		}(url)
	}
	time.Sleep(time.Second)
}
