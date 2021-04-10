package main

import (
	"fmt"
)

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func main() {

	fmt.Println(isValid("{}[]"))
}

func isValid(s string) bool {
	// write code here
	if len(s) < 2 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '}' {
			stack = append(stack, s[i])
		} else {
			if len(stack) < 1 {
				return false
			}
			switch s[i] {
			case ')':
				if stack[len(stack)-1] != ')' {
					return false
				}
			case ']':
				if stack[len(stack)-1] != ']' {
					return false
				}
			case '}':
				if stack[len(stack)-1] != '}' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		}
	}
	return true
}
