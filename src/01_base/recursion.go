package main

import "fmt"

/**
 * 递归
 * recursion
 */
func main() {
	//var i uint64 = 15
	//fmt.Printf("%d 的阶乘是 %v", i, Factorial(i))

	t()
	kda(10)

}

func kda(size int) {
	var arr []int = make([]int, size, size)

	for i := 0; i < size; i++ {
		arr[i] = i
	}
	fmt.Println(arr)

}

func t() {
	size := 15
	arr := [15]int{}
	for i := 0; i < size; i++ {
		arr[i] = fibonacci(i)
	}
	fmt.Println(arr)

	var slice = arr[5:10]
	slice[0] = 999
	fmt.Println(arr)
}

/**
 * 阶乘
 */
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * (n - 1)
		return result
	}
	return 1

}

/**
 * 斐波那契数列
 */
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
