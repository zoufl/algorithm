package sort

import "testing"

// 分治
// 元素的比较和交换是从两端向中间进行的，较大的元素一轮就能交换到后面的位置，而较小的元素一轮就能交换到前面的位置
// 1. 在待排序的元素任取一个元素作为基准（通常选第一个元素，但最好的方法是从待排序元素中随机选取一个为基准），称为基准元素（pivot）
// 2. 将待排序的元素进行分区，比基准元素大的元素放在它的右边，比基准元素小的放在它的左边
// 3. 对左右两个分区重复以上步骤，直到所有的元素都是有序的

func QuickSort(nums []int, start, end int) {
	if start > end {
		return
	}

	pivotIndex := partition(nums, start, end)
	QuickSort(nums, start, pivotIndex-1)
	QuickSort(nums, pivotIndex+1, end)
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	left, right := start, end

	for left < right {
		for left < right && nums[right] > pivot {
			right--
		}

		for left < right && nums[left] <= pivot {
			left++
		}

		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	nums[left], nums[start] = nums[start], nums[left]

	return left
}

func TestQuickSort(t *testing.T) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	QuickSort(nums, 0, len(nums)-1)
	t.Log(nums)
}

func BenchmarkQuickSort(b *testing.B) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	for i := 0; i < b.N; i++ {
		QuickSort(nums, 0, len(nums)-1)
	}
}
