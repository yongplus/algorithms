package main

import (
	"log"
	"math"
)

/**
* 第120题：给定一个三角形，找出自顶向下的最小路径和。
每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
*/

func dp(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	result := math.MaxInt
	for i := 1; i < len(triangle); i++ {
		for j := 0; j <= i; j++ {
			var x int
			if j == 0 { //如果是每行的第一个元素
				x = triangle[i-1][0] + triangle[i][j]
			} else if j == i { //如果是每行最后一个元素
				x = triangle[i-1][j-1] + triangle[i][j]
			} else { //中间元素
				x = min(triangle[i-1][j-1]+triangle[i][j], triangle[i-1][j]+triangle[i][j])
			}
			triangle[i][j] = x
			if i == len(triangle)-1 { //如果是最后一行
				result = min(result, x)
			}
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {

	triangle1 := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	triangle2 := [][]int{
		{-1},
		{-2, -3},
	}
	log.Println(dp(triangle1))
	log.Println(dp(triangle2))
}
