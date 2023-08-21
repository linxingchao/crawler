package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.thepaper.cn/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("fetch url err: %v", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read content err: %v", err)
	}
	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links\n", numLinks)
}
