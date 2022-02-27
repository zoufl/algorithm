package leetcode

import (
	"testing"
)

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

func MaxProfitII(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

func TestMaxProfitII(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	t.Log(MaxProfitII(nums))

	nums = []int{1, 2, 3, 4, 5}
	t.Log(MaxProfitII(nums))

	nums = []int{7, 6, 4, 3, 1}
	t.Log(MaxProfitII(nums))
}
