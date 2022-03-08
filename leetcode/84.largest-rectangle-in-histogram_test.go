package leetcode

import (
	"math"
	"testing"
)

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/

// 执行会超时
func LargestRectangleArea1(heights []int) int {
	maxArea := 0

	for i := range heights {
		minHeight := math.MaxInt64

		for j := i; j < len(heights); j++ {
			minHeight = int(math.Min(float64(minHeight), float64(heights[j])))

			maxArea = int(math.Max(float64(maxArea), float64(minHeight*(j-i+1))))
		}
	}

	return maxArea
}

// 单栈
func LargestRectangleArea(heights []int) int {
	stack := []int{-1}
	ans := 0

	for i, height := range heights {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > height {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans = int(math.Max(float64(ans), float64(heights[n]*(i-stack[len(stack)-1]-1))))
		}

		stack = append(stack, i)
	}

	if len(stack) > 1 {
		i := stack[len(stack)-1]
		for len(stack) > 1 {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans = int(math.Max(float64(ans), float64(heights[n]*(i-stack[len(stack)-1]))))
		}
	}

	return ans
}

func TestLargestRectangleArea(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	t.Log(LargestRectangleArea(heights))

	heights = []int{2, 4}
	t.Log(LargestRectangleArea(heights))
}
