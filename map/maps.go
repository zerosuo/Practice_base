package main

import "fmt"

// map 映射键到值；map在使用之前必须make而不是new来创建 ；值为nil的mao是空的，并且不能赋值

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main(){
	m = make(map[string]Vertex)
	m["Nell Labs"] = Vertex{
		40.68433, -74.399967,
	}
	fmt.Println(m["Bell Labs"])
}
