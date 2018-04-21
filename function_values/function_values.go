package main

import (
	"math"
	"fmt"
)

// 函数也可以作为值

func main(){
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))
}
