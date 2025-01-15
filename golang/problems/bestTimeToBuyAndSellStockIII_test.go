package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][]int, 3)
	for k := 0; k <= 2; k++ {
		dp[k] = make([]int, len(prices))
	}

	for k := 1; k <= 2; k++ {
		minVal := prices[0]
		for i := 1; i < len(prices); i++ {
			minVal = min(minVal, prices[i]-dp[k-1][i-1])
			dp[k][i] = max(dp[k][i-1], prices[i]-minVal)
		}
	}

	return dp[2][len(prices)-1]
}

// func maxProfit(prices []int) int {
// 	buy1 := math.MaxInt
// 	sell1 := 0
// 	buy2 := math.MaxInt
// 	sell2 := 0
// 	for i := 0; i < len(prices); i++ {
// 		price := prices[i]
// 		buy1 = min(buy1, price)
// 		diff1 := prices[i] - buy1
// 		sell1 = max(sell1, diff1)
// 		price2 := prices[i] - sell1
// 		buy2 = min(buy2, price2)
// 		diff2 := prices[i] - buy2
// 		sell2 = max(sell2, diff2)
// 	}
//
// 	return sell2
// }

func TestMaxProfit(t *testing.T) {
	testCases := []maxProfitTestCase{
		{
			prices: []int{1, 6, 5, 8, 0, 9},
			expect: 16,
		},
		{
			prices: []int{3, 3, 5, 0, 0, 1, 4},
			expect: 6,
		},
		// {
		// 	prices: []int{1, 2, 3, 4, 5},
		// 	expect: 4,
		// },
		// {
		// 	prices: []int{7, 6, 4, 3, 1},
		// 	expect: 0,
		// },
		// {
		// 	prices: []int{7, 6, 4, 3, 4},
		// 	expect: 0,
		// },
	}

	for index, testCase := range testCases {
		result := maxProfit(testCase.prices)
		assert.Equal(t, testCase.expect, result, fmt.Sprintf("第%d个案例结果不一致", index+1))
	}
}

type maxProfitTestCase struct {
	prices []int
	expect int
}
