package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
)

const seedUrl = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main() {

	seed := engine.Request{
		Url:        seedUrl,
		ParserFunc: parser.ParseCityList,
	}
	engine.Run(seed)

}
