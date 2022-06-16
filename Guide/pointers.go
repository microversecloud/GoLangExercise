package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) { //指针接收者声明
	v.X = v.X * f //指针接收者v修改接收者指向的值。
	v.Y = v.Y * f //方法经常需要修改它的接收者，指针接收者比值接收者更常用
}

func (v Vertex) ScaleWithoutMark(f float64) {
	//值接收者v，只会对原始Vertex值副本进行操作。
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestWithMark() {
	v := Vertex{3, 4}
	v.Scale(10)          //此时执行后会产生结果：f=10,v.X=30,v.Y=40
	fmt.Println(v.Abs()) //平方根后返回50
}

func TestWithoutMark() {
	v := Vertex{3, 4}      //X=3,Y=4
	v.ScaleWithoutMark(10) //并不会对v任何修改，只会在ScaleWithoutMark方法的局部参数（原始数据的副本）进行修改
	//此时X=3,Y=4
	fmt.Println(v.Abs()) //平方根后等于5
}
func main() {

	TestWithMark()

	TestWithoutMark()
}
