package main

import (
	"fmt"
	"go/constant"
)

func main() {

	_, n, s := numbers()
	fmt.Println(n, s)
	fmt.Println(constant.Unknown)
}

func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a,b,c
}
