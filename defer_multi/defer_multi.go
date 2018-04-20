package main

import "fmt"

// defer延迟函数调用被压入一个栈中，当函数返回时，按照后进先出的顺序调用被延迟的函数调用

func main(){
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
