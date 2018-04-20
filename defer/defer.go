package main

import "fmt"

// defer 语句延迟函数的执行直到上层函数返回
// 延迟调用的参数会立即生成，但是上次函数返回前函数都不会调用

func main(){
	defer fmt.Println("world")

	fmt.Println("hello")
}
