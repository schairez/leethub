/*
time: O(n)
space: O(1)

*/

func maxProfit(prices []int) int {
    maxProfit := 0 
    buy := prices[0]
    for i := 1; i < len(prices); i++ {
        curr := prices[i]
        maxProfit = max(maxProfit, curr - buy)
        buy = min(curr, buy)
        
    }
    return maxProfit
}

func min(a, b int) int {
    if a < b {
        return a 
    }
    return b 
}

func max(a, b int) int {
    if a > b {
        return a 
    }
    return b 
}
