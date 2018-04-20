package main

import "fmt"

// 简短声明， 在函数中：= 简洁赋值语句在明确类型的地方，可以用于替代var定义（在函数外一定不能使用）

func main(){
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no"

	fmt.Println(i, j, k, c, python, java)
}
