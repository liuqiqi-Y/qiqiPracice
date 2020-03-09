// 股票最大收益
/*
	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
	一天只能进行一次交易（买入或卖出），设计一个算法计算你所能
	获取的最大利润。
	示例:
	输入: [7,1,5,3,6,4]
	输出: 7
*/
package main

import "fmt"

func maxProfit(prices []int) int {
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		val := prices[i+1] - prices[i]
		if val > 0 {
			sum += val
		}
	}
	return sum
}
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	val := maxProfit(prices)
	fmt.Printf("%d\n", val)
}
