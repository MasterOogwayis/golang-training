package main

import (
	"fmt"
	"unsafe"
)

const b = "String"

const (
	Unknown = 0
	Female = 1
	Male = 2
)

const (
	A = "abc"
	B = len(A)
	C = unsafe.Sizeof(A)
)

const (
	X = iota
	Y = iota
	Z
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	fmt.Println(a, b, c)


	fmt.Println(A, B, C)

	fmt.Println(X, Y, Z)
}
