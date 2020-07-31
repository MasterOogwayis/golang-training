package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 0
	arr[1] = 1
	arr[4] = 4

	fmt.Println(len(arr))

	x, y, z := multiple()
	fmt.Println(x, y, z)

}

func multiple() (int, string, bool) {
	a, b, c := 1, "String", true
	return a, b, c
}
