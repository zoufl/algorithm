package sort

import "testing"

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
