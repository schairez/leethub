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
        incl := dfs(idx+1, total - nums[idx])
        excl := dfs(idx+1, total)
        
        memo[pair{idx, total}] = incl || excl
        
        return memo[pair{idx, total}]
    }
    
    return dfs(0, target)
}