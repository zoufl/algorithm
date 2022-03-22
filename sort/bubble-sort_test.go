package sort

import "testing"

// 比较相邻元素，如果前者比后者大，则进行位置交换
// 从第一个开始比较到未确定位置的最后一个，每经过一轮，则代表当前最大的数到达本次最后的位置

func BubbleSort(nums []int) {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	BubbleSort(nums)
	t.Log(nums)
}

func BenchmarkBubbleSort(b *testing.B) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	for i := 0; i < b.N; i++ {
		BubbleSort(nums)
	}
}
