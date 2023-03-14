package main

import (
	"log"
)

/*
第300题：给定一个无序的整数数组，找到其中最长上升子序列的长度。
示例:
输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:
可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
*/
func FindLongestIncreasingSubsequence(nums []int) int {
	dp := make([]int, len(nums))

	result := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	log.Println(dp)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	//nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	nums := []int{1, 9, 5, 9, 3}
	log.Println(FindLongestIncreasingSubsequence(nums))
}
