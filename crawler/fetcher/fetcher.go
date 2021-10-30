/*
 * @Author: wenbin
 * @Date: 2021-09-05 22:17:07
 * @LastEditTime: 2021-09-12 23:06:24
 * @Description:
 * @FilePath: /crawler/crawler/fetcher/fetcher.go
 */

package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	/**
	* @description: 判断并返回文本编码类型
	* @param r: io.Reader
	* @return : 文本编码类型
	 */
	types, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(types, "")
	return e
}

func Fetche(url string) ([]byte, error) {
	/**
	 * @description: 根据url获取相应的响应体信息
	 * @param url:
	 * @return (响应内容，error)
	 */
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// 判断网址文本编码类型，统一转换为utf-8格式
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}
