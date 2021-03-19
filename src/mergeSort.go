package leetcode

func mergeSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	mid := length / 2
	left := mergeSort(nums[:mid])
	rigth := mergeSort(nums[mid:])
	return merge(left, rigth)
}

func merge(l []int, r []int) []int {
	temp := make([]int, 0)
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			temp = append(temp, l[i])
			i++
		} else {
			temp = append(temp, r[j])
			j++
		}
		if i < len(l) {
			temp = append(temp, l[i:]...)
		}
		if j < len(r) {
			temp = append(temp, r[j:]...)
		}
	}
	return l
}
