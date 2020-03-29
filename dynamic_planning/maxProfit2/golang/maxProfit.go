package dp

/*
给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每次交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

示例 1:

输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
输出: 8
解释: 能够达到的最大利润:
在此处买入 prices[0] = 1
在此处卖出 prices[3] = 8
在此处买入 prices[4] = 4
在此处卖出 prices[5] = 9
总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.

注意:

    0 < prices.length <= 50000.
    0 < prices[i] < 50000.
    0 <= fee < 50000.
*/
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][2]int, n) //dp[i][0]:该天卖出的最大利润，dp[i][1]:该天买入的最大利润
	dp[0][0] = 0            //第一天不可能卖出，利润为0
	dp[0][1] = -prices[0] - fee
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])     //1.前一天已卖出，今天什么都不做 2.前一天买入了，今天卖出
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee) //买入时扣除fee
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
