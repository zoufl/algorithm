package sort

import "testing"

// 每一次选出最小者直接交换到左侧，省出了多余的元素交换

func SelectionSort(nums []int) {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func TestSelectionSort(t *testing.T) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	SelectionSort(nums)
	t.Log(nums)
}

func BenchmarkSelectionSort(b *testing.B) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	for i := 0; i < b.N; i++ {
		SelectionSort(nums)
	}
}
