package main

import "fmt"

func main() {
	var arr [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var slice []int = arr[:]

	s := make([]int, 5, 10)

	count := copy(s, slice)

	fmt.Printf("%d %d", s, count)


}

func t0() {
	arr := [5]int{1, 2, 7, 4, 5}

	for index, i := range arr {
		fmt.Println(index, i)
	}

	balance := [...]float32{1.1, 1.2, 1.9, 1.4, 1.5}

	for index, b := range balance {
		fmt.Println(index, b)
	}

	fmt.Errorf(string(len(balance)))

	cr := [7]int{1, 2, 3, 4, 5, 6, 7}

	length(cr[0:])
}

func length(arr []int) int {
	fmt.Println(arr)
	return len(arr)
}
