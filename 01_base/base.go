package main

import "fmt"

var (
	T int = 12
	K string = "K"
)

func main() {
	var a string = "Go"
	var b, c int = 1,2
	fmt.Println(a, b, c)


	var d  ="RUNOOB"

	var e int

	var f bool

	fmt.Println(d, e, f)


	g, h := 1.5, "asd"

	fmt.Println(g, h)


	fmt.Println(T, K)

	fmt.Printf("内存地址 T = %v\n", &T)

	i := T

	fmt.Printf("i = %v, address = %T\n", i, &i)

	j := 1
	k := j
	fmt.Println(&j, &k)
	j = 2
	fmt.Println(&j, &k, j, k)
}
