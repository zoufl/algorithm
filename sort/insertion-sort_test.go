package sort

import "testing"

// 维护一个有序区，把元素一个一个插入到有序区的适当位置，直到所有元素有序为止
// 对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		current := array[i]
		preIndex := i - 1

		for preIndex >= 0 && array[preIndex] > current {
			array[preIndex+1] = array[preIndex]
			preIndex--
		}

		array[preIndex+1] = current
	}
}

func insertionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}

func TestInsertionSort(t *testing.T) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	InsertionSort(nums)
	t.Log(nums)
}

func BenchmarkInsertionSort(b *testing.B) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	for i := 0; i < b.N; i++ {
		InsertionSort(nums)
	}
}

func BenchmarkInsertionSort1(b *testing.B) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	for i := 0; i < b.N; i++ {
		insertionSort(nums)
	}
}
