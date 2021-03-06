package main

import (
	"math"
	"fmt"
)

// 接口类型由一组方法定义的集合，接口类型的值可存放实现这些方法的任何值

type ABser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a ABser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 5}

	a = f
	a = &v



	fmt.Println(a.Abs())

}