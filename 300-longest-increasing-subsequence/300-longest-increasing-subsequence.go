// 300. Longest Increasing Subsequence
// time: O(nlogn)
// space: O(n)

func lengthOfLIS(nums []int) int {
    n := len(nums)
    // dp[i] = smallest LIS value for subsequence len i
    dp := make([]int, n+1)
    k := 0
    dp[k] = -1 << 31
    for i := 0; i < n; i++ {
        if nums[i] > dp[k] {
            k++
            dp[k] = nums[i]
        } else {
            lo, hi := 0, k
            for lo+1 < hi {
                mid := lo + (hi - lo) >> 1
                if dp[mid] >= nums[i] {
                    hi = mid
                } else {
                    lo = mid
                }
            }
            dp[hi] = nums[i]
        }
    }
    return k
}