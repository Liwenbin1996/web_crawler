/*
 * @Author: wenbin
 * @Date: 2021-09-05 23:14:09
 * @LastEditTime: 2021-09-05 23:47:03
 * @Description:
 * @FilePath: /crawler/crawler/engine/engine.go
 */
package engine

import (
	"crawler/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	/**
	 * @description: 根据给定URL，递归爬取网页信息并进行解析
	 * @param seeds: 请求信息，包括URL和对应的解析方法
	 * @return
	 */
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		log.Printf("Fetching Url: %s", req.Url)
		body, err := fetcher.Fetche(req.Url)
		if err != nil {
			log.Printf("Fetcher error: fetching url %s: %v", req.Url, err)
			continue
		}

		parseResult := req.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item: %v\n", item)
		}
	}

}
