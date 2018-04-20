package main

import (
	"math"
	"fmt"
)

//if 语句可以在条件之前执行一个简单的语句 这个语句定义的变量作用于仅仅在if的范围内有效

func pow(x, n, lim float64) float64{
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main(){
	fmt.Println(
		pow(3, 2, 10),
		pow(2, 3, 20),
	)
}

