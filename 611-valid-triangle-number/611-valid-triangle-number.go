


func triangleNumber(nums []int) int {
    n := len(nums)
    if n < 3 {
        return 0
    }
    sort.Ints(nums)
    res := 0
    //for targetIdx := n-1; targetIdx > n
    //prev := n-1
    //init indices
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
    
    
    
    