package main

import "fmt"

// 常量的定义和变量类似，使用const关键字 常量不能用 ：= 语法定义  常量可以是字符， 字符串， 布尔， 数字类型

const Pi = 3.14

func main(){
	const World = "世界"

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
