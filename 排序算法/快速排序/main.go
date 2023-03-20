package main

import (
	"log"
)

/**
 * 快排是一种原地、不稳定的排序算法,快排的时间复杂度是 O(nlogn),最差O(n^2),空间复杂度O(1)
 *
 */
func quickSort(nums []int, p, r int) {
	if p >= r {
		return
	}

	partition(nums, p, r)

	q := (p + r) / 2
	quickSort(nums, p, q)
	quickSort(nums, q+1, r)

}

func partition(nums []int, p, r int) {
	pivot := nums[r]
	i := p
	for j := p; j < r; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
}

func main() {
	nums := []int{6, 5, 4, 1, 2, 3}
	//nums = []int{1, 2, 3}
	quickSort(nums, 0, len(nums)-1)
	log.Println(nums)
}
