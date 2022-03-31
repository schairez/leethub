// 209. Minimum Size Subarray Sum
// time: O(n)
// space: O(1)

func min(a, b int) int { if a <= b { return a}; return b}

func minSubArrayLen(target int, nums []int) int {
    const maxInt32 = 1 << 31 -1
    res := maxInt32
    n := len(nums)
    start, end := 0, 0
    subArrSum := 0
    for end < n {
        subArrSum += nums[end]
        for subArrSum >= target {
            res = min(res, end - start + 1)
            subArrSum -= nums[start]
            start++
        }
        end++
    }
    if res == maxInt32 {
        return 0
    } 
    return res
}