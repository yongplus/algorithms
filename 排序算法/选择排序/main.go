package main

import "log"

func selectionSort(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		//minidx := i
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

func main() {
	nums := []int{4, 5, 6, 3, 2, 1}
	log.Println(selectionSort(nums))
}
