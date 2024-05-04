func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    res := 0
    for i := 1; i < len(prices); i++ {
        if prices[i-1] < prices[i] {
            res += prices[i] - prices[i-1] 
        }
    }
    return res
}
