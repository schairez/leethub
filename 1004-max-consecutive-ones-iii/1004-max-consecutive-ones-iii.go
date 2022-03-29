
// 1004. Max Consecutive Ones III
// time: O(n)
// space: O(1)

func max(a, b int) int { if a >= b { return a}; return b}

func longestOnes(nums []int, k int) int {
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