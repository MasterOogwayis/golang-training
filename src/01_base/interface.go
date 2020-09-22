package main

import "fmt"

func main() {
	var phone Phone = new(IPhone)
	phone.call()

	var phoneNokia Phone = new(Nokia)
	phoneNokia.call()

}


type Phone interface {
	call()
}

type IPhone struct {
	
}

type Nokia struct {

}

func (iPhone IPhone) call()  {
	fmt.Println("iPhone ...")
}

func (nokia Nokia) call()  {
	fmt.Println("Nokia ...")
}




