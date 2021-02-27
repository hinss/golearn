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
	"regexp"
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
	printCityList(bytes)

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

func printCityList(bytes []byte) {
	// + 表示一个或者多个  * 表示0个或者多个
	// 加括号表示正则表达式的提取
	reg := regexp.MustCompile(`<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`)
	// -1 表示全部匹配
	matches := reg.FindAllSubmatch(bytes, -1)

	// [][]string 二维数组
	// 每一行代表一个整体的匹配，每行的每列代表正则括号提取的内容
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n",
			m[2], m[1])
	}

}
