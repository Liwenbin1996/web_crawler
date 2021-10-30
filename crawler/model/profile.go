/*
 * @Author: wenbin
 * @Date: 2021-10-04 10:18:20
 * @LastEditTime: 2021-10-04 10:26:08
 * @Description:
 * @FilePath: /crawler/crawler/model/profile.go
 */
package model

type Profile struct {
	Name          string
	Gender        string
	Age           int
	Height        int
	Weight        int
	Income        string
	Marriage      string
	Education     string
	Occupation    string
	Census        string // 户籍
	Constellation string // 星座
	House         string
	Car           string
}
