//state transition fn
// dp[i] = max{ dp[i], dp[j]+1 | nums[i] > nums[j] for all j < i}

//time: O(n^2)
//space: O(n)

func max(a, b int) int { if a >= b { return a }; return b}
//iterative
func lengthOfLIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    res := 1
    dp[0] = 1
    for i:=1; i < n; i++ {
        dp[i] = 1
        for j:=0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
        res = max(dp[i], res)
    }
    
    return res
}

