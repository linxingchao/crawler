package collect

import (
	"bufio"
	"fmt"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

type Fetcher interface {
	Get(url string) ([]byte, error)
}

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Printf("fetech error:%v", err)

	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
