
// 611. Valid Triangle Number
// time: O(nlogn) to sort + O(n^2) ≈ O(n^2)
// space: O(logn) space depth internal sort space

func triangleNumber(nums []int) int {
    n := len(nums)
    if n < 3 {
        return 0
    }
    sort.Ints(nums)
    res := 0
    for target := n-1; target > 1; target-- {
        targetVal := nums[target]
        start, end := 0, target-1
        for start < end {
            twoSum := nums[start] + nums[end] 
            if twoSum > targetVal {
                res += end - start
                end--
            } else {
                start++
            }
        }
    }
    
    return res
}
    
    
    
    