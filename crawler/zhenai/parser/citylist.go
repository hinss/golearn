package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityListParseRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

func ParseCityList(contents []byte) engine.ParseResult {

	// + 表示一个或者多个  * 表示0个或者多个
	// 加括号表示正则表达式的提取
	reg := regexp.MustCompile(cityListParseRe)
	// -1 表示全部匹配
	matches := reg.FindAllSubmatch(contents, -1)

	// [][]string 二维数组
	// 每一行代表一个整体的匹配，每行的每列代表正则括号提取的内容
	result := engine.ParseResult{}
	for _, m := range matches {
		// Items 装城市的名字
		result.Items = append(result.Items, "City: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
