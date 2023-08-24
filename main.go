package main

import (
	"bufio"
	"errors"
	"strconv"

	//"encoding"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Printf("fetech error:%v", err)

	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, errors.New(strconv.Itoa(resp.StatusCode))
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func NewProxy() (*httputil.ReverseProxy, error) {
	targetHost := "http://127.0.0.1:8500"
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	return proxy, nil
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	proxy, err := NewProxy()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", ProxyRequestHandler(proxy))
	http.ListenAndServe(":8080", nil)
}
