package main

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func main() {
	fmt.Println(test())
}
func test() int {
	result := 5
	defer func() {
		fmt.Println("defer 1")
	}()

	defer func() {
		fmt.Println("defer 2")
	}()
	return result
}
