package main


import (

	"fmt"
	"math/cmplx"
)

var (

	Tobe bool = false
	maxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(- 5 + 21i)

)


func main(){

	fmt.Printf("type: %T, value: %v\n", Tobe, Tobe)
	fmt.Printf("type: %T, value: %v\n", maxInt, maxInt)
	fmt.Printf("type: %T, value: %v\n", z, z)
}