package leetcode

import (
	"math"
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

func MaxProfit(prices []int) int {
	min, max := math.MaxInt64, 0

	for _, price := range prices {
		if price < min {
			min = price
		} else if price-min > max {
			max = price - min
		}
	}

	return max
}

// 保存所有交易结果
func MaxProfit1(prices []int) int {
	min, max := prices[0], prices[0]
	l := make([]int, 0)
	l = append(l, max-min)

	for _, price := range prices {
		if price < min {
			min = price
			max = price

			continue
		}

		if price > max {
			max = price

			l = append(l, max-min)
		}
	}

	sort.Ints(l)

	return l[len(l)-1]
}

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	t.Log(MaxProfit(nums))

	nums = []int{7, 6, 4, 3, 1}
	t.Log(MaxProfit(nums))

	nums = []int{2, 4, 1}
	t.Log(MaxProfit(nums))
}
