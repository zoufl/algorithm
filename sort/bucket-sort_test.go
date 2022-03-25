package sort

import (
	"sort"
	"testing"
)

// 设置固定数量的空桶。
// 求要排序数组的最大值和最小值，根据这两个值，再根据定几个桶，来定每个桶的范围。
// 对每个桶都进行排序。
// 按照顺序输出每个桶的数据

func BucketSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	max, min := nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		}
	}

	var bucketSize = 5
	var bucketNum = (max-min+1)/bucketSize + 1
	var buckets [][]int

	for i := 0; i < bucketSize; i++ {
		buckets = append(buckets, make([]int, 0))
	}

	for i := 0; i < len(nums); i++ {
		index := (nums[i] - min) / bucketNum
		buckets[index] = append(buckets[index], nums[i])
	}

	for _, bucket := range buckets {
		sort.Ints(bucket)
	}

	var index int
	for _, bucket := range buckets {
		for _, v := range bucket {
			nums[index] = v
			index++
		}
	}
}

func TestBucketSort(t *testing.T) {
	nums := []int{5, 16, 95, 1, 45, 66, 75, 20, 62, 38, 76, 94, 43}
	BucketSort(nums)
	t.Log(nums)
}
