package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	// var URLS []string
	// URLS = append(os.Args[1:], os.Args[1:]...)
	// for _, url := range URLS {
	//	....
	//	}

	for _, url := range os.Args[1:] {
		for i := 0; i < 2; i++ {
			go fetch(url, ch)
		}
	}

	for range os.Args[1:] {
		for i := 0; i < 2; i++ {
			fmt.Println(<-ch)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("url:%s, error:%s", url, err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("url:%s, reading error:%s", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
