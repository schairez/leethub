func max(a, b int) int { if a >= b { return a}; return b }

func findMaxConsecutiveOnes(nums []int) int {
    maxTotal, localTotal := 0, 0
    for i := range nums {
        if nums[i] == 1 {
            localTotal++
            maxTotal = max(maxTotal, localTotal)
        } else {
            localTotal = 0
        }
    }
    return maxTotal
}