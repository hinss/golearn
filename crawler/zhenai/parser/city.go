package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {

	reg := regexp.MustCompile(cityRe)
	matches := reg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User: "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// 使用匿名函数，函数式编程方式
			ParserFunc: func(b []byte) engine.ParseResult {
				return ParseProfile(b, name)
			},
		})
	}

	return result
}
