package main

import "fmt"

/**
类型转换
*/
func main() {
	expression()
}

func expression() {
	var a int = 10
	var b int = 2
	var sum float32

	sum = float32(a + b)

	fmt.Printf("sum = %v", sum)

}
