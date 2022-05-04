func canPartition(nums []int) bool {
    n := len(nums)
    // sum values
    sum := 0
    for i := range nums {
        sum += nums[i]
    }
    // even or odd
    if sum % 2 == 1 {
        return false
    }
    // time: O(n*target)
    target := sum >> 1
    dp := make([]bool, target+1)
    dp[0] = true
    for i := 0; i < n; i++ {
        for j := target; j >= nums[i]; j-- {
            if dp[target] { 
                return true
            }
            //      excl     incl
            dp[j] = dp[j] || dp[j - nums[i]]
        }
    }
    return dp[target]
    
}

// top down memoized
func canPartitionMemo(nums []int) bool {
    n := len(nums)
    // sum values
    sum := 0
    for i := range nums {
        sum += nums[i]
    }
    // even or odd
    if sum % 2 == 1 {
        return false
    }
    target := sum >> 1
    type pair struct {idx, target int}
    memo := make(map[pair]bool)
    
    var dfs func(int, int) bool
    dfs = func(idx, target int) bool {
        if target == 0 {
            return true
        }
        if target < 0 || idx >= n {
            return false
        }
        if isPossible, exists := memo[pair{idx, target}]; exists {
            return isPossible
        }
        includeItem := false
        if target - nums[idx] >= 0 {
            includeItem = dfs(idx+1, target - nums[idx])
        }
        memo[pair{idx, target}] = includeItem || dfs(idx+1, target)
        
        return memo[pair{idx, target}]
    }
    return dfs(0, target)
    
}