package main

import "fmt"


// 函数可以没有参数 或者接受收多个参数 返回值也是一样道理
func add(x int, y int) int {
	return x + y
}

// 如果是两个或者多个函数命名参数都是同一类型则 只写最后一个，其他都可以忽略
func sum(x, y int) int {
	return x - y
}

// 多值返回 ,返回2个string类型值，注意的是 在go中交换2值就这么简单
func swap(x, y string) (string, string) {
	return y, x
}

func main(){
	fmt.Println(add(5, 6))
	fmt.Println(sum(6, 3))
	fmt.Println(swap("hello", "world"))
}
