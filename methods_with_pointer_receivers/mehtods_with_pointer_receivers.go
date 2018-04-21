package main

import (
	"math"
	"fmt"
)

// 方法可以与命名类型或者命名类型的指针关联
// 避免在每个方法调用中拷贝值， 方法可以修改接收者的指定的参数
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 5}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
