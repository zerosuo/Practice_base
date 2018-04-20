package main

import (
	"math"
	"fmt"
)

// 还是i强调if的作用域

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(2,3, 20),
	)
}
