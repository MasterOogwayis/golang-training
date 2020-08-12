package main

import "fmt"

func main() {
	var kvs map[string]string

	kvs = make(map[string]string)

	kvs["msg"] = "This is a test demo."

	fmt.Println(kvs)

}
