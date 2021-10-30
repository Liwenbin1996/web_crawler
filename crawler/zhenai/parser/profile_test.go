/*
 * @Author: wenbin
 * @Date: 2021-10-04 11:03:24
 * @LastEditTime: 2021-10-10 22:32:30
 * @Description:
 * @FilePath: /crawler/crawler/zhenai/parser/profile_test.go
 */
package parser

import (
	"crawler/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	fileContent, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := PraseProfile(fileContent, "高冷绅士小仙女")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Age:           74,
		Height:        190,
		Weight:        277,
		Income:        "1-2000",
		Gender:        "女",
		Name:          "高冷绅士小仙女",
		Constellation: "天蝎座",
		Occupation:    "销售",
		Marriage:      "离异",
		House:         "租房",
		Census:        "武汉市",
		Education:     "小学",
		Car:           "有车",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
