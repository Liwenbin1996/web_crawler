/*
 * @Author: wenbin
 * @Date: 2021-09-05 22:39:24
 * @LastEditTime: 2021-09-05 23:04:38
 * @Description:
 * @FilePath: /crawler/crawler/engine/types.go
 */
package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
