package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe 	bool			= false
	MaxInt  uint64			= 1<<64-1
	z 		complex128		= cmplx.Sqrt(-5 + 12i)
)

// Go的基本类型：bool string int int8 int16 int32 int64  complex64 complex128
//			   uint uint8 uint16 uint64 uintptr byte rune float32 float64

func main(){
	const f = "%T(%V)\n"

	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
