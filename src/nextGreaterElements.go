package leetcode

import (
	"fmt"
)

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}
	s := []int{}
	for i := 0; i < 2*n-1; i++ {
		fmt.Println(i)
		for len(s) > 0 && nums[i%n] > nums[s[len(s)-1]] {
			result[s[len(s)-1]] = nums[i%n]
			s = s[:len(s)-1]

		}
		s = append(s, i%n)
	}
	return result
}
