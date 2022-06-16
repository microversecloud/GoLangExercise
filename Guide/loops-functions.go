// 题目原文链接地址：https://tour.go-zh.org/flowcontrol/8
// 题目分析：
// 1.定义一个浮点值并对其进行初始化
// 2.套用题目中提供的公式：z -= (z*z - x) / (2*z)，化简公式可得：	z = (z + x/z) / 2
package main

import (
	"fmt"
	"math"
)

func getAbs(x float64) float64 {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

func sqrt(x float64) float64 {
	z := 1.0
	if x < 0 {
		return 0
	} else if x == 0 {
		return 0
	} else {
		for getAbs(z*z-x) > 1e-6 {
			z = (z + x/z) / 2
		}
		return z
	}
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Sqrt(2)-sqrt(2) < 1e-6)
}
