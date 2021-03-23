package leetcode

import "strconv"

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	n := 0
	if la > lb {
		n = la
	} else {
		n = lb
	}
	result := ""
	tn := 0
	for i := 0; i < n; i++ {
		if i < la {
			tn += int(a[la-1-i])
		}
		if i < lb {
			tn += int(b[lb-1-i] - '0')
		}
		result = strconv.Itoa(tn%2) + result
		tn = tn / 2
	}
	if tn == 1 {
		result = "1" + result
	}
	return result
}
