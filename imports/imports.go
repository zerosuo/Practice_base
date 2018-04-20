package main

import (
	"fmt"
	"math"
)
// 如果需要倒入多个包则可以写成如上，注意：包导入必须要用到，如果只是初始化则前面需要加下划线_只导入不使用
func main(){
	fmt.Print("Now you hava %g problems", math.Nextafter(2, 3))
}



