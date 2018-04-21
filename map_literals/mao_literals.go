package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map语法和结构体语法相似 不过必须有键名

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.3999,
	},
	"Google": Vertex{
		37.42204, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}

