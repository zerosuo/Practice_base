package main

import "fmt"

// channel 可以是带缓存的 为make提供第二个参数作为缓冲长度来初始化一个缓冲的channel

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
