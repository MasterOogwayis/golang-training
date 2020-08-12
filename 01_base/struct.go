package main

import "fmt"

func main() {
	var structPointer *Books

	var book Books = Books{"title", "author", "subject", 1}
	fmt.Println(book)
	
	structPointer = &book
	setTitle(structPointer)
	fmt.Println(structPointer)
	//t1()
	//t2()
}
func setTitle(book *Books)  {
	book.title = "new title"
}


func t2() {
	var book1 Books
	var book2 Books

	book1.title = "java"
	book1.author = "gs"
	book1.subject = "DDD"

	book2.title = "go"
	book2.author = "c"
	book2.subject = "lang"
	book2.bookId = 2

	fmt.Println(book1, book2)

	fmt.Println(getAttr(book1))

}

func t1() {
	var book Books = Books{"java", "gs", "面向对象编程", 8080}
	fmt.Println(book.title)
	fmt.Println(book.subject)
	fmt.Println(book.author)
	fmt.Println(book.bookId)
}

func getAttr(book Books) string {
	return book.author
}

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}
