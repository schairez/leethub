//constraints
//2 <= cost.length <= 1000
//0 <= cost[i] <= 999

//if n == 1 || n == 0 cost to arrive at stair = 0
//recurrece relation
// dp[n] = min{dp[n-1]+cost[n-1], dp[n-2]+cost[n-2]}

func min(a, b int) int { if a <= b {return a}; return b}


//time: O(n)
//cost: O(n)
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    dp := make([]int, n+1)
    dp[0], dp[1] = 0, 0
    for stair := 2; stair <= n; stair++ {
        costFromPrevStair := dp[stair-1] + cost[stair-1]
        costFromTwoPrevStair := dp[stair-2] + cost[stair-2]
        dp[stair] = min(costFromPrevStair, costFromTwoPrevStair)
    }
    return dp[n]
}
























/*

//time: O(n)
//cost: O(n)
func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}
	dp := make([]int, len(cost)+2)
	for i := 0; i <= 2; i++ {
		dp[i] = 0
	}
	for i := 3; i < len(dp); i++ {
		dp[i] = min(dp[i-1]+cost[i-2], dp[i-2]+cost[i-3])
	}
	return dp[len(dp)-1]

    
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b 
}
*/