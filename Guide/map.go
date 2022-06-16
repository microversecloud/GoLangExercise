// 题目描述：https://tour.go-zh.org/moretypes/23
// 题目分析：
// 1.使用strings.Fields获取字符串的分割信息
// 2.以位的形式返回[]string数组
// 3.计算字符串中单个单词出现的次数
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func wordCount(s string) map[string]int {
	m := make(map[string]int) //函数会返回给定类型的映射，并将其初始化备用
	c := strings.Fields(s)    //以[]byte形式返回

	for _, v := range c {
		m[v] += 1 //每出现相同的单词（字符串），出现后次数加1
	}
	return m
}

func main() {
	wc.Test(wordCount)
}
