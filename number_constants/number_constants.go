package main

import "fmt"

// 数值常量 数值常量是高精度的值 ，如果未指定也是由于上下文决定其类型

const (
	Big 	= 1 << 100
	Small	= Big >> 99
)

func needInt(x int) int { return  x*10 + 1}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main(){
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

