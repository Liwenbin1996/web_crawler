/*
 * @Author: wenbin
 * @Date: 2021-09-05 23:14:09
 * @LastEditTime: 2021-10-31 21:28:28
 * @Description:
 * @FilePath: /crawler/crawler/engine/simple.go
 */
package engine

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	/**
	 * @description: 根据给定URL，递归爬取网页信息并进行解析
	 * @param seeds: 请求信息，包括URL和对应的解析方法
	 * @return
	 */

	var requests []Request
	requests = append(requests, seeds...)

	// 遍历任务列表
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		// 执行爬虫任务
		parseResult, err := worker(req)
		if err != nil {
			continue
		}

		// 处理爬虫返回结果
		resultHandle(parseResult.Items)

		requests = append(requests, parseResult.Requests...)
	}

}
