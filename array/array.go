package main

import "fmt"

// 类型[n]T 是一个有n个类型为T的值的数组
// 定义变量a是一个有是10个整数的数组
// 数组的长度 因为数组不能改变大小，这是一个制约

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0],a[1])
	fmt.Println(a)
}
