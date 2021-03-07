package leetcode

func minOperations(boxes string) []int {
	n := len(boxes)
	result := make([]int, n)
	for i := 0; i < len(boxes); i++ {
		if i-1 >= 0 {
			left := boxes[:i]
			for j := 0; j < len(left); j++ {
				if left[j] == '1' {
					result[i] += i - j
				}
			}
		}
		if i+1 < len(boxes) {
			right := boxes[i+1:]
			for h := 0; h < len(right); h++ {
				if right[h] == '1' {
					result[i] += h + 1
				}
			}
		}

	}
	return result
}
