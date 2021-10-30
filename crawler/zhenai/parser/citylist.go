/*
 * @Author: wenbin
 * @Date: 2021-09-05 22:50:42
 * @LastEditTime: 2021-09-13 00:21:53
 * @Description:
 * @FilePath: /crawler/crawler/zhenai/parser/citylist.go
 */
package parser

import (
	"crawler/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]+>([^<]+)</a>`

func ParseCityList(context []byte) engine.ParseResult {
	/**
	 * @description: 解析获取所有城市的网址
	 * @param context: 网站响应体信息
	 * @return 解析结果，包括所有城市的网址和位置信息
	 */
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(context, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City: " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
