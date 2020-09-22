package main

import "fmt"

func main() {

	a := 60 // 00111100
	b := 13 // 00001101

	fmt.Println(a ^ b) // 0011 0001 = 49

	fmt.Println(a << 2) // 1111 0000 = 240

	c := a
	c <<= 2
	fmt.Println(c)

	var pointer *int = &a

	fmt.Println(pointer, *pointer, &a)



}
