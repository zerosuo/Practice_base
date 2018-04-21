package main

import "fmt"

// 一个slice 会指向一个序列的值， 并且包含了长度信息。
// []T 是一个元素类型为T的slice

func main() {
	p := []int{2, 3, 4, 5, 6, 9}
	fmt.Println("p===", p)

	for i := 0; i < len(p); i ++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}
