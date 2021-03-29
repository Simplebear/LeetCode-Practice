package leetcode

func findPeakElement(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	i := 0
	for i < l {
		if i == 0 && nums[i] > nums[i+1] {
			return i
		}
		if i > 0 && i < l-1 && nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
		if i == l-1 && nums[i] > nums[i-1] {
			return i
		}
		i++
	}
	return 0
}
