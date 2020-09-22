package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Int()

	fmt.Println(a)

	if a > 5 {
		fmt.Println(a > 5)
	}

}
