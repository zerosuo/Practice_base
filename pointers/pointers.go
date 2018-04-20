package main

import "fmt"

// Go具有指针。指针保存了变量的内存地址。
// *T 是指向 T 的值的指针 。零值是nil
// & 符号会生成一个指向其作用对象的指针  // 间接引用和直接引用

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
