package main

func maxProfit(prices []int) int {
	max := 0
	buy := prices[0]
	l := len(prices)
	for i := 0; i < l; i++ {
		if prices[i] < buy {
			buy = prices[i]
		}
		if prices[i]-buy > max {
			max = prices[i] - buy
		}
	}
	return max
}
