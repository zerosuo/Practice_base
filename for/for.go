package main

import "fmt"

//Go只有一种循环结构 for 循环 ； 基本的for循环除了没有 ()外 而{}是必须有的

func main(){
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

/*
// 可以让前置或者后置语句未空
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

}


*/
