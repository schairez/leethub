// 416. Partition Equal Subset Sum
// rolling array DP table
// time: O(n*target)
// space: O(target)
func canPartition(nums []int) bool {
    n := len(nums)
    sumTotal := 0
    for i := range nums {
        sumTotal += nums[i]
    }
    if sumTotal % 2 == 1 {
        return false
    }
    target := sumTotal >> 1
    dp := make([]bool, target+1)
    dp[0] = true
    for i := 0; i < n; i++ {
        for j := target; j >= nums[i]; j-- {
            if dp[target] {
                return true
            }
            dp[j] = dp[j] || dp[j - nums[i]]
        }
    }
    return dp[target]
}

// top down dfs approach
// time: O(n*target)
// space: O(n*target)
func canPartitionMemo(nums []int) bool {
    n := len(nums)
    sumTotal := 0
    for i := range nums {
        sumTotal += nums[i]
    }
    if sumTotal % 2 == 1 {
        return false
    }
    type pair struct {idx, total int}
    memo := make(map[pair]bool)
    var (
        dfs func(int, int) bool
    )
    // target is half
    target := sumTotal >> 1
    dfs = func(idx, total int) bool {
        if total == 0 {
            return true
        }
        if total < 0 || idx >= n {
            return false
        }
        if isPossible, exists := memo[pair{idx, total}]; exists { 
            return isPossible
        }
        includeItem := false
        if total - nums[idx] >= 0 {
            includeItem = dfs(idx+1, total - nums[idx])
        }
        // include item or exclude item
        memo[pair{idx, total}] = includeItem || dfs(idx+1, total)
        
        
        return memo[pair{idx, total}]
    }
    
    return dfs(0, target)
}