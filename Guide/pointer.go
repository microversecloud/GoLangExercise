package main

import "fmt"

func main() {
	//类型*T是指向T类型值的指针。其零值为nil
	//&操作符会生成一个指向其操作数的指针
	//操作符表示指针指向的底层值
	i, j := 42, 2701

	p := &i         //指针p，指向i
	fmt.Println(*p) //通过指针p读取i的值

	*p = 21        //通过指针p设置i的值
	fmt.Println(i) //查看i设置后的新值

	p = &j         //重新设置指针指向j
	*p = *p / 37   //通过指针p对j进行除法运算
	fmt.Println(j) //查看j的值
	fmt.Println(i) //查看i的值
}
