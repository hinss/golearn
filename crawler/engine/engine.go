package engine

import (
	"learngo/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {

	// 装载Request请求的队列
	var requests []Request

	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	// 当队列中有请求时，不断拿出第一个执行
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetcher(r.Url)
		if err != nil {
			// 遇到请求错误就继续下一个请求
			log.Printf("Fetcher error fetching url %s %v",
				r.Url, err)
			continue
		}

		// 拿出请求的解析器解析获得结果
		parseResult := r.ParserFunc(body)

		// 把请求添加到 处理请求队列中
		requests = append(requests, parseResult.Requests...)

		// 打印结果
		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}

}
