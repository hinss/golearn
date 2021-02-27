package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	// 获取推断的类型
	encoding := determineEncoding(resp.Body)

	// 解决不同http页面的编码问题
	utf8Reader := transform.NewReader(resp.Body, encoding.NewDecoder())
	bytes, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)
}

/**
从网页的前1024个bytes中推断 页面编码
*/
func determineEncoding(reader io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
