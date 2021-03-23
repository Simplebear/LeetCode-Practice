package leetcode

func maxSubArray(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	result := nums[0]
	for i := 0; i < length; i++ {
		dp[i] = nums[i]
	}
	for j := 1; j < length; j++ {
		if dp[j-1] > 0 {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	for p := range dp {
		if p > result {
			result = p
		}
	}
	return result
}
