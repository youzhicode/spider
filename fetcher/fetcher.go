package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"spider/iencoding"
)

func Fetcher(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status code: %v", resp.StatusCode)
	}
	// 如果有乱码问题
	bufioReader := bufio.NewReader(resp.Body)
	e := iencoding.DetermineEncoding(bufioReader)
	utf8Reader := transform.NewReader(bufioReader, e.NewDecoder())
	all, err := io.ReadAll(utf8Reader)
	if err != nil {
		return nil, err
	}
	return all, nil
}
