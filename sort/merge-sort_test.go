package sort

import "testing"

// 分治法
// 1. 申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
// 2. 设定两个指针，最初位置分别为两个已经排序序列的起始位置；
// 3. 比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
// 4. 重复步骤 3 直到某一指针达到序列尾；
// 5. 将另一个序列剩下的所有元素直接复制到合并序列尾。

func MergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	middle := n / 2
	nums1, nums2 := nums[:middle], nums[middle:]

	return merge(MergeSort(nums1), MergeSort(nums2))
}

func merge(nums1, nums2 []int) []int {
	nums := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
	}

	if i < len(nums1) {
		nums = append(nums, nums1[i:]...)
		i++
	}

	if j < len(nums2) {
		nums = append(nums, nums2[j:]...)
		j++
	}

	return nums
}

func TestMergeSort(t *testing.T) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	t.Log(MergeSort(nums))
}
