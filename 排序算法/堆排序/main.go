package main

import (
	"log"
)

/**
O(nlogn)，所以，堆排序整体的时间复杂度是 O(nlogn)
堆排序不是稳定的排序算法
*/
func buildHeap(a []int, n int) {
	for i := n / 2; i >= 1; i-- {
		heapify(a, n, i)
	}
}
func heapify(a []int, n, i int) {
	for {
		maxPos := i
		if i*2 <= n && a[i] < a[i*2] {
			maxPos = i * 2
		}

		if i*2+1 <= n && a[maxPos] < a[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		a[i], a[maxPos] = a[maxPos], a[i]
		i = maxPos
	}
}

func sort(a []int, n int) {
	buildHeap(a, n)
	k := n
	for k > 1 {
		a[1], a[k] = a[k], a[1]
		k--
		heapify(a, k, 1)
	}
}

func main() {
	arr := []int{0, 7, 5, 19, 8, 4, 1, 20, 13, 16}
	//buildHeap(arr, len(arr)-1)
	sort(arr, len(arr)-1)
	log.Println(arr)
}

// 输出结果为 [1 2 3 5 6 8 10]

//在上述代码中，heapSort 函数实现了堆排序的过程。首先，我们调用 heapify 函数对整个数组构建大根堆。然后，我们从后往前遍历数组，每次将 arr[0] 与当前未排序的最后一个元素交换位置，然后对 arr[:i] 执行 heapify 操作，使得剩余数组仍然是一个大根堆。最终得到的数组即为排序后的结果。在 heapify 函数中，我们从 n/2 - 1 开始，对当前的节点 i 执行 sift down 操作，将 i 节点及其子树构建成一个大根堆。在 siftDown 函数中，我们将 i 节点与其左右子节点进行比较，找到三个元素中最大的那个，并将其下标存储在变量 largest 中。如果最大值下标不是 i，就将 arr[i] 和 arr[largest] 交换位置，并继续执行 sift down 操作，将 largest 节点及其子树构建成一个大根堆。
