package sort

import "testing"

// 堆排序是一种选择排序，完全二叉树
// 大顶堆：每个节点的值都大于或等于其子节点的值，在堆排序算法中用于升序排列	arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
// 小顶堆：每个节点的值都小于或等于其子节点的值，在堆排序算法中用于降序排列	arr[i] <= arr[2i+1] && arr[i] <= arr[2i+2]

// 创建一个堆 H[0……n-1]；
// 把堆首（最大值）和堆尾互换；
// 把堆的尺寸缩小 1，并调用 shift_down(0)，目的是把新的数组顶端数据调整到相应位置；
// 重复步骤 2，直到堆的尺寸为 1。

func HeapSort(nums []int) {
	n := len(nums)

	buildMaxHeap(nums, n)
	for i := n - 1; i >= 0; i-- {
		swap(nums, 0, i)
		n--
		heapify(nums, 0, n)
	}
}

func buildMaxHeap(nums []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}
}

func heapify(nums []int, i, n int) {
	lNode, rNode := i*2+1, i*2+2
	current := i

	if lNode < n && nums[lNode] > nums[current] {
		current = lNode
	}

	if rNode < n && nums[rNode] > nums[current] {
		current = rNode
	}

	if current != i {
		swap(nums, i, current)
		heapify(nums, current, n)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func TestHeapSort(t *testing.T) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	HeapSort(nums)
	t.Log(nums)
}
