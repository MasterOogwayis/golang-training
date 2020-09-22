package main

import "fmt"

func main() {

	var slice = make([]int, 5, 10)
	fmt.Println(slice)
	//slice[5] = 1
	b := append(slice, 12)
	fmt.Println(b)

	c := make([]int, len(b), 2*cap(b))

	copy(c, b)
	fmt.Println(c)

	//t3()
	//t4()
	//t5()
}

func t5() {
	var arr = [10]int{1, 2, 3, 4, 5}
	var slice = arr[5:]
	slice[0] = 11
	fmt.Println(arr)

}

func t4() {
	var arr = make([]int, 5, 10)
	fmt.Printf("len = %v, cap = %d, slice = %x", len(arr), cap(arr), arr)
}

func t3() {
	var arr = [7]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(arr[1:])
	fmt.Printf("%T", make([]int, 5))
}

type Object struct {
}
