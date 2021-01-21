package coin

func minCoins(coins []int, amount int) int {
	f := func(n int) int {
		if n <= 0 {
			return 0
		}
		for _, v := range coins {
			minCoin := f(n - v)
			 
		}
	}
	return f(amount)
}
