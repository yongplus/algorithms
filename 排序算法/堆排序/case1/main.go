package main

import "log"

func buildHeap(arr []int, n int) {
	for k := n / 2; k >= 1; k-- {
		log.Println(k, arr)
		heapify(arr, n, k)
	}
}

func heapify(arr []int, n, k int) {
	for {
		maxPos := k
		if k*2 <= n && arr[k*2] > arr[k] {
			maxPos = k * 2
		}
		if k*2+1 <= n && arr[k*2+1] > arr[maxPos] {
			maxPos = k*2 + 1
		}
		if maxPos == k {
			break
		}
		arr[maxPos], arr[k] = arr[k], arr[maxPos]
		k = maxPos

	}
}

func sort(arr []int, n int) {
	buildHeap(arr, n) //构建堆
	for n > 1 {
		arr[n], arr[1] = arr[1], arr[n]
		n--
		heapify(arr, n, 1) //重新堆化
	}
}

/*
	   5
	 /   \
	2     3
   /  \   /
  1    2  1
 #求最后一个非叶子节点公式：节点数/2,如上述堆，第一个非叶子节点是3
*/

func main() {
	arr := []int{0, 7, 5, 19, 8, 4, 1, 20, 13, 16}
	//buildHeap(arr, len(arr)-1)
	sort(arr, len(arr)-1)
	log.Println(arr)
}
