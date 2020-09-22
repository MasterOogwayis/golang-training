package main

import "fmt"

func main() {
	var str string = "Hello"
	c := []byte(str)
	c[0] = 'c'
	str = string(c)
	fmt.Println(str)

	b := c[:3]

	fmt.Printf("%s\n", b)

	var t string = `
	Hello
	World
	.
	This is a test
	.
	`
	fmt.Println(t)

}
