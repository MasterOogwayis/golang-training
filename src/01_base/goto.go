package main

import "fmt"

func main() {

	testFor()

	for i := 0; i < 10; i++ {

	}

}

func testFor() {
	sum := 1

	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)
}

func testGoto() {
	i := 0

Here:
	fmt.Println(i)
	i++
	if i < 5 {
		goto Here
	}
}
