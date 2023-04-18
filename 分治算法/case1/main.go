package main

import "log"

/*
```
假设我们有 n 个数据，我们期望数据从小到大排列，那完全有序的数据的有序度就是 n(n-1)/2，逆序度等于 0；相反，倒序排列的数据的有序度就是 0，逆序度是 n(n-1)/2。除了这两种极端情况外，我们通过计算有序对或者逆序对的个数，来表示数据的有序度或逆序度。

用分治法求有序度
```
*/
var num = 0

func mergeCount(nums []int, p, r int) {
	if p >= r {
		return
	}
	q := (r + p) / 2
	mergeCount(nums, p, q)
	mergeCount(nums, q+1, r)
	merge(nums, p, q, r)
	return
}

func merge(nums []int, p, q, r int) {

	i := p
	j := q + 1
	k := 0
	tmp := make([]int, r-p+1)
	for i <= q && j <= r {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {

			num += (q - i + 1)
			tmp[k] = nums[j]
			j++
		}
		k++
	}

	for ; i <= q; i++ {
		tmp[k] = nums[i]
		k++
	}

	for ; j <= r; j++ {
		tmp[k] = nums[j]
		k++
	}

	for i = 0; i <= r-p; i++ {
		nums[p+i] = tmp[i]
	}
}

func main() {

	//nums := rand.Perm(1000)
	nums := []int{2, 4, 3, 1, 5, 6}
	mergeCount(nums, 0, len(nums)-1)
	log.Println(nums)
	log.Println(num)
}
