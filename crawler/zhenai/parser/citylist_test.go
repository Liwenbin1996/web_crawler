/*
 * @Author: wenbin
 * @Date: 2021-09-12 22:52:34
 * @LastEditTime: 2021-09-13 00:22:19
 * @Description:
 * @FilePath: /crawler/crawler/zhenai/parser/citylist_test.go
 */
package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	context, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	results := ParseCityList(context)

	const resultSize = 470
	expectUrls := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}
	expectCitys := []string{
		"City: 阿坝", "City: 阿克苏", "City: 阿拉善盟",
	}

	if len(results.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(results.Requests))
	}

	if len(results.Items) != resultSize {
		t.Errorf("result should have %d items; but had %d", resultSize, len(results.Items))
	}

	for i, url := range expectUrls {
		if results.Requests[i].Url != url {
			t.Errorf("expect url #%d: %s, but was %s", i, url, results.Requests[i].Url)
		}
	}

	for i, city := range expectCitys {
		if results.Items[i].(string) != city {
			t.Errorf("expect city #%d: %s, but was %s", i, city, results.Items[i].(string))
		}
	}
}
