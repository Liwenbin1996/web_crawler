/*
 * @Author: wenbin
 * @Date: 2021-10-04 10:07:34
 * @LastEditTime: 2021-10-10 22:45:08
 * @Description:
 * @FilePath: /crawler/crawler/zhenai/parser/profile.go
 */
package parser

import (
	"crawler/crawler/engine"
	"crawler/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)元</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var censusRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var constellationRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)

func PraseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{
		Name:          name,
		Age:           extractInt(contents, ageRe),
		Height:        extractInt(contents, heightRe),
		Income:        extractString(contents, incomeRe),
		Marriage:      extractString(contents, marriageRe),
		Education:     extractString(contents, educationRe),
		Occupation:    extractString(contents, occupationRe),
		Census:        extractString(contents, censusRe),
		Gender:        extractString(contents, genderRe),
		Weight:        extractInt(contents, weightRe),
		House:         extractString(contents, houseRe),
		Car:           extractString(contents, carRe),
		Constellation: extractString(contents, constellationRe),
	}

	return engine.ParseResult{
		Items: []interface{}{profile},
	}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	/**
	 * @description:
	 * @param {*}
	 * @return {*}
	 */
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	/**
	 * @description:
	 * @param {*}
	 * @return {*}
	 */
	numStr := extractString(contents, re)
	number, err := strconv.Atoi(numStr)
	if err != nil {
		return 0
	}
	return number
}
