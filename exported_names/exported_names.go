package main

import (

	"fmt"
	"math"
)

// 导入一个包，就可以用其导出的名称调用它  Go中，首字母大些是可以被导出的。注意！！！
func main(){
	fmt.Println(math.Pi)
}

