package main

import "math"

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	n := len(floor)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, numCarpets+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	if numCarpets*carpetLen >= n {
		return 0
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= numCarpets; j++ {
			// 情况 1：不使用地毯覆盖第 i 块砖块
			dp[i][j] = dp[i-1][j] + int(floor[i-1]-'0')

			// 情况 2：使用一条地毯覆盖从 i - carpetLen 到 i - 1 的砖块
			if j > 0 && i >= carpetLen {
				dp[i][j] = min(dp[i][j], dp[i-carpetLen][j-1])
			}
		}
	}

	return dp[n][numCarpets]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// floor := "10110101"
	// numCarpets := 2
	// carpetLen := 2

	floor := "11111"
	numCarpets := 2
	carpetLen := 3
	result := minimumWhiteTiles(floor, numCarpets, carpetLen)
	print("未被覆盖的白色砖块的最小数量为：\n", result)
}
