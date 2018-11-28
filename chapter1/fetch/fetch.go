package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := "www.google.com"
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Http Status Code: %s\n", resp.Status)

	// b, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("%s", b)
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
