

// nums[i] + nums[j] + nums[k] < target.

// time: sort takes O(nlogn) + O(n^2) ≈ O(n^2)
// space: O(logn)
func threeSumSmaller(nums []int, target int) int {
    n := len(nums)
    if n < 3 {
        return 0
    }
    sort.Ints(nums)
    // a + b + c < target
    // a + b < target - c
    // 
    res := 0
    for thirdIdx := n-1; thirdIdx > 1; thirdIdx-- {
        thirdVal := nums[thirdIdx]
        start, end := 0, thirdIdx-1
        for start < end {
            threeSum := nums[start] + nums[end] + thirdVal
            if threeSum < target {
                res += end - start
                start++
            } else {
                end--
            } 
        }
    } 
    return res
}