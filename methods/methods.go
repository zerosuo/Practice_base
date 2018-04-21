package main

import (
	"math"
	"fmt"
)

// Go没有类但是仍然可以在结构体类型上定义方法。方法接收者出现在func关键字和方法名之间的参数中
// 包中任何类型都可以定义任何方法不仅仅是结构体 ，但是不能呢个对其他包的类型或者基础类型定义方法
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abb() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abb())
}
