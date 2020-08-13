package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(Sqrt(-100))

}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return 1, nil
}
