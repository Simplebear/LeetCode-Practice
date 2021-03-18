package leetcode

func selectSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		max := 0
		for j := 1; j < length-i; j++ {
			if nums[j] > nums[max] {
				max = nums[j]
			}
		}
		nums[length-1-i], nums[max] = nums[max], nums[length-1-i]
	}
}
