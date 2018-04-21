package main

import "fmt"

// slice 的零值是nil， 一个nil的slice的长度和容量是0
// 向slice添加元素是一种常见的操作，go提供了一个内建函数append 。

func main(){
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	var a []int
	printSlice("a", a)

	// append on nil slices
	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n," ,
		s, len(x), cap(x), x)
}
