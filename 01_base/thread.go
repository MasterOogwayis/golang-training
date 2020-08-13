package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 10)
	go Fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}

	
}

func testChannel()  {
	//go say("World")

	var ch = make(chan int, 1)

	ch <- 1
	go getCh(ch)
	ch <- 2
	go getCh(ch)

	fmt.Println("a")
}

func getCh(ch chan int)  {
	time.Sleep(2 * time.Second)
	fmt.Println(<- ch)
}


func testAsync()  {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan int, 1)

	go sum(s[:4], ch)
	go sum(s[4:], ch)

	x := <-ch
	time.Sleep(5 * time.Second)
	y := <-ch
	fmt.Println(x, y, x+y)
}

func sum(s []int, ch chan int) {
	sum := 0
	for _, num := range s {
		sum += num
	}
	ch <- sum
	fmt.Println(s)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}


func Fibonacci(n int, ch chan int) {
	x,y := 0,1
	for i := 0; i < n; i++ {
		ch <- x
		x,y=y,x+y
	}
	close(ch)


}
