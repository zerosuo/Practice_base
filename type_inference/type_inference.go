package main

import "fmt"

//定义一个变量但是不指定其类型时候，变量的类型由右边的值推导得出，但是的当右边包含未指名类型
// 的数字常量时，取决于常量的精度

func main(){
	v := 42
	fmt.Printf("v is of type %T\n", v)
}
