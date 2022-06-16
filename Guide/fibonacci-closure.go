// 题目描述：https://tour.go-zh.org/moretypes/26
// 题目分析：
// 1.实现一个fibonacci函数
// 2.使用该函数返回一个比吧
// 3.该闭包返回一个斐波那契额数列
// 4.闭包函数引用函数体外的值，并对其进行修改
package main

import "fmt"

func fibonacci() func() int {
	backA, backB := 0, 1
	return func() int {
		temp := backA
		backA, backB = backB, backA+backB
		return temp
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
