package problem

/*
问题：
1143.最长公共子序列
https://leetcode-cn.com/problems/longest-common-subsequence/
*/

// dp[i][j] 代表l长度为i，r长度为j的时候最长的公共子序列个数
// 易得：
// 当l[i] == r[j]时，dp[i][j] = dp[i-1][j-1]+1
// 当l[i] != r[j]时，dp[i][j] = max(dp[i][j-1], dp[i][j-1])
func longestCommonSubsequence1(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

// 优化版：o（n）空间，n^2时间
func longestCommonSubsequence2(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)
	for i := 1; i <= len(text1); i++ {
		last := 0
		for j := 1; j <= len(text2); j++ {
			temp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = last + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			last = temp
		}
	}
	return dp[len(text2)]
}
