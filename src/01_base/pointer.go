package main

import (
	"fmt"
	"reflect"
)

/**
 * 指针
 */
func main() {
	var a int = 20
	var ip *int

	ip = &a

	fmt.Println(ip, *ip)

	var ptr *int

	if (ptr == nil) {
		fmt.Println("is")
	} else {
		fmt.Println("is not")
	}


	//testPointer()
}

func testPointer() {
	var number int = 10

	var pointer *int = &number

	fmt.Println("a's address = ", pointer)

	fmt.Println("a = ", *pointer)

	fmt.Printf("%T", pointer)

	t := reflect.TypeOf(pointer)

	fmt.Printf("%T", t)

	fmt.Println(t.String())
}
