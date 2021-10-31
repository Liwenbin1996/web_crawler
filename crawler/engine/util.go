/*
 * @Author: wenbin
 * @Date: 2021-10-31 21:27:53
 * @LastEditTime: 2021-10-31 21:29:25
 * @Description:
 * @FilePath: /crawler/crawler/engine/util.go
 */
package engine

import (
	"log"
	"crawler/crawler/fetcher"
)

func worker(req Request) (ParseResult, error) {
	log.Printf("Fetching Url: %s", req.Url)
	body, err := fetcher.Fetche(req.Url)
	if err != nil {
		return ParseResult{}, err
	}

	parseResult := req.ParseFunc(body)
	return parseResult, nil
}

func resultHandle(results []interface{}) {
	for _, item := range results {
		log.Printf("Got item: %v\n", item)
	}
}