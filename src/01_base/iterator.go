package main

import "fmt"

func main() {

	a := 1
	for a < 100 {
		a++
		fmt.Println("循环 ", a)
		if a == 50 {
			break
		}
	}

}
