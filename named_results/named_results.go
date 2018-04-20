package main

import "fmt"

// Go的返回值可以被命名，像变量那样使用 下面的sum x y 都是; 没有参数的 return 语句返回结果当前值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
