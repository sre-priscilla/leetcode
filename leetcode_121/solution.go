package main;

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	} else {
		min, profit := prices[0], 0
		for i := 1; i < n; i += 1 {
			if prices[i] < min {
				min = prices[i]
			} else {
				delta := prices[i] - min
				if delta > profit {
					profit = delta
				}
			}
		}
		return profit
	}
}
