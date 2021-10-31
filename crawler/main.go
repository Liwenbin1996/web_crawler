/*
 * @Author: wenbin
 * @Date: 2021-09-05 09:33:58
 * @LastEditTime: 2021-11-01 00:44:48
 * @Description: 网络爬虫
 * @FilePath: /crawler/crawler/main.go
 */
package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/scheduler"
	"crawler/crawler/zhenai/parser"
)

const siteUrl = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:       siteUrl,
		ParseFunc: parser.ParseCityList,
	})
}
