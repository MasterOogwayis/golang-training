package main

import "fmt"

func main() {

	testInterface()

}

func testInterface() {

	var kvs map[string]interface{} = make(map[string]interface{})

	kvs["name"] = "zsw"
	kvs["age"] = 18
	kvs["attr"] = MapMembers{"title"}

	fmt.Println(kvs)
	fmt.Println(&kvs)

	var m = MapMembers{title: "title"}
	m.title = "1"
	fmt.Println(m)
	var mp *MapMembers = &m
	mp.title = "2"
	fmt.Println(m)

}

func testMap() {
	var kvs map[string]string

	kvs = make(map[string]string)

	kvs["msg"] = "This is a test demo."
	kvs["notice"] = "nil + n"

	for s2 := range kvs {
		fmt.Println(kvs[s2])
	}

	delete(kvs, "notice")

	fmt.Println(kvs)
}

type MapMembers struct {
	title string
}
