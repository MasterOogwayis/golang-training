package main

import "fmt"

func main() {
	x,y := 0,1
	x,y=y,x+y
	fmt.Println(x, y)
}
