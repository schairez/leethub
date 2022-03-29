func max(a, b int) int { if a >= b { return a}; return b}

func findMaxConsecutiveOnes(nums []int) int {
    k := 1
    n := len(nums)
    res := 0
    numZero := 0
    start, end := 0, 0
    for end < n {
        if nums[end] == 0 {
            numZero++
        }
        // generalized for any k
        for numZero > k {
            if nums[start] == 0 {
                numZero--
            }
            start++
        } 
        res = max(res, end - start + 1)
        end++
    }
    
    return res
}