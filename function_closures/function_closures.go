package main

import "fmt"

// 闭包是一个函数值，它来自函数体的外部的变量引用。函数可以对这个引用值进行访问和赋值；
// 这个函数被绑定在这个变量上了

func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := addr(), addr()
	for i := 0; i < 10; i ++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}