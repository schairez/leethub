func min(a, b int) int { if a <= b { return a}; return b}

// time: O(n^2) space: O(n)
func jump(nums []int) int {
    n := len(nums)
    memo := make([]int, n)
    for i := range memo {
        memo[i] = 10_000
    }
    var dfs func(int) int
    dfs = func(idx int) int {
        if idx >= n - 1 {
            return 0
        }
        if memo[idx] != 10_000 {
            return memo[idx]
        }
        minJumps := 10_000
        nextIdx, furthestIdx := idx + 1, idx + nums[idx]
        for nextIdx <= furthestIdx {
            minJumps = min(minJumps, 1 + dfs(nextIdx)) 
            nextIdx++
        }
        memo[idx] = minJumps
        return memo[idx]
    }
    return dfs(0)
}