package main

import (
	"log"
	"math"
)

/**
希尔排序不稳定
空间复杂度O(1)，
时间复杂度，平均O(1^1..3)，最坏O(n^2)
*/

//The function below wrote by self according to WikiPedia's description about the principles of shell sorting

func ShellSort(nums []int) {
	n := len(nums)
	for stepSize := n / 2; stepSize > 0; stepSize = stepSize / 2 {
		cols := int(math.Ceil(float64(len(nums)) / float64(stepSize)))
		for c := 0; c <= cols; c++ {
			i := c + stepSize
			for ; i < len(nums); i += stepSize {
				value := nums[i]
				j := i
				for ; j > stepSize-1; j -= stepSize {
					if nums[j-stepSize] > value {
						nums[j] = nums[j-stepSize]
					} else {
						break
					}
				}
				nums[j] = value
			}
		}
		log.Println(stepSize)
	}

	log.Println(nums)
}

//The function below was copied from WikiPedia, link: https://zh.wikipedia.org/wiki/%E5%B8%8C%E5%B0%94%E6%8E%92%E5%BA%8F
func ShellSort2(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	stepSize := n / 2 //步长
	for stepSize > 0 {
		for i := stepSize; i < n; i++ {
			j := i
			for j >= stepSize && nums[j] < nums[j-stepSize] {
				nums[j], nums[j-stepSize] = nums[j-stepSize], nums[j]
				j = j - stepSize
			}
		}
		stepSize = stepSize / 2
	}
	log.Println(nums)
}

func main() {
	nums := []int{13, 14, 94, 33, 82, 25, 59, 94, 65, 23, 45, 27, 73, 25, 39, 10}
	ShellSort2(nums)
}
