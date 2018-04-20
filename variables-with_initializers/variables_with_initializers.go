package main

import "fmt"

// 变量定义可以包含初始值，每个变量对应一个，如果初始化使用表达式可以省略类型，从初始值中获取类型

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no"
	fmt.Println(i, j, c, python, java)
}
