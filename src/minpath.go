package leetcode

func minpath(grid [][]int) int {
	rows, clos := len(grid), len(grid[0])
	if rows == 0 || clos == 0 {
		return 0
	}
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[0] = make([]int, clos)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j < clos; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < clos; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][clos-1]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
