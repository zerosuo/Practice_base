package main

import "fmt"

// slice 重新切片，创建一个新的slice值指向相同的数组


func main() {
	p := []int {2, 3, 4, 6, 8, 2,}

	fmt.Println("p==", p)
	fmt.Println("p[1:4]==", p[1:4])

	// 省略下标代表从0开始
	fmt.Println("p[:3] ==", p[:3])
	fmt.Println("p[4:] ==", p[4:])
}
