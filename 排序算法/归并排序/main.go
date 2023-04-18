package main

import "log"

/**
第一，归并排序是稳定的排序算法吗？
结合我前面画的那张图和归并排序的伪代码，你应该能发现，归并排序稳不稳定关键要看 merge() 函数，也就是两个有序子数组合并成一个有序数组的那部分代码。在合并的过程中，如果 A[p...q]和 A[q+1...r]之间有值相同的元素，那我们可以像伪代码中那样，先把 A[p...q]中的元素放入 tmp 数组。这样就保证了值相同的元素，在合并前后的先后顺序不变。所以，归并排序是一个稳定的排序算法。

第二，归并排序的时间复杂度是多少？
不管是最好情况、最坏情况，还是平均情况，时间复杂度都是 O(nlogn)

空间复杂度是 O(n)。
*/
func mergeSort(nums []int, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	log.Println(p, q+1, r)
	mergeSort(nums, p, q)
	mergeSort(nums, q+1, r)
	merge(nums, p, q+1, r)
}

func merge(nums []int, p, q, r int) {
	tmp := make([]int, r-p+1)
	i := 0
	j := 0
	y := 0
	for (p+i) < q && (j+q) <= r {
		if nums[p+i] <= nums[j+q] {
			tmp[y] = nums[p+i]
			i++
		} else {
			tmp[y] = nums[j+q]
			j++
		}
		y++
	}

	var start, end int
	if (p + i) == q {
		start = j + q
		end = r
	} else {
		start = p + i
		end = q - 1
	}

	for ; start <= end; start++ {
		tmp[y] = nums[start]
		y++
	}

	for x := 0; x <= r-p; x++ {
		nums[p+x] = tmp[x]
	}

}

func main() {
	nums := []int{6, 5, 4, 1, 2, 3}
	mergeSort(nums, 0, len(nums)-1)
	log.Println(nums)
}
