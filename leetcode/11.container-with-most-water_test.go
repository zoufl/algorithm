package leetcode

import (
	"math"
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/container-with-most-water/

// 双指针
func MaxArea1(height []int) int {
	max := 0

	for l, r := 0, len(height)-1; l < r; {
		area := int(math.Min(float64(height[l]), float64(height[r]))) * (r - l)

		if area > max {
			max = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return max
}

// 双指针
func MaxArea(height []int) int {
	max := 0

	for l, r := 0, len(height)-1; l < r; {
		width, high := r-l, 0
		if height[l] < height[r] {
			high = height[l]
			l++
		} else {
			high = height[r]
			r--
		}

		if temp := width * high; temp > max {
			max = temp
		}
	}

	return max
}

func TestMaxArea(t *testing.T) {
	//nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//t.Log(MaxArea(nums))
	//
	//nums = []int{1, 1}
	//t.Log(MaxArea(nums))

	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}

		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, n-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})

				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			}
		}
	}

	return ans
}
