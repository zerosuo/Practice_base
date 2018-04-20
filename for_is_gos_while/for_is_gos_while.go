package main

import "fmt"

func main(){
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// 死循环
	for {
		a :=1
		a += a+1
	fmt.Println(a)
	}
}


