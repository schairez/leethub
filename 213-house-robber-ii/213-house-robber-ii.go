
// 213. House Robber II
// time: O(2n) ≈ O(n)
// space: O(2n) ≈ O(n)


func max(a, b int) int { if a >= b { return a}; return b}

func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    return max(robHelper(nums[1:]), robHelper(nums[:n-1]))
}

func robHelper(nums []int) int {
    n := len(nums)
    dp := make([]int, n+1)
    dp[1] = nums[0]
    for i := 2; i <= n; i++ {
        dp[i] = max(nums[i-1]+dp[i-2], dp[i-1])
    }
    return dp[n]
}