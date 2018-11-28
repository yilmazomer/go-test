package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	var urls [3]string
	urls[0] = "https://www.google.com"
	urls[1] = "https://www.yahoo.com"
	urls[2] = "https://www.outlook.com"

	ch := make(chan string)
	for _, url := range urls[0:] {
		//fmt.Printf("url: %s\n", url)
		go fetch(url, ch)
	}

	for range urls[0:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("Total elapsed time: %.3fs seconds\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	receivedBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("error while reading %s: %v", url, err)
		return
	}

	ch <- fmt.Sprintf("%s %7d bytes %.3fs", url, receivedBytes, time.Since(start).Seconds())
}
