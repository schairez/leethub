func canPartition(nums []int) bool {
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