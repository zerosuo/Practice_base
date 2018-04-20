package main

import "fmt"

// struct 结构体 就是一个字段的集合

type Vertex struct {
	X int
	Y int
}

// 通过结构体字段的值作为列表来新分配一个结构体
// &返回一个指向结构体的指针

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = &Vertex{1, 2}
	p = &Vertex{1, 2}
)
func main(){

	v := Vertex{1, 2}
	//v访问结构体中字段 是通过.号来直接访问的
	v.X = 4
	// 结构体字段可以通过结构体指针来访问，通过指针间接的访问是透明的
	p := &v
	p.X = 1e9
	fmt.Println(Vertex{1, 2})
	fmt.Println(v1, p, v2, v3)
}
