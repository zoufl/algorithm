package leetcode

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/3sum/

// 记录出现次数，去重，排序
func ThreeSum1(nums []int) [][]int {
	counter := make(map[int]int)
	ans := make([][]int, 0)

	for _, num := range nums {
		counter[num]++
	}

	uniqNums := make([]int, 0)
	for k := range counter {
		uniqNums = append(uniqNums, k)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		// x=0 3x=0
		if uniqNums[i]*3 == 0 && counter[uniqNums[i]] >= 3 {
			ans = append(ans, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}

		for j := i + 1; j < len(uniqNums); j++ {
			// x+x+y=0
			if uniqNums[i]*2+uniqNums[j] == 0 && counter[uniqNums[i]] > 1 {
				ans = append(ans, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}

			// x+y+y=0
			if uniqNums[j]*2+uniqNums[i] == 0 && counter[uniqNums[j]] > 1 {
				ans = append(ans, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}

			// x+y+z=0
			third := 0 - uniqNums[i] - uniqNums[j]
			if third > uniqNums[j] && counter[third] > 0 {
				ans = append(ans, []int{uniqNums[i], uniqNums[j], third})
			}
		}
	}

	return ans
}

// 排序+双指针
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := n - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum == 0 {
				ans = append(ans, []int{i, nums[l], nums[r]})
				for nums[l] == nums[l+1] {
					l++
				}
				for nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			}
		}
	}

	return ans
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	t.Log(ThreeSum(nums))

	nums = []int{}
	t.Log(ThreeSum(nums))

	nums = []int{0}
	t.Log(ThreeSum(nums))
}
