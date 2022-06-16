// 题目：https://tour.go-zh.org/moretypes/18
// 题目分析：
// 1.外层切片长度为dy
// 2.内层切片长度为dx
// 3.内层切片的每个元素值可选题目中任意一个计算公式，譬如：x%(y+1)
// 4.使用嵌套循环方式计算颜色值
package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy) //外层切片
	for x := range a {
		b := make([]uint8, dx) //内层切片
		for y := range b {
			b[y] = uint8(x % (y + 1)) //给内层切片里的每一个元素赋值。其中函数课题换成其他的函数展示不同的图形
		}
		a[x] = b //给外层切片里的每一个元素赋值
	}
	return a
}

func main() {
	pic.Show(Pic)
}
