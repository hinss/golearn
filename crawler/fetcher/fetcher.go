package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetcher(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d",
			resp.StatusCode)
	}

	// 获取推断的类型
	bufioReader := bufio.NewReader(resp.Body)
	encoding := determineEncoding(bufioReader)

	// 解决不同http页面的编码问题
	utf8Reader := transform.NewReader(bufioReader, encoding.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

/**
从网页的前1024个bytes中推断 页面编码
*/
func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)
	if err != nil {
		log.Panicf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
