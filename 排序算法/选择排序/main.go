package main

import "log"

func selectionSort(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		minidx := i
		for j := len(nums) - 1; j > i; j-- {
			if nums[minidx] > nums[j] {
				minidx = j
			}
		}
		if minidx != i {
			nums[i], nums[minidx] = nums[minidx], nums[i]
		}
	}

	return nums
}

func main() {
	nums := []int{4, 5, 6, 1, 2, 3}
	log.Println(selectionSort(nums))
}
