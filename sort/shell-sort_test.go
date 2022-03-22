package sort

import "testing"

// 分组
// 把记录按下标的一定增量分组，对每组使用直接插入排序算法排序；
// 随着增量逐渐减少，每组包含的关键词越来越多，当增量减至1时，整个文件恰被分成一组，算法便终止。

func ShellSort(nums []int) {
	n := len(nums)

	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			current := nums[i]
			preIndex := i - gap

			for preIndex >= 0 && nums[preIndex] > current {
				nums[preIndex+gap] = nums[preIndex]
				preIndex -= gap
			}

			nums[preIndex+gap] = current
		}
	}
}

func TestShellSort(t *testing.T) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	ShellSort(nums)
	t.Log(nums)
}

func BenchmarkShellSort(b *testing.B) {
	nums := []int{5, 8, 6, 3, 9, 2, 1, 7}
	for i := 0; i < b.N; i++ {
		ShellSort(nums)
	}
}
