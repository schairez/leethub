


func canPartition(nums []int) bool {
    n := len(nums)
    totalSum := 0
    for i := range nums {
        totalSum += nums[i]
    }
    if totalSum % 2 != 0 {
        return false
    }
    partTarget := totalSum / 2
    
    dp := make([]bool, partTarget + 1)
    dp[0] = true
    
    // 11 -  _ = -values is err
    for i := 0; i < n; i++ {
        for j := partTarget; j >= nums[i]; j-- {
            //t or f
            dp[j] = dp[j] || dp[j-nums[i]]
        }
    }
    return dp[partTarget]
}