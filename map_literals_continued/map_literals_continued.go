package main

import "fmt"

// 如果顶级类型只有类型名的话，可以在文法的元素中省略键名 ---语句有点不通？

type Vertex struct {
	Lat, Long float64
}
// 创建一个key --》string ， Value --->struct
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.3996},
	"Google": {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}