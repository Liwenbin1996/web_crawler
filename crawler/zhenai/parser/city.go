/*
 * @Author: wenbin
 * @Date: 2021-09-13 00:10:44
 * @LastEditTime: 2021-10-10 23:08:54
 * @Description:
 * @FilePath: /crawler/crawler/zhenai/parser/city.go
 */
package parser

import (
	"crawler/crawler/engine"
	"fmt"
	"regexp"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)">([^<]*)</a>`

func ParseCity(context []byte) engine.ParseResult {
	/**
	 * @description: 解析获取所有城市的网址
	 * @param context: 网站响应体信息
	 * @return 解析结果，包括所有城市的网址和位置信息
	 */
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(context, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		fmt.Println("User: " + name)
		result.Items = append(result.Items, "User: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return PraseProfile(c, name)
			},
		})
	}
	return result
}
