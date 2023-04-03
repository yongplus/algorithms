package main

/**
第一，冒泡排序是原地排序算法吗？
冒泡的过程只涉及相邻数据的交换操作，只需要常量级的临时空间，所以它的空间复杂度为 O(1)，是一个原地排序算法。

第二，冒泡排序是稳定的排序算法吗？
在冒泡排序中，只有交换才可以改变两个元素的前后顺序。为了保证冒泡排序算法的稳定性，当有相邻的两个元素大小相等的时候，我们不做交换，相同大小的数据在排序前后不会改变顺序，所以冒泡排序是稳定的排序算法。

第三，冒泡排序的时间复杂度是多少？
最好情况下，要排序的数据已经是有序的了，我们只需要进行一次冒泡操作，就可以结束了，所以最好情况时间复杂度是 O(n)。而最坏的情况是，要排序的数据刚好是倒序排列的，我们需要进行 n 次冒泡操作，所以最坏情况时间复杂度为 O(n2)。
*/
import "log"

func bubbleSort(nums []int) []int {
	for i := len(nums); i > 0; i-- {
		haveSwap := false //如果没有交换数据 则说明有序，这也是最好情况下时间复杂度O(n)的关键逻辑
		for j := 0; j < i-1; j++ {
			if nums[j] > nums[j+1] {
				haveSwap = true
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
		if !haveSwap {
			break
		}
	}
	return nums
}

func main() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	log.Println(bubbleSort(nums))
}
