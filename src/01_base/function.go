package main

import "fmt"

func init() {
	fmt.Println("init ...")
}

func main() {

	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven) // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)

}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func isOdd(integer int) bool {
	return integer%2 != 0
}

func isEven(integer int) bool {
	return integer%2 == 0
}

type testInt func(int) bool

func testMultiple() {
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
