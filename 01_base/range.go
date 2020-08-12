package main

import "fmt"

/**
 * 语言范围(Range)
 */
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	for index, num := range nums {
		sum += num
		fmt.Println(index)
	}
	fmt.Println(sum)

	kvs := map[string]string{"name": "zsw", "age": "18"}

	for key, value := range kvs {
		fmt.Println(key, value)
	}
	// c - Unicode的值
	for i, c := range "golang" {
		fmt.Println(i, c)
	}

}
