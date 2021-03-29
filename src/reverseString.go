package leetcode

import "fmt"

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		temp := s[l-1-i]
		s[l-1-i] = s[i]
		s[i] = temp
	}
	fmt.Println(s)
}
