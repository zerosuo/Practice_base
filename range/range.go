package main

import "fmt"

// range for循环的range格式可以对slice或者map进行迭代循环

var pow = []int {1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	sum := make([]int, 10)
	for i := range sum {
		sum[i] = 1 << uint(i)
	}

	// 可以通过赋值给——来忽略序号和值，如果只需要索引值，去掉 value的部分即可

	for _, value := range sum {
		fmt.Printf("%d\n", value)
	}
}
