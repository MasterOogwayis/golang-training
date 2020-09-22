package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	go t2(arr, c)
	go p(c)

	time.Sleep(2 * time.Second)

	select {
	case i := <-c:
		fmt.Println(i)
	case <-time.After(5 * time.Second):
		fmt.Println("close")
		close(c)
	}

}

func t2(arr []int, c chan int) {
	for i := range arr {
		c <- i
		time.Sleep(time.Second)
	}
	close(c)
}

func p(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func testSum() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	c := make(chan int, 2)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func sum(a []int, c chan int) {
	total := 0
	for _, value := range a {
		total += value
	}
	c <- total
}

func t1() {
	c := make(chan int, 2) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
