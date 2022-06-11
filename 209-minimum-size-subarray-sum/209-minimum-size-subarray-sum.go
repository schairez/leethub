func minSubArrayLen(target int, nums []int) int {
    maxInt32 := 1 << 31 -1
    bestL := maxInt32
    n := len(nums)
    currSum := 0
    left, right := 0, 0
    for right < n {
        currSum += nums[right]
        for currSum >= target {
            if cand := right - left + 1; cand < bestL {
                bestL = cand
            }
            currSum -= nums[left]
            left++
        }
        right++
    }
    
    if bestL == maxInt32 {
        return 0
    }
    
    return bestL
}